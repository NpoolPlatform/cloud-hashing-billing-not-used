package goodincoming

import (
	"context"

	accountcrud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coin-account-info"
	crud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/goodincoming"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"golang.org/x/xerrors"
)

func validateGoodIncomingAccount(ctx context.Context, accountID string) error {
	resp, err := accountcrud.Get(ctx, &npool.GetCoinAccountRequest{
		ID: accountID,
	})
	if err != nil {
		return xerrors.Errorf("fail get account: %v", err)
	}
	if resp.Info == nil {
		return xerrors.Errorf("fail get account")
	}
	if resp.Info.PlatformHoldPrivateKey {
		return xerrors.Errorf("good incoming account should be an offline address")
	}

	// TODO: check of account is used for other

	return nil
}

func Create(ctx context.Context, in *npool.CreateGoodIncomingRequest) (*npool.CreateGoodIncomingResponse, error) {
	if err := validateGoodIncomingAccount(ctx, in.GetInfo().GetAccountID()); err != nil {
		return nil, xerrors.Errorf("invalid good incoming account: %v", err)
	}
	return crud.Create(ctx, in)
}

func Update(ctx context.Context, in *npool.UpdateGoodIncomingRequest) (*npool.UpdateGoodIncomingResponse, error) {
	if err := validateGoodIncomingAccount(ctx, in.GetInfo().GetAccountID()); err != nil {
		return nil, xerrors.Errorf("invalid good incoming account: %v", err)
	}
	return crud.Update(ctx, in)
}
