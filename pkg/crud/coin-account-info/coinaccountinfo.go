package coinaccountinfo

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-billing/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/coinaccountinfo" //nolint

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateCoinAccount(info *npool.CoinAccountInfo) error {
	if _, err := uuid.Parse(info.CoinTypeID); err != nil {
		return xerrors.Errorf("invalid coin type id: %v", err)
	}
	if _, err := uuid.Parse(info.AppID); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.UserID); err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}
	return nil
}

func dbRowToCoinAccount(row *ent.CoinAccountInfo) *npool.CoinAccountInfo {
	return &npool.CoinAccountInfo{
		ID:          row.ID.String(),
		CoinTypeID:  row.CoinTypeID.String(),
		Address:     row.Address,
		GeneratedBy: string(row.GeneratedBy),
		UsedFor:     string(row.UsedFor),
		Idle:        row.Idle,
		AppID:       row.AppID.String(),
		UserID:      row.UserID.String(),
	}
}

func Create(ctx context.Context, in *npool.CreateCoinAccountRequest) (*npool.CreateCoinAccountResponse, error) {
	if err := validateCoinAccount(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	info, err := db.Client().
		CoinAccountInfo.
		Create().
		SetCoinTypeID(uuid.MustParse(in.GetInfo().GetCoinTypeID())).
		SetAddress(in.GetInfo().GetAddress()).
		SetGeneratedBy(coinaccountinfo.GeneratedBy(in.GetInfo().GetGeneratedBy())).
		SetUsedFor(coinaccountinfo.UsedFor(in.GetInfo().GetUsedFor())).
		SetIdle(true).
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create coin account: %v", err)
	}

	return &npool.CreateCoinAccountResponse{
		Info: dbRowToCoinAccount(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetCoinAccountRequest) (*npool.GetCoinAccountResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid coin account id: %v", err)
	}

	infos, err := db.Client().
		CoinAccountInfo.
		Query().
		Where(
			coinaccountinfo.And(
				coinaccountinfo.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query coin account: %v", err)
	}
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty coin account")
	}

	return &npool.GetCoinAccountResponse{
		Info: dbRowToCoinAccount(infos[0]),
	}, nil
}

func GetCoinAccountsByAppUser(ctx context.Context, in *npool.GetCoinAccountsByAppUserRequest) (*npool.GetCoinAccountsByAppUserResponse, error) {
	return nil, nil
}

func Delete(ctx context.Context, in *npool.DeleteCoinAccountRequest) (*npool.DeleteCoinAccountResponse, error) {
	return nil, nil
}
