package coinsetting

import (
	"context"

	accountcrud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coin-account-info"
	crud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coinsetting"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"golang.org/x/xerrors"
)

func validateCoinSettingAccount(ctx context.Context, accountID string, platformHoldPrivateKey bool) error {
	resp, err := accountcrud.Get(ctx, &npool.GetCoinAccountRequest{
		ID: accountID,
	})
	if err != nil {
		return xerrors.Errorf("fail get account: %v", err)
	}
	if resp.Info == nil {
		return xerrors.Errorf("fail get account")
	}
	if resp.Info.PlatformHoldPrivateKey != platformHoldPrivateKey {
		return xerrors.Errorf("different hold private key by platform")
	}

	// TODO: check of account is used for other

	return nil
}

func validateCoinSetting(ctx context.Context, info *npool.CoinSetting) error {
	if err := validateCoinSettingAccount(ctx, info.GetPlatformOfflineAccountID(), false); err != nil {
		return xerrors.Errorf("invalid coin setting account: %v", err)
	}
	if err := validateCoinSettingAccount(ctx, info.GetUserOfflineAccountID(), false); err != nil {
		return xerrors.Errorf("invalid coin setting account: %v", err)
	}
	if err := validateCoinSettingAccount(ctx, info.GetUserOnlineAccountID(), true); err != nil {
		return xerrors.Errorf("invalid coin setting account: %v", err)
	}
	if err := validateCoinSettingAccount(ctx, info.GetGoodIncomingAccountID(), false); err != nil {
		return xerrors.Errorf("invalid coin setting account: %v", err)
	}
	if err := validateCoinSettingAccount(ctx, info.GetGasProviderAccountID(), true); err != nil {
		return xerrors.Errorf("invalid coin setting account: %v", err)
	}

	validates := map[string]struct{}{}
	validates[info.GetPlatformOfflineAccountID()] = struct{}{}
	validates[info.GetUserOfflineAccountID()] = struct{}{}
	validates[info.GetUserOnlineAccountID()] = struct{}{}
	validates[info.GetGoodIncomingAccountID()] = struct{}{}
	validates[info.GetGasProviderAccountID()] = struct{}{}

	if len(validates) < 5 {
		return xerrors.Errorf("cannot use same account for different usage")
	}

	return nil
}

func Create(ctx context.Context, in *npool.CreateCoinSettingRequest) (*npool.CreateCoinSettingResponse, error) {
	if err := validateCoinSetting(ctx, in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid coin setting account: %v", err)
	}
	return crud.Create(ctx, in)
}

func Update(ctx context.Context, in *npool.UpdateCoinSettingRequest) (*npool.UpdateCoinSettingResponse, error) {
	if err := validateCoinSetting(ctx, in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid coin setting account: %v", err)
	}
	return crud.Update(ctx, in)
}
