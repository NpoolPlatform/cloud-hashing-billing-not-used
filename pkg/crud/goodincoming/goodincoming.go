package goodincoming

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/goodincoming"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateGoodIncoming(info *npool.GoodIncoming) error {
	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		return xerrors.Errorf("invalid good id: %v", err)
	}
	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		return xerrors.Errorf("invalid coin type id: %v", err)
	}
	if _, err := uuid.Parse(info.GetAccountID()); err != nil {
		return xerrors.Errorf("invalid account id: %v", err)
	}
	return nil
}

func dbRowToGoodIncoming(row *ent.GoodIncoming) *npool.GoodIncoming {
	return &npool.GoodIncoming{
		ID:         row.ID.String(),
		GoodID:     row.GoodID.String(),
		CoinTypeID: row.CoinTypeID.String(),
		AccountID:  row.AccountID.String(),
	}
}

func Create(ctx context.Context, in *npool.CreateGoodIncomingRequest) (*npool.CreateGoodIncomingResponse, error) {
	if err := validateGoodIncoming(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		GoodIncoming.
		Create().
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetCoinTypeID(uuid.MustParse(in.GetInfo().GetCoinTypeID())).
		SetAccountID(uuid.MustParse(in.GetInfo().GetAccountID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create good incoming: %v", err)
	}

	return &npool.CreateGoodIncomingResponse{
		Info: dbRowToGoodIncoming(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateGoodIncomingRequest) (*npool.UpdateGoodIncomingResponse, error) {
	if err := validateGoodIncoming(in.GetInfo()); err != nil {
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
		GoodIncoming.
		UpdateOneID(id).
		SetAccountID(uuid.MustParse(in.GetInfo().GetAccountID())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update good incoming: %v", err)
	}

	return &npool.UpdateGoodIncomingResponse{
		Info: dbRowToGoodIncoming(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetGoodIncomingRequest) (*npool.GetGoodIncomingResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		GoodIncoming.
		Query().
		Where(
			goodincoming.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query good incoming: %v", err)
	}

	var setting *npool.GoodIncoming
	for _, info := range infos {
		setting = dbRowToGoodIncoming(info)
		break
	}

	return &npool.GetGoodIncomingResponse{
		Info: setting,
	}, nil
}

func GetByGood(ctx context.Context, in *npool.GetGoodIncomingByGoodRequest) (*npool.GetGoodIncomingByGoodResponse, error) {
	goodID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		GoodIncoming.
		Query().
		Where(
			goodincoming.GoodID(goodID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query good incoming: %v", err)
	}

	settings := []*npool.GoodIncoming{}
	for _, info := range infos {
		settings = append(settings, dbRowToGoodIncoming(info))
	}

	return &npool.GetGoodIncomingByGoodResponse{
		Infos: settings,
	}, nil
}

func GetByGoodCoin(ctx context.Context, in *npool.GetGoodIncomingByGoodCoinRequest) (*npool.GetGoodIncomingByGoodCoinResponse, error) {
	goodID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
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
		GoodIncoming.
		Query().
		Where(
			goodincoming.And(
				goodincoming.GoodID(goodID),
				goodincoming.CoinTypeID(coinID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query good incoming: %v", err)
	}

	var setting *npool.GoodIncoming
	for _, info := range infos {
		setting = dbRowToGoodIncoming(info)
		break
	}

	return &npool.GetGoodIncomingByGoodCoinResponse{
		Info: setting,
	}, nil
}
