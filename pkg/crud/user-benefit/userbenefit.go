package userbenefit

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-billing/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userbenefit"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateUserBenefit(info *npool.UserBenefit) error {
	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		return xerrors.Errorf("invalid good id: %v", err)
	}
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}
	if _, err := uuid.Parse(info.GetOrderID()); err != nil {
		return xerrors.Errorf("invalid order id: %v", err)
	}
	return nil
}

func dbRowToUserBenefit(row *ent.UserBenefit) *npool.UserBenefit {
	return &npool.UserBenefit{
		ID:       row.ID.String(),
		GoodID:   row.GoodID.String(),
		AppID:    row.AppID.String(),
		UserID:   row.UserID.String(),
		Amount:   price.DBPriceToVisualPrice(row.Amount),
		OrderID:  row.OrderID.String(),
		CreateAt: row.CreateAt,
	}
}

func Create(ctx context.Context, in *npool.CreateUserBenefitRequest) (*npool.CreateUserBenefitResponse, error) {
	if err := validateUserBenefit(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		UserBenefit.
		Create().
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetAmount(price.VisualPriceToDBPrice(in.GetInfo().GetAmount())).
		SetOrderID(uuid.MustParse(in.GetInfo().GetOrderID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create user benefit: %v", err)
	}

	return &npool.CreateUserBenefitResponse{
		Info: dbRowToUserBenefit(info),
	}, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetUserBenefitsByAppUserRequest) (*npool.GetUserBenefitsByAppUserResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	userID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	infos, err := db.Client().
		UserBenefit.
		Query().
		Where(
			userbenefit.And(
				userbenefit.AppID(appID),
				userbenefit.UserID(userID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user benefit: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty user benefit")
	}

	benefits := []*npool.UserBenefit{}
	for _, info := range infos {
		benefits = append(benefits, dbRowToUserBenefit(info))
	}

	return &npool.GetUserBenefitsByAppUserResponse{
		Infos: benefits,
	}, nil
}
