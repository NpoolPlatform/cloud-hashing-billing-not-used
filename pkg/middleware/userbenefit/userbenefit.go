package userbenefit

import (
	"context"

	trans "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coin-account-transaction"
	crud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/user-benefit"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	constant "github.com/NpoolPlatform/cloud-hashing-billing/pkg/const"

	"golang.org/x/xerrors"
)

func filter(ctx context.Context, infos []*npool.UserBenefit) ([]*npool.UserBenefit, error) {
	txs := map[string]*npool.CoinAccountTransaction{}
	benefits := []*npool.UserBenefit{}

	for _, info := range infos {
		tx, ok := txs[info.PlatformTransactionID]
		if !ok {
			resp1, err := trans.Get(ctx, &npool.GetCoinAccountTransactionRequest{
				ID: info.PlatformTransactionID,
			})
			if err != nil {
				return nil, xerrors.Errorf("fail get platform transaction: %v", err)
			}
			txs[info.PlatformTransactionID] = resp1.Info
			tx = resp1.Info
		}

		if tx.State == constant.CoinTransactionStateFail || tx.State == constant.CoinTransactionStateRejected {
			continue
		}

		benefits = append(benefits, info)
	}

	return benefits, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetUserBenefitsByAppUserRequest) (*npool.GetUserBenefitsByAppUserResponse, error) {
	resp, err := crud.GetByAppUser(ctx, in)
	if err != nil {
		return nil, xerrors.Errorf("fail get by app user: %v", err)
	}

	benefits, err := filter(ctx, resp.Infos)
	if err != nil {
		return nil, xerrors.Errorf("fail filter user benefits: %v", err)
	}

	return &npool.GetUserBenefitsByAppUserResponse{
		Infos: benefits,
	}, nil
}

func GetByAppUserCoin(ctx context.Context, in *npool.GetUserBenefitsByAppUserCoinRequest) (*npool.GetUserBenefitsByAppUserCoinResponse, error) {
	resp, err := crud.GetByAppUserCoin(ctx, in)
	if err != nil {
		return nil, xerrors.Errorf("fail get by app user coin: %v", err)
	}

	benefits, err := filter(ctx, resp.Infos)
	if err != nil {
		return nil, xerrors.Errorf("fail filter user benefits: %v", err)
	}

	return &npool.GetUserBenefitsByAppUserCoinResponse{
		Infos: benefits,
	}, nil
}
