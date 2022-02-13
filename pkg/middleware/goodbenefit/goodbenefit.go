package goodbenefit

import (
	"context"

	accountcrud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coin-account-info"
	crud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/goodbenefit"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"golang.org/x/xerrors"
)

func validateGoodBenefitAccount(ctx context.Context, accountID string) error {
	resp, err := accountcrud.Get(ctx, &npool.GetCoinAccountRequest{
		ID: accountID,
	})
	if err != nil {
		return xerrors.Errorf("fail get account: %v", err)
	}
	if resp.Info == nil {
		return xerrors.Errorf("fail get account")
	}
	if !resp.Info.PlatformHoldPrivateKey {
		return xerrors.Errorf("should hold private key by platform")
	}

	// TODO: check of account is used for other

	return nil
}

func validateGoodBenefit(ctx context.Context, info *npool.GoodBenefit) error {
	if err := validateGoodBenefitAccount(ctx, info.GetBenefitAccountID()); err != nil {
		return xerrors.Errorf("invalid good benefit account: %v", err)
	}
	if err := validateGoodBenefitAccount(ctx, info.GetPlatformOfflineAccountID()); err != nil {
		return xerrors.Errorf("invalid good benefit account: %v", err)
	}
	if err := validateGoodBenefitAccount(ctx, info.GetUserOfflineAccountID()); err != nil {
		return xerrors.Errorf("invalid good benefit account: %v", err)
	}
	if err := validateGoodBenefitAccount(ctx, info.GetUserOnlineAccountID()); err != nil {
		return xerrors.Errorf("invalid good benefit account: %v", err)
	}
	return nil
}

func Create(ctx context.Context, in *npool.CreateGoodBenefitRequest) (*npool.CreateGoodBenefitResponse, error) {
	if err := validateGoodBenefit(ctx, in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid good benefit account: %v", err)
	}
	return crud.Create(ctx, in)
}

func Update(ctx context.Context, in *npool.UpdateGoodBenefitRequest) (*npool.UpdateGoodBenefitResponse, error) {
	if err := validateGoodBenefit(ctx, in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid good benefit account: %v", err)
	}
	return crud.Update(ctx, in)
}
