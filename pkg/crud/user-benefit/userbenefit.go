package userbenefit

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

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
	if _, err := uuid.Parse(info.GetPlatformTransactionID()); err != nil {
		return xerrors.Errorf("invalid platform transaction id: %v", err)
	}
	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		return xerrors.Errorf("invalid coin type id: %v", err)
	}
	if info.GetLastBenefitTimestamp() == 0 {
		return xerrors.Errorf("invalid last benefit timestamp")
	}
	return nil
}

func dbRowToUserBenefit(row *ent.UserBenefit) *npool.UserBenefit {
	return &npool.UserBenefit{
		ID:                    row.ID.String(),
		GoodID:                row.GoodID.String(),
		AppID:                 row.AppID.String(),
		UserID:                row.UserID.String(),
		Amount:                price.DBPriceToVisualPrice(row.Amount),
		OrderID:               row.OrderID.String(),
		CoinTypeID:            row.CoinTypeID.String(),
		CreateAt:              row.CreateAt,
		LastBenefitTimestamp:  row.LastBenefitTimestamp,
		PlatformTransactionID: row.PlatformTransactionID.String(),
	}
}

func Create(ctx context.Context, in *npool.CreateUserBenefitRequest) (*npool.CreateUserBenefitResponse, error) {
	if err := validateUserBenefit(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		UserBenefit.
		Create().
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetAmount(price.VisualPriceToDBPrice(in.GetInfo().GetAmount())).
		SetOrderID(uuid.MustParse(in.GetInfo().GetOrderID())).
		SetCoinTypeID(uuid.MustParse(in.GetInfo().GetCoinTypeID())).
		SetLastBenefitTimestamp(in.GetInfo().GetLastBenefitTimestamp()).
		SetPlatformTransactionID(uuid.MustParse(in.GetInfo().GetPlatformTransactionID())).
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

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
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

	benefits := []*npool.UserBenefit{}
	for _, info := range infos {
		benefits = append(benefits, dbRowToUserBenefit(info))
	}

	return &npool.GetUserBenefitsByAppUserResponse{
		Infos: benefits,
	}, nil
}

func GetLatestByGoodAppUser(ctx context.Context, in *npool.GetLatestUserBenefitByGoodAppUserRequest) (*npool.GetLatestUserBenefitByGoodAppUserResponse, error) {
	goodID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

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
		UserBenefit.
		Query().
		Order(
			ent.Desc(userbenefit.FieldCreateAt),
		).
		Where(
			userbenefit.GoodID(goodID),
			userbenefit.AppID(appID),
			userbenefit.UserID(userID),
		).
		Limit(1).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user benefit: %v", err)
	}

	var benefit *npool.UserBenefit
	if len(infos) > 0 {
		benefit = dbRowToUserBenefit(infos[0])
	}

	return &npool.GetLatestUserBenefitByGoodAppUserResponse{
		Info: benefit,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetUserBenefitsByAppRequest) (*npool.GetUserBenefitsByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserBenefit.
		Query().
		Where(
			userbenefit.AppID(appID),
		).
		Limit(200).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("failed to query user benefit db, %v", err)
	}

	benefits := []*npool.UserBenefit{}
	for _, info := range infos {
		benefits = append(benefits, dbRowToUserBenefit(info))
	}

	return &npool.GetUserBenefitsByAppResponse{
		Infos: benefits,
	}, nil
}

func GetByAppUserCoin(ctx context.Context, in *npool.GetUserBenefitsByAppUserCoinRequest) (*npool.GetUserBenefitsByAppUserCoinResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	userID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		return nil, xerrors.Errorf("invalid user id: %v", err)
	}

	coinTypeID, err := uuid.Parse(in.GetCoinTypeID())
	if err != nil {
		return nil, xerrors.Errorf("invalid coin type id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserBenefit.
		Query().
		Where(
			userbenefit.And(
				userbenefit.AppID(appID),
				userbenefit.UserID(userID),
				userbenefit.CoinTypeID(coinTypeID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user benefit: %v", err)
	}

	benefits := []*npool.UserBenefit{}
	for _, info := range infos {
		benefits = append(benefits, dbRowToUserBenefit(info))
	}

	return &npool.GetUserBenefitsByAppUserCoinResponse{
		Infos: benefits,
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetUserBenefitsRequest) (*npool.GetUserBenefitsResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		UserBenefit.
		Query().
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query user benefit: %v", err)
	}

	benefits := []*npool.UserBenefit{}
	for _, info := range infos {
		benefits = append(benefits, dbRowToUserBenefit(info))
	}

	return &npool.GetUserBenefitsResponse{
		Infos: benefits,
	}, nil
}
