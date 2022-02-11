package platformsetting

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/platformsetting"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validatePlatformSetting(info *npool.PlatformSetting) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	return nil
}

func dbRowToPlatformSetting(row *ent.PlatformSetting) *npool.PlatformSetting {
	return &npool.PlatformSetting{
		ID:                   row.ID.String(),
		AppID:                row.AppID.String(),
		WarmAccountUSDAmount: price.DBPriceToVisualPrice(row.WarmAccountUsdAmount),
	}
}

func Create(ctx context.Context, in *npool.CreatePlatformSettingRequest) (*npool.CreatePlatformSettingResponse, error) {
	if err := validatePlatformSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		PlatformSetting.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetWarmAccountUsdAmount(price.VisualPriceToDBPrice(in.GetInfo().GetWarmAccountUSDAmount())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create platform setting: %v", err)
	}

	return &npool.CreatePlatformSettingResponse{
		Info: dbRowToPlatformSetting(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdatePlatformSettingRequest) (*npool.UpdatePlatformSettingResponse, error) {
	if err := validatePlatformSetting(in.GetInfo()); err != nil {
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
		PlatformSetting.
		UpdateOneID(id).
		SetWarmAccountUsdAmount(price.VisualPriceToDBPrice(in.GetInfo().GetWarmAccountUSDAmount())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update platform setting: %v", err)
	}

	return &npool.UpdatePlatformSettingResponse{
		Info: dbRowToPlatformSetting(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetPlatformSettingRequest) (*npool.GetPlatformSettingResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		PlatformSetting.
		Query().
		Where(
			platformsetting.And(
				platformsetting.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query platform setting: %v", err)
	}

	var setting *npool.PlatformSetting
	for _, info := range infos {
		setting = dbRowToPlatformSetting(info)
		break
	}

	return &npool.GetPlatformSettingResponse{
		Info: setting,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetPlatformSettingByAppRequest) (*npool.GetPlatformSettingByAppResponse, error) {
	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		PlatformSetting.
		Query().
		Where(
			platformsetting.And(
				platformsetting.AppID(appID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query platform setting: %v", err)
	}

	var setting *npool.PlatformSetting
	for _, info := range infos {
		setting = dbRowToPlatformSetting(info)
		break
	}

	return &npool.GetPlatformSettingByAppResponse{
		Info: setting,
	}, nil
}
