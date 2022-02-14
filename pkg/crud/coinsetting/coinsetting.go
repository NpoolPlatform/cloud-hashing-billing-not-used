package coinsetting

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/coinsetting"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateCoinSetting(info *npool.CoinSetting) error {
	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		return xerrors.Errorf("invalid coin type id: %v", err)
	}
	return nil
}

func dbRowToCoinSetting(row *ent.CoinSetting) *npool.CoinSetting {
	return &npool.CoinSetting{
		ID:                       row.ID.String(),
		CoinTypeID:               row.CoinTypeID.String(),
		WarmAccountCoinAmount:    price.DBPriceToVisualPrice(row.WarmAccountCoinAmount),
		PaymentAccountCoinAmount: price.DBPriceToVisualPrice(row.PaymentAccountCoinAmount),
	}
}

func Create(ctx context.Context, in *npool.CreateCoinSettingRequest) (*npool.CreateCoinSettingResponse, error) {
	if err := validateCoinSetting(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		CoinSetting.
		Create().
		SetCoinTypeID(uuid.MustParse(in.GetInfo().GetCoinTypeID())).
		SetWarmAccountCoinAmount(price.VisualPriceToDBPrice(in.GetInfo().GetWarmAccountCoinAmount())).
		SetPaymentAccountCoinAmount(price.VisualPriceToDBPrice(in.GetInfo().GetPaymentAccountCoinAmount())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create coin setting: %v", err)
	}

	return &npool.CreateCoinSettingResponse{
		Info: dbRowToCoinSetting(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateCoinSettingRequest) (*npool.UpdateCoinSettingResponse, error) {
	if err := validateCoinSetting(in.GetInfo()); err != nil {
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
		CoinSetting.
		UpdateOneID(id).
		SetWarmAccountCoinAmount(price.VisualPriceToDBPrice(in.GetInfo().GetWarmAccountCoinAmount())).
		SetPaymentAccountCoinAmount(price.VisualPriceToDBPrice(in.GetInfo().GetPaymentAccountCoinAmount())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update coin setting: %v", err)
	}

	return &npool.UpdateCoinSettingResponse{
		Info: dbRowToCoinSetting(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetCoinSettingRequest) (*npool.GetCoinSettingResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		CoinSetting.
		Query().
		Where(
			coinsetting.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query coin setting: %v", err)
	}

	var setting *npool.CoinSetting
	for _, info := range infos {
		setting = dbRowToCoinSetting(info)
		break
	}

	return &npool.GetCoinSettingResponse{
		Info: setting,
	}, nil
}

func GetByCoin(ctx context.Context, in *npool.GetCoinSettingByCoinRequest) (*npool.GetCoinSettingByCoinResponse, error) {
	coinID, err := uuid.Parse(in.GetCoinTypeID())
	if err != nil {
		return nil, xerrors.Errorf("invalid coin id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		CoinSetting.
		Query().
		Where(
			coinsetting.CoinTypeID(coinID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query coin setting: %v", err)
	}

	var setting *npool.CoinSetting
	for _, info := range infos {
		setting = dbRowToCoinSetting(info)
		break
	}

	return &npool.GetCoinSettingByCoinResponse{
		Info: setting,
	}, nil
}
