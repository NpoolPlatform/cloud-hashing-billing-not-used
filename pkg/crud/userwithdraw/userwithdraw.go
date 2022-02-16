package userwithdraw

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userwithdraw"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateUserWithdraw(info *npool.UserWithdraw) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}
	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		return xerrors.Errorf("invalid coin type id: %v", err)
	}
	if _, err := uuid.Parse(info.GetAccountID()); err != nil {
		return xerrors.Errorf("invalid account id: %v", err)
	}
	return nil
}

func dbRowToUserWithdraw(row *ent.UserWithdraw) *npool.UserWithdraw {
	return &npool.UserWithdraw{
		ID:         row.ID.String(),
		AppID:      row.AppID.String(),
		UserID:     row.UserID.String(),
		CoinTypeID: row.CoinTypeID.String(),
		AccountID:  row.AccountID.String(),
		CreateAt:   row.CreateAt,
		Name:       row.Name,
		Message:    row.Message,
	}
}

func Create(ctx context.Context, in *npool.CreateUserWithdrawRequest) (*npool.CreateUserWithdrawResponse, error) {
	if err := validateUserWithdraw(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		UserWithdraw.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetCoinTypeID(uuid.MustParse(in.GetInfo().GetCoinTypeID())).
		SetAccountID(uuid.MustParse(in.GetInfo().GetAccountID())).
		SetName(in.GetInfo().GetName()).
		SetMessage(in.GetInfo().GetMessage()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create user withdraw: %v", err)
	}

	return &npool.CreateUserWithdrawResponse{
		Info: dbRowToUserWithdraw(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateUserWithdrawRequest) (*npool.UpdateUserWithdrawResponse, error) {
	if err := validateUserWithdraw(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	id, err := uuid.Parse(in.GetInfo().GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		UserWithdraw.
		UpdateOneID(id).
		SetName(in.GetInfo().GetName()).
		SetMessage(in.GetInfo().GetMessage()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update user withdraw: %v", err)
	}

	return &npool.UpdateUserWithdrawResponse{
		Info: dbRowToUserWithdraw(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetUserWithdrawRequest) (*npool.GetUserWithdrawResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserWithdraw.
		Query().
		Where(
			userwithdraw.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user withdraw: %v", err)
	}

	var userWithdraw *npool.UserWithdraw
	for _, info := range infos {
		userWithdraw = dbRowToUserWithdraw(info)
		break
	}

	return &npool.GetUserWithdrawResponse{
		Info: userWithdraw,
	}, nil
}

func GetByAccount(ctx context.Context, in *npool.GetUserWithdrawByAccountRequest) (*npool.GetUserWithdrawByAccountResponse, error) {
	accountID, err := uuid.Parse(in.GetAccountID())
	if err != nil {
		return nil, xerrors.Errorf("invalid account id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserWithdraw.
		Query().
		Where(
			userwithdraw.AccountID(accountID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user withdraw: %v", err)
	}

	var userWithdraw *npool.UserWithdraw
	for _, info := range infos {
		userWithdraw = dbRowToUserWithdraw(info)
		break
	}

	return &npool.GetUserWithdrawByAccountResponse{
		Info: userWithdraw,
	}, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetUserWithdrawsByAppUserRequest) (*npool.GetUserWithdrawsByAppUserResponse, error) {
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
		UserWithdraw.
		Query().
		Where(
			userwithdraw.And(
				userwithdraw.AppID(appID),
				userwithdraw.UserID(userID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user withdraw: %v", err)
	}

	userWithdraws := []*npool.UserWithdraw{}
	for _, info := range infos {
		userWithdraws = append(userWithdraws, dbRowToUserWithdraw(info))
	}

	return &npool.GetUserWithdrawsByAppUserResponse{
		Infos: userWithdraws,
	}, nil
}
