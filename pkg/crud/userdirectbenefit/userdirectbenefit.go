package userdirectbenefit

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userdirectbenefit"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateUserDirectBenefit(info *npool.UserDirectBenefit) error {
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

func dbRowToUserDirectBenefit(row *ent.UserDirectBenefit) *npool.UserDirectBenefit {
	return &npool.UserDirectBenefit{
		ID:         row.ID.String(),
		AppID:      row.AppID.String(),
		UserID:     row.UserID.String(),
		CoinTypeID: row.CoinTypeID.String(),
		AccountID:  row.AccountID.String(),
	}
}

func Create(ctx context.Context, in *npool.CreateUserDirectBenefitRequest) (*npool.CreateUserDirectBenefitResponse, error) {
	if err := validateUserDirectBenefit(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		UserDirectBenefit.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetCoinTypeID(uuid.MustParse(in.GetInfo().GetCoinTypeID())).
		SetAccountID(uuid.MustParse(in.GetInfo().GetAccountID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create user direct benefit: %v", err)
	}

	return &npool.CreateUserDirectBenefitResponse{
		Info: dbRowToUserDirectBenefit(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateUserDirectBenefitRequest) (*npool.UpdateUserDirectBenefitResponse, error) {
	if err := validateUserDirectBenefit(in.GetInfo()); err != nil {
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
		UserDirectBenefit.
		UpdateOneID(id).
		SetAccountID(uuid.MustParse(in.GetInfo().GetAccountID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update user direct benefit: %v", err)
	}

	return &npool.UpdateUserDirectBenefitResponse{
		Info: dbRowToUserDirectBenefit(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetUserDirectBenefitRequest) (*npool.GetUserDirectBenefitResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserDirectBenefit.
		Query().
		Where(
			userdirectbenefit.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user direct benefit: %v", err)
	}

	var userWithdraw *npool.UserDirectBenefit
	for _, info := range infos {
		userWithdraw = dbRowToUserDirectBenefit(info)
		break
	}

	return &npool.GetUserDirectBenefitResponse{
		Info: userWithdraw,
	}, nil
}

func GetByAccount(ctx context.Context, in *npool.GetUserDirectBenefitByAccountRequest) (*npool.GetUserDirectBenefitByAccountResponse, error) {
	accountID, err := uuid.Parse(in.GetAccountID())
	if err != nil {
		return nil, xerrors.Errorf("invalid account id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserDirectBenefit.
		Query().
		Where(
			userdirectbenefit.AccountID(accountID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user direct benefit: %v", err)
	}

	var userWithdraw *npool.UserDirectBenefit
	for _, info := range infos {
		userWithdraw = dbRowToUserDirectBenefit(info)
		break
	}

	return &npool.GetUserDirectBenefitByAccountResponse{
		Info: userWithdraw,
	}, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetUserDirectBenefitsByAppUserRequest) (*npool.GetUserDirectBenefitsByAppUserResponse, error) {
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
		UserDirectBenefit.
		Query().
		Where(
			userdirectbenefit.And(
				userdirectbenefit.AppID(appID),
				userdirectbenefit.UserID(userID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user direct benefit: %v", err)
	}

	userWithdraws := []*npool.UserDirectBenefit{}
	for _, info := range infos {
		userWithdraws = append(userWithdraws, dbRowToUserDirectBenefit(info))
	}

	return &npool.GetUserDirectBenefitsByAppUserResponse{
		Infos: userWithdraws,
	}, nil
}
