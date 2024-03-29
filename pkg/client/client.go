package client

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/cloud-hashing-billing/pkg/message/const"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get good payment connection: %v", err)
	}
	defer conn.Close()

	cli := npool.NewCloudHashingBillingClient(conn)

	return fn(_ctx, cli)
}

func CreateGoodPayment(ctx context.Context, in *npool.GoodPayment) (*npool.GoodPayment, error) {
	// conds: NOT USED NOW, will be used after refactor code
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.CreateGoodPayment(ctx, &npool.CreateGoodPaymentRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create good payment: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create good payment: %v", err)
	}
	return info.(*npool.GoodPayment), nil
}

func GetGoodPayments(ctx context.Context, conds cruder.FilterConds) ([]*npool.GoodPayment, error) {
	// conds: NOT USED NOW, will be used after refactor code
	infos, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.GetGoodPayments(ctx, &npool.GetGoodPaymentsRequest{})
		if err != nil {
			return nil, fmt.Errorf("fail get good payments: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get good payments: %v", err)
	}
	return infos.([]*npool.GoodPayment), nil
}

func GetAccountGoodPayment(ctx context.Context, accountID string) (*npool.GoodPayment, error) {
	// conds: NOT USED NOW, will be used after refactor code
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.GetGoodPaymentByAccount(ctx, &npool.GetGoodPaymentByAccountRequest{
			AccountID: accountID,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get good payment: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get good payment: %v", err)
	}
	return info.(*npool.GoodPayment), nil
}

func GetGoodPayment(ctx context.Context, id string) (*npool.GoodPayment, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.GetGoodPayment(ctx, &npool.GetGoodPaymentRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get good payment: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get good payment: %v", err)
	}
	return info.(*npool.GoodPayment), nil
}

func GetIdleGoodPayments(ctx context.Context, goodID, coinTypeID string) ([]*npool.GoodPayment, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.GetIdleGoodPaymentsByGoodPaymentCoin(ctx, &npool.GetIdleGoodPaymentsByGoodPaymentCoinRequest{
			GoodID:            goodID,
			PaymentCoinTypeID: coinTypeID,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get good payments: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get good payments: %v", err)
	}
	return infos.([]*npool.GoodPayment), nil
}

func UpdateGoodPayment(ctx context.Context, in *npool.GoodPayment) (*npool.GoodPayment, error) {
	// conds: NOT USED NOW, will be used after refactor code
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.UpdateGoodPayment(ctx, &npool.UpdateGoodPaymentRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update good payment: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update good payment: %v", err)
	}
	return info.(*npool.GoodPayment), nil
}

func CreateAccount(ctx context.Context, in *npool.CoinAccountInfo) (*npool.CoinAccountInfo, error) {
	// conds: NOT USED NOW, will be used after refactor code
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.CreateCoinAccount(ctx, &npool.CreateCoinAccountRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create coin account: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create coin account: %v", err)
	}
	return info.(*npool.CoinAccountInfo), nil
}

func GetAccount(ctx context.Context, id string) (*npool.CoinAccountInfo, error) {
	// conds: NOT USED NOW, will be used after refactor code
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.GetCoinAccount(ctx, &npool.GetCoinAccountRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get coin account: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get coin account: %v", err)
	}
	return info.(*npool.CoinAccountInfo), nil
}

func GetAccounts(ctx context.Context) ([]*npool.CoinAccountInfo, error) {
	// conds: NOT USED NOW, will be used after refactor code
	infos, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.GetCoinAccounts(ctx, &npool.GetCoinAccountsRequest{})
		if err != nil {
			return nil, fmt.Errorf("fail get coin accounts: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get coin accounts: %v", err)
	}
	return infos.([]*npool.CoinAccountInfo), nil
}

func GetCoinSettings(ctx context.Context) ([]*npool.CoinSetting, error) {
	// conds: NOT USED NOW, will be used after refactor code
	infos, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.GetCoinSettings(ctx, &npool.GetCoinSettingsRequest{})
		if err != nil {
			return nil, fmt.Errorf("fail get coin settings: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get coin settings: %v", err)
	}
	return infos.([]*npool.CoinSetting), nil
}

func GetCoinSetting(ctx context.Context, coinTypeID string) (*npool.CoinSetting, error) {
	// conds: NOT USED NOW, will be used after refactor code
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.GetCoinSettingByCoin(ctx, &npool.GetCoinSettingByCoinRequest{
			CoinTypeID: coinTypeID,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get coin setting: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get coin setting: %v", err)
	}
	return info.(*npool.CoinSetting), nil
}

func GetPlatformSetting(ctx context.Context) (*npool.PlatformSetting, error) {
	// conds: NOT USED NOW, will be used after refactor code
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.GetPlatformSetting(ctx, &npool.GetPlatformSettingRequest{})
		if err != nil {
			return nil, fmt.Errorf("fail get coin setting: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get coin setting: %v", err)
	}
	return info.(*npool.PlatformSetting), nil
}

func CreateTransaction(ctx context.Context, tx *npool.CoinAccountTransaction) (*npool.CoinAccountTransaction, error) {
	// conds: NOT USED NOW, will be used after refactor code
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.CreateCoinAccountTransaction(ctx, &npool.CreateCoinAccountTransactionRequest{
			Info: tx,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create transaction: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create transaction: %v", err)
	}
	return info.(*npool.CoinAccountTransaction), nil
}

func UpdateTransaction(ctx context.Context, tx *npool.CoinAccountTransaction) (*npool.CoinAccountTransaction, error) {
	// conds: NOT USED NOW, will be used after refactor code
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.UpdateCoinAccountTransaction(ctx, &npool.UpdateCoinAccountTransactionRequest{
			Info: tx,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update transaction: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update transaction: %v", err)
	}
	return info.(*npool.CoinAccountTransaction), nil
}

func GetPaymentBalances(ctx context.Context, appID, userID string) ([]*npool.UserPaymentBalance, error) {
	// conds: NOT USED NOW, will be used after refactor code
	infos, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.GetUserPaymentBalancesByAppUser(ctx, &npool.GetUserPaymentBalancesByAppUserRequest{
			AppID:  appID,
			UserID: userID,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get user payment balance: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get user payment balance: %v", err)
	}
	return infos.([]*npool.UserPaymentBalance), nil
}

func CreatePaymentBalance(ctx context.Context, in *npool.UserPaymentBalance) (*npool.UserPaymentBalance, error) {
	// conds: NOT USED NOW, will be used after refactor code
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.CreateUserPaymentBalance(ctx, &npool.CreateUserPaymentBalanceRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create user payment balance: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create user payment balance: %v", err)
	}
	return info.(*npool.UserPaymentBalance), nil
}

func GetTransactions(ctx context.Context, state string) ([]*npool.CoinAccountTransaction, error) {
	// conds: NOT USED NOW, will be used after refactor code
	infos, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.GetCoinAccountTransactionsByState(ctx, &npool.GetCoinAccountTransactionsByStateRequest{
			State: state,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get transactions: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get transactions: %v", err)
	}
	return infos.([]*npool.CoinAccountTransaction), nil
}

func GetAccountTransactions(ctx context.Context, coinTypeID, accountID string) ([]*npool.CoinAccountTransaction, error) {
	// conds: NOT USED NOW, will be used after refactor code
	infos, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.GetCoinAccountTransactionsByCoinAccount(ctx, &npool.GetCoinAccountTransactionsByCoinAccountRequest{
			CoinTypeID: coinTypeID,
			AddressID:  accountID,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get transactions: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get transactions: %v", err)
	}
	return infos.([]*npool.CoinAccountTransaction), nil
}

func GetTransaction(ctx context.Context, id string) (*npool.CoinAccountTransaction, error) {
	// conds: NOT USED NOW, will be used after refactor code
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.GetCoinAccountTransaction(ctx, &npool.GetCoinAccountTransactionRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get transaction: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get transaction: %v", err)
	}
	return info.(*npool.CoinAccountTransaction), nil
}

func GetGoodBenefit(ctx context.Context, goodID string) (*npool.GoodBenefit, error) {
	// conds: NOT USED NOW, will be used after refactor code
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.GetGoodBenefitByGood(ctx, &npool.GetGoodBenefitByGoodRequest{
			GoodID: goodID,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get good benefit: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get good benefit: %v", err)
	}
	return info.(*npool.GoodBenefit), nil
}

func GetWithdrawSetting(ctx context.Context, appID, coinTypeID string) (*npool.AppWithdrawSetting, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.GetAppWithdrawSettingByAppCoin(ctx, &npool.GetAppWithdrawSettingByAppCoinRequest{
			AppID:      appID,
			CoinTypeID: coinTypeID,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get withdraw setting: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get withdraw setting: %v", err)
	}
	return info.(*npool.AppWithdrawSetting), nil
}

func GetWithdrawAccounts(ctx context.Context, appID, userID string) ([]*npool.UserWithdraw, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.GetUserWithdrawsByAppUser(ctx, &npool.GetUserWithdrawsByAppUserRequest{
			AppID:  appID,
			UserID: userID,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get user withdraw: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get user withdraw: %v", err)
	}
	return infos.([]*npool.UserWithdraw), nil
}

func GetAppWithdrawAccounts(ctx context.Context, appID string) ([]*npool.UserWithdraw, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.GetUserWithdrawsByApp(ctx, &npool.GetUserWithdrawsByAppRequest{
			AppID: appID,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get user withdraw: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get user withdraw: %v", err)
	}
	return infos.([]*npool.UserWithdraw), nil
}

func GetWithdrawAccount(ctx context.Context, accountID string) (*npool.UserWithdraw, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.CloudHashingBillingClient) (cruder.Any, error) {
		resp, err := cli.GetUserWithdrawByAccount(ctx, &npool.GetUserWithdrawByAccountRequest{
			AccountID: accountID,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get user withdraw: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get user withdraw: %v", err)
	}
	return info.(*npool.UserWithdraw), nil
}
