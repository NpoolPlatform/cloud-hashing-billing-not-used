package appwithdrawsetting

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/appwithdrawsetting"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateAppWithdrawSetting(info *npool.AppWithdrawSetting) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		return xerrors.Errorf("invalid coin type id: %v", err)
	}
	return nil
}

func dbRowToAppWithdrawSetting(row *ent.AppWithdrawSetting) *npool.AppWithdrawSetting {
	return &npool.AppWithdrawSetting{
		ID:                           row.ID.String(),
		AppID:                        row.AppID.String(),
		CoinTypeID:                   row.CoinTypeID.String(),
		WithdrawAutoReviewCoinAmount: price.DBPriceToVisualPrice(row.WithdrawAutoReviewCoinAmount),
	}
}

func Create(ctx context.Context, in *npool.CreateAppWithdrawSettingRequest) (*npool.CreateAppWithdrawSettingResponse, error) {
	if err := validateAppWithdrawSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		AppWithdrawSetting.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetCoinTypeID(uuid.MustParse(in.GetInfo().GetCoinTypeID())).
		SetWithdrawAutoReviewCoinAmount(price.VisualPriceToDBPrice(in.GetInfo().GetWithdrawAutoReviewCoinAmount())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app withdraw setting: %v", err)
	}

	return &npool.CreateAppWithdrawSettingResponse{
		Info: dbRowToAppWithdrawSetting(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateAppWithdrawSettingRequest) (*npool.UpdateAppWithdrawSettingResponse, error) {
	if err := validateAppWithdrawSetting(in.GetInfo()); err != nil {
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
		AppWithdrawSetting.
		UpdateOneID(id).
		SetWithdrawAutoReviewCoinAmount(price.VisualPriceToDBPrice(in.GetInfo().GetWithdrawAutoReviewCoinAmount())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update app withdraw setting: %v", err)
	}

	return &npool.UpdateAppWithdrawSettingResponse{
		Info: dbRowToAppWithdrawSetting(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetAppWithdrawSettingRequest) (*npool.GetAppWithdrawSettingResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppWithdrawSetting.
		Query().
		Where(
			appwithdrawsetting.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app withdraw setting: %v", err)
	}

	var setting *npool.AppWithdrawSetting
	for _, info := range infos {
		setting = dbRowToAppWithdrawSetting(info)
		break
	}

	return &npool.GetAppWithdrawSettingResponse{
		Info: setting,
	}, nil
}

func GetByAppCoin(ctx context.Context, in *npool.GetAppWithdrawSettingByAppCoinRequest) (*npool.GetAppWithdrawSettingByAppCoinResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid coin id: %v", err)
	}

	coinID, err := uuid.Parse(in.GetCoinTypeID())
	if err != nil {
		return nil, xerrors.Errorf("invalid coin id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppWithdrawSetting.
		Query().
		Where(
			appwithdrawsetting.AppID(appID),
			appwithdrawsetting.CoinTypeID(coinID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app withdraw setting: %v", err)
	}

	var setting *npool.AppWithdrawSetting
	for _, info := range infos {
		setting = dbRowToAppWithdrawSetting(info)
		break
	}

	return &npool.GetAppWithdrawSettingByAppCoinResponse{
		Info: setting,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetAppWithdrawSettingsByAppRequest) (*npool.GetAppWithdrawSettingsByAppResponse, error) {
	coinID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid coin id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		AppWithdrawSetting.
		Query().
		Where(
			appwithdrawsetting.AppID(coinID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app withdraw setting: %v", err)
	}

	settings := []*npool.AppWithdrawSetting{}
	for _, info := range infos {
		settings = append(settings, dbRowToAppWithdrawSetting(info))
	}

	return &npool.GetAppWithdrawSettingsByAppResponse{
		Infos: settings,
	}, nil
}
