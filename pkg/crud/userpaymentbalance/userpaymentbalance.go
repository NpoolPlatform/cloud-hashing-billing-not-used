package userpaymentbalance

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userpaymentbalance"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func validateUserPaymentBalance(info *npool.UserPaymentBalance) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}
	if _, err := uuid.Parse(info.GetPaymentID()); err != nil {
		return xerrors.Errorf("invalid payment id: %v", err)
	}
	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		return xerrors.Errorf("invalid coin type id: %v", err)
	}
	if info.GetAmount() <= 0.0 {
		return xerrors.Errorf("invalid amount")
	}
	if info.GetCoinUSDCurrency() <= 0 {
		return xerrors.Errorf("invalid coin usd currency")
	}
	return nil
}

func dbRowToUserPaymentBalance(row *ent.UserPaymentBalance) *npool.UserPaymentBalance {
	return &npool.UserPaymentBalance{
		ID:              row.ID.String(),
		AppID:           row.AppID.String(),
		UserID:          row.UserID.String(),
		PaymentID:       row.PaymentID.String(),
		UsedByPaymentID: row.UsedByPaymentID.String(),
		Amount:          price.DBPriceToVisualPrice(row.Amount),
		CoinTypeID:      row.CoinTypeID.String(),
		CoinUSDCurrency: price.DBPriceToVisualPrice(row.CoinUsdCurrency),
	}
}

func Create(ctx context.Context, in *npool.CreateUserPaymentBalanceRequest) (*npool.CreateUserPaymentBalanceResponse, error) {
	if err := validateUserPaymentBalance(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		UserPaymentBalance.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetPaymentID(uuid.MustParse(in.GetInfo().GetPaymentID())).
		SetUsedByPaymentID(uuid.UUID{}).
		SetAmount(price.VisualPriceToDBPrice(in.GetInfo().GetAmount())).
		SetUsedByPaymentID(uuid.UUID{}).
		SetCoinTypeID(uuid.MustParse(in.GetInfo().GetCoinTypeID())).
		SetCoinUsdCurrency(price.VisualPriceToDBPrice(in.GetInfo().GetCoinUSDCurrency())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create user payment balance: %v", err)
	}

	return &npool.CreateUserPaymentBalanceResponse{
		Info: dbRowToUserPaymentBalance(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateUserPaymentBalanceRequest) (*npool.UpdateUserPaymentBalanceResponse, error) {
	if err := validateUserPaymentBalance(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	usedByPaymentID, err := uuid.Parse(in.GetInfo().GetUsedByPaymentID())
	if err != nil {
		return nil, xerrors.Errorf("invalid used by payment id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		UserPaymentBalance.
		UpdateOneID(id).
		SetUsedByPaymentID(usedByPaymentID).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create user payment balance: %v", err)
	}

	return &npool.UpdateUserPaymentBalanceResponse{
		Info: dbRowToUserPaymentBalance(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetUserPaymentBalanceRequest) (*npool.GetUserPaymentBalanceResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserPaymentBalance.
		Query().
		Where(
			userpaymentbalance.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user payment balance: %v", err)
	}

	var userPaymentBalance *npool.UserPaymentBalance
	for _, info := range infos {
		userPaymentBalance = dbRowToUserPaymentBalance(info)
		break
	}

	return &npool.GetUserPaymentBalanceResponse{
		Info: userPaymentBalance,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetUserPaymentBalancesByAppRequest) (*npool.GetUserPaymentBalancesByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserPaymentBalance.
		Query().
		Where(
			userpaymentbalance.AppID(appID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user payment balance: %v", err)
	}

	userPaymentBalances := []*npool.UserPaymentBalance{}
	for _, info := range infos {
		userPaymentBalances = append(userPaymentBalances, dbRowToUserPaymentBalance(info))
	}

	return &npool.GetUserPaymentBalancesByAppResponse{
		Infos: userPaymentBalances,
	}, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetUserPaymentBalancesByAppUserRequest) (*npool.GetUserPaymentBalancesByAppUserResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	userID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserPaymentBalance.
		Query().
		Where(
			userpaymentbalance.And(
				userpaymentbalance.AppID(appID),
				userpaymentbalance.UserID(userID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user payment balance: %v", err)
	}

	userPaymentBalances := []*npool.UserPaymentBalance{}
	for _, info := range infos {
		userPaymentBalances = append(userPaymentBalances, dbRowToUserPaymentBalance(info))
	}

	return &npool.GetUserPaymentBalancesByAppUserResponse{
		Infos: userPaymentBalances,
	}, nil
}
