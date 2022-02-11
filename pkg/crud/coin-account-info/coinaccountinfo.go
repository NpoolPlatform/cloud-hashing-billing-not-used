package coinaccountinfo

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/coinaccountinfo" //nolint

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateCoinAccount(info *npool.CoinAccountInfo) error {
	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		return xerrors.Errorf("invalid coin type id: %v", err)
	}
	if info.GetAddress() == "" {
		return xerrors.Errorf("invalid coin address")
	}
	return nil
}

func dbRowToCoinAccount(row *ent.CoinAccountInfo) *npool.CoinAccountInfo {
	return &npool.CoinAccountInfo{
		ID:                     row.ID.String(),
		CoinTypeID:             row.CoinTypeID.String(),
		Address:                row.Address,
		PlatformHoldPrivateKey: row.PlatformHoldPrivateKey,
	}
}

func Create(ctx context.Context, in *npool.CreateCoinAccountRequest) (*npool.CreateCoinAccountResponse, error) {
	if err := validateCoinAccount(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		CoinAccountInfo.
		Create().
		SetCoinTypeID(uuid.MustParse(in.GetInfo().GetCoinTypeID())).
		SetAddress(in.GetInfo().GetAddress()).
		SetPlatformHoldPrivateKey(in.GetInfo().GetPlatformHoldPrivateKey()).
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

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
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

func GetByCoinAddress(ctx context.Context, in *npool.GetCoinAccountByCoinAddressRequest) (*npool.GetCoinAccountByCoinAddressResponse, error) {
	coinInfoID, err := uuid.Parse(in.GetCoinInfoID())
	if err != nil {
		return nil, xerrors.Errorf("invalid coin info id: %v", err)
	}

	address := in.GetAddress()
	if address == "" {
		return nil, xerrors.Errorf("invalid coin address")
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		CoinAccountInfo.
		Query().
		Where(
			coinaccountinfo.And(
				coinaccountinfo.CoinTypeID(coinInfoID),
				coinaccountinfo.Address(address),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query coin account: %v", err)
	}

	var account *npool.CoinAccountInfo
	for _, info := range infos {
		account = dbRowToCoinAccount(info)
		break
	}

	return &npool.GetCoinAccountByCoinAddressResponse{
		Info: account,
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteCoinAccountRequest) (*npool.DeleteCoinAccountResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		CoinAccountInfo.
		UpdateOneID(uuid.MustParse(in.GetID())).
		SetDeleteAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail delete coin account: %v", err)
	}

	return &npool.DeleteCoinAccountResponse{
		Info: dbRowToCoinAccount(info),
	}, nil
}
