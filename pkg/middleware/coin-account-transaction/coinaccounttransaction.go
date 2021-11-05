package coinaccounttransaction

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-billing/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coin-account-info"        //nolint
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coin-account-transaction" //nolint

	"golang.org/x/xerrors"
)

func Get(ctx context.Context, in *npool.GetCoinAccountTransactionDetailRequest) (*npool.GetCoinAccountTransactionDetailResponse, error) {
	transaction, err := coinaccounttransaction.Get(ctx, &npool.GetCoinAccountTransactionRequest{
		ID: in.GetID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get transaction: %v", err)
	}

	fromAccount, err := coinaccountinfo.Get(ctx, &npool.GetCoinAccountRequest{
		ID: transaction.Info.FromAddressID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get transaction from address: %v", err)
	}

	toAccount, err := coinaccountinfo.Get(ctx, &npool.GetCoinAccountRequest{
		ID: transaction.Info.ToAddressID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get transaction to address: %v", err)
	}

	return &npool.GetCoinAccountTransactionDetailResponse{
		Detail: &npool.CoinAccountTransactionDetail{
			ID:                    transaction.Info.ID,
			UserID:                transaction.Info.UserID,
			AppID:                 transaction.Info.AppID,
			FromAddress:           fromAccount.Info,
			ToAddress:             toAccount.Info,
			CoinTypeID:            transaction.Info.CoinTypeID,
			Amount:                transaction.Info.Amount,
			Message:               transaction.Info.Message,
			ChainTransactionID:    transaction.Info.ChainTransactionID,
			PlatformTransactionID: transaction.Info.PlatformTransactionID,
			CreateAt:              transaction.Info.CreateAt,
			State:                 transaction.Info.State,
		},
	}, nil
}