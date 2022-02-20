package userwithdrawitem

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userwithdrawitem"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func validateUserWithdrawItem(info *npool.UserWithdrawItem) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}
	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		return xerrors.Errorf("invalid coin type id: %v", err)
	}
	if _, err := uuid.Parse(info.GetWithdrawToAccountID()); err != nil {
		return xerrors.Errorf("invalid account id: %v", err)
	}
	return nil
}

func dbRowToUserWithdrawItem(row *ent.UserWithdrawItem) *npool.UserWithdrawItem {
	return &npool.UserWithdrawItem{
		ID:                    row.ID.String(),
		AppID:                 row.AppID.String(),
		UserID:                row.UserID.String(),
		CoinTypeID:            row.CoinTypeID.String(),
		WithdrawToAccountID:   row.WithdrawToAccountID.String(),
		Amount:                price.DBPriceToVisualPrice(row.Amount),
		PlatformTransactionID: row.PlatformTransactionID.String(),
	}
}

func Create(ctx context.Context, in *npool.CreateUserWithdrawItemRequest) (*npool.CreateUserWithdrawItemResponse, error) {
	if err := validateUserWithdrawItem(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		UserWithdrawItem.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetCoinTypeID(uuid.MustParse(in.GetInfo().GetCoinTypeID())).
		SetWithdrawToAccountID(uuid.MustParse(in.GetInfo().GetWithdrawToAccountID())).
		SetAmount(price.VisualPriceToDBPrice(in.GetInfo().GetAmount())).
		SetPlatformTransactionID(uuid.UUID{}).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create user withdraw item: %v", err)
	}

	return &npool.CreateUserWithdrawItemResponse{
		Info: dbRowToUserWithdrawItem(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateUserWithdrawItemRequest) (*npool.UpdateUserWithdrawItemResponse, error) {
	if err := validateUserWithdrawItem(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	txID, err := uuid.Parse(in.GetInfo().GetPlatformTransactionID())
	if err != nil {
		return nil, xerrors.Errorf("invalid platform tx id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		UserWithdrawItem.
		UpdateOneID(id).
		SetPlatformTransactionID(txID).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update user withdraw item: %v", err)
	}

	return &npool.UpdateUserWithdrawItemResponse{
		Info: dbRowToUserWithdrawItem(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetUserWithdrawItemRequest) (*npool.GetUserWithdrawItemResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserWithdrawItem.
		Query().
		Where(
			userwithdrawitem.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user withdraw item: %v", err)
	}

	var userWithdrawItem *npool.UserWithdrawItem
	for _, info := range infos {
		userWithdrawItem = dbRowToUserWithdrawItem(info)
		break
	}

	return &npool.GetUserWithdrawItemResponse{
		Info: userWithdrawItem,
	}, nil
}

func GetByAccount(ctx context.Context, in *npool.GetUserWithdrawItemsByAccountRequest) (*npool.GetUserWithdrawItemsByAccountResponse, error) {
	accountID, err := uuid.Parse(in.GetAccountID())
	if err != nil {
		return nil, xerrors.Errorf("invalid account id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserWithdrawItem.
		Query().
		Where(
			userwithdrawitem.WithdrawToAccountID(accountID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user withdraw item: %v", err)
	}

	userWithdrawItems := []*npool.UserWithdrawItem{}
	for _, info := range infos {
		userWithdrawItems = append(userWithdrawItems, dbRowToUserWithdrawItem(info))
	}

	return &npool.GetUserWithdrawItemsByAccountResponse{
		Infos: userWithdrawItems,
	}, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetUserWithdrawItemsByAppUserRequest) (*npool.GetUserWithdrawItemsByAppUserResponse, error) {
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
		UserWithdrawItem.
		Query().
		Where(
			userwithdrawitem.And(
				userwithdrawitem.AppID(appID),
				userwithdrawitem.UserID(userID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user withdraw item: %v", err)
	}

	userWithdrawItems := []*npool.UserWithdrawItem{}
	for _, info := range infos {
		userWithdrawItems = append(userWithdrawItems, dbRowToUserWithdrawItem(info))
	}

	return &npool.GetUserWithdrawItemsByAppUserResponse{
		Infos: userWithdrawItems,
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetUserWithdrawItemsRequest) (*npool.GetUserWithdrawItemsResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserWithdrawItem.
		Query().
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user withdraw item: %v", err)
	}

	userWithdrawItems := []*npool.UserWithdrawItem{}
	for _, info := range infos {
		userWithdrawItems = append(userWithdrawItems, dbRowToUserWithdrawItem(info))
	}

	return &npool.GetUserWithdrawItemsResponse{
		Infos: userWithdrawItems,
	}, nil
}
