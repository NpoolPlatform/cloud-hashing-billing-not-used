package userwithdraw

import (
	"context"

	accountcrud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coin-account-info"
	crud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/userwithdraw"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"golang.org/x/xerrors"
)

func GetUserWithdrawAccount(ctx context.Context, in *npool.GetUserWithdrawAccountRequest) (*npool.GetUserWithdrawAccountResponse, error) {
	resp, err := crud.Get(ctx, &npool.GetUserWithdrawRequest{
		ID: in.GetID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get user withdraw: %v", err)
	}

	resp1, err := accountcrud.Get(ctx, &npool.GetCoinAccountRequest{
		ID: resp.Info.AccountID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get account: %v", err)
	}

	return &npool.GetUserWithdrawAccountResponse{
		Info: &npool.UserWithdrawAccount{
			Withdraw: resp.Info,
			Account:  resp1.Info,
		},
	}, nil
}

func GetUserWithdrawAccountsByAppUser(ctx context.Context, in *npool.GetUserWithdrawAccountsByAppUserRequest) (*npool.GetUserWithdrawAccountsByAppUserResponse, error) {
	resp, err := crud.GetByAppUser(ctx, &npool.GetUserWithdrawsByAppUserRequest{
		AppID:  in.GetAppID(),
		UserID: in.GetUserID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get user withdraw: %v", err)
	}

	accounts := []*npool.UserWithdrawAccount{}
	for _, info := range resp.Infos {
		account, err := accountcrud.Get(ctx, &npool.GetCoinAccountRequest{
			ID: info.AccountID,
		})
		if err != nil {
			return nil, xerrors.Errorf("fail get account: %v", err)
		}
		accounts = append(accounts, &npool.UserWithdrawAccount{
			Withdraw: info,
			Account:  account.Info,
		})
	}

	return &npool.GetUserWithdrawAccountsByAppUserResponse{
		Infos: accounts,
	}, nil
}
