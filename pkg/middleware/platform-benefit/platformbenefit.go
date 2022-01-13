package platformbenefit

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coin-account-info" //nolint
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/platform-benefit"  //nolint

	"golang.org/x/xerrors"
)

func Get(ctx context.Context, in *npool.GetPlatformBenefitDetailRequest) (*npool.GetPlatformBenefitDetailResponse, error) {
	benefit, err := platformbenefit.Get(ctx, &npool.GetPlatformBenefitRequest{
		ID: in.GetID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get platform benefit: %v", err)
	}

	account, err := coinaccountinfo.Get(ctx, &npool.GetCoinAccountRequest{
		ID: benefit.Info.BenefitAccountID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get coin account: %v", err)
	}

	return &npool.GetPlatformBenefitDetailResponse{
		Detail: &npool.PlatformBenefitDetail{
			ID:                 benefit.Info.ID,
			GoodID:             benefit.Info.GoodID,
			BenefitAddress:     account.Info,
			Amount:             benefit.Info.Amount,
			ChainTransactionID: benefit.Info.ChainTransactionID,
		},
	}, nil
}
