package goodsetting

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/goodsetting"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateGoodSetting(info *npool.GoodSetting) error {
	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		return xerrors.Errorf("invalid coin type id: %v", err)
	}
	return nil
}

func dbRowToGoodSetting(row *ent.GoodSetting) *npool.GoodSetting {
	return &npool.GoodSetting{
		ID:                    row.ID.String(),
		GoodID:                row.GoodID.String(),
		WarmAccountCoinAmount: price.DBPriceToVisualPrice(row.WarmAccountCoinAmount),
		WarmAccountUSDAmount:  price.DBPriceToVisualPrice(row.WarmAccountUsdAmount),
	}
}

func Create(ctx context.Context, in *npool.CreateGoodSettingRequest) (*npool.CreateGoodSettingResponse, error) {
	if err := validateGoodSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		GoodSetting.
		Create().
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetWarmAccountCoinAmount(price.VisualPriceToDBPrice(in.GetInfo().GetWarmAccountCoinAmount())).
		SetWarmAccountUsdAmount(price.VisualPriceToDBPrice(in.GetInfo().GetWarmAccountUSDAmount())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create good setting: %v", err)
	}

	return &npool.CreateGoodSettingResponse{
		Info: dbRowToGoodSetting(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateGoodSettingRequest) (*npool.UpdateGoodSettingResponse, error) {
	if err := validateGoodSetting(in.GetInfo()); err != nil {
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
		GoodSetting.
		UpdateOneID(id).
		SetWarmAccountCoinAmount(price.VisualPriceToDBPrice(in.GetInfo().GetWarmAccountCoinAmount())).
		SetWarmAccountUsdAmount(price.VisualPriceToDBPrice(in.GetInfo().GetWarmAccountUSDAmount())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update good setting: %v", err)
	}

	return &npool.UpdateGoodSettingResponse{
		Info: dbRowToGoodSetting(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetGoodSettingRequest) (*npool.GetGoodSettingResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		GoodSetting.
		Query().
		Where(
			goodsetting.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query good setting: %v", err)
	}

	var setting *npool.GoodSetting
	for _, info := range infos {
		setting = dbRowToGoodSetting(info)
		break
	}

	return &npool.GetGoodSettingResponse{
		Info: setting,
	}, nil
}

func GetByGood(ctx context.Context, in *npool.GetGoodSettingByGoodRequest) (*npool.GetGoodSettingByGoodResponse, error) {
	coinID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid coin id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		GoodSetting.
		Query().
		Where(
			goodsetting.GoodID(coinID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query good setting: %v", err)
	}

	var setting *npool.GoodSetting
	for _, info := range infos {
		setting = dbRowToGoodSetting(info)
		break
	}

	return &npool.GetGoodSettingByGoodResponse{
		Info: setting,
	}, nil
}
