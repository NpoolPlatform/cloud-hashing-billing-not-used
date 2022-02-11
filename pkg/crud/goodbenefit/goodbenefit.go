package goodbenefit

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/goodbenefit"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validateGoodBenefit(info *npool.GoodBenefit) error {
	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		return xerrors.Errorf("invalid good id: %v", err)
	}
	if _, err := uuid.Parse(info.GetBenefitAccountID()); err != nil {
		return xerrors.Errorf("invalid benefit account id: %v", err)
	}
	if _, err := uuid.Parse(info.GetPlatformOfflineAccountID()); err != nil {
		return xerrors.Errorf("invalid platform offline account id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserOfflineAccountID()); err != nil {
		return xerrors.Errorf("invalid user offline account id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserOnlineAccountID()); err != nil {
		return xerrors.Errorf("invalid user online account id: %v", err)
	}
	if info.GetBenefitIntervalHours() == 0 {
		return xerrors.Errorf("invalid benefit interval")
	}
	return nil
}

func dbRowToGoodBenefit(row *ent.GoodBenefit) *npool.GoodBenefit {
	return &npool.GoodBenefit{
		ID:                       row.ID.String(),
		GoodID:                   row.GoodID.String(),
		BenefitAccountID:         row.BenefitAccountID.String(),
		PlatformOfflineAccountID: row.PlatformOfflineAccountID.String(),
		UserOfflineAccountID:     row.UserOfflineAccountID.String(),
		UserOnlineAccountID:      row.UserOnlineAccountID.String(),
		BenefitIntervalHours:     row.BenefitIntervalHours,
	}
}

func Create(ctx context.Context, in *npool.CreateGoodBenefitRequest) (*npool.CreateGoodBenefitResponse, error) {
	if err := validateGoodBenefit(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		GoodBenefit.
		Create().
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetBenefitAccountID(uuid.MustParse(in.GetInfo().GetBenefitAccountID())).
		SetPlatformOfflineAccountID(uuid.MustParse(in.GetInfo().GetPlatformOfflineAccountID())).
		SetUserOfflineAccountID(uuid.MustParse(in.GetInfo().GetUserOfflineAccountID())).
		SetUserOnlineAccountID(uuid.MustParse(in.GetInfo().GetUserOnlineAccountID())).
		SetBenefitIntervalHours(in.GetInfo().GetBenefitIntervalHours()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create good benefit: %v", err)
	}

	return &npool.CreateGoodBenefitResponse{
		Info: dbRowToGoodBenefit(info),
	}, nil
}

func Update(ctx context.Context, in *npool.UpdateGoodBenefitRequest) (*npool.UpdateGoodBenefitResponse, error) {
	if err := validateGoodBenefit(in.GetInfo()); err != nil {
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
		GoodBenefit.
		UpdateOneID(id).
		SetBenefitAccountID(uuid.MustParse(in.GetInfo().GetBenefitAccountID())).
		SetPlatformOfflineAccountID(uuid.MustParse(in.GetInfo().GetPlatformOfflineAccountID())).
		SetUserOfflineAccountID(uuid.MustParse(in.GetInfo().GetUserOfflineAccountID())).
		SetUserOnlineAccountID(uuid.MustParse(in.GetInfo().GetUserOnlineAccountID())).
		SetBenefitIntervalHours(in.GetInfo().GetBenefitIntervalHours()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update good benefit: %v", err)
	}

	return &npool.UpdateGoodBenefitResponse{
		Info: dbRowToGoodBenefit(info),
	}, nil
}

func Get(ctx context.Context, in *npool.GetGoodBenefitRequest) (*npool.GetGoodBenefitResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		GoodBenefit.
		Query().
		Where(
			goodbenefit.ID(id),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query good benefit: %v", err)
	}

	var setting *npool.GoodBenefit
	for _, info := range infos {
		setting = dbRowToGoodBenefit(info)
		break
	}

	return &npool.GetGoodBenefitResponse{
		Info: setting,
	}, nil
}

func GetByGood(ctx context.Context, in *npool.GetGoodBenefitByGoodRequest) (*npool.GetGoodBenefitByGoodResponse, error) {
	coinID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid coin id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		GoodBenefit.
		Query().
		Where(
			goodbenefit.GoodID(coinID),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query good benefit: %v", err)
	}

	var setting *npool.GoodBenefit
	for _, info := range infos {
		setting = dbRowToGoodBenefit(info)
		break
	}

	return &npool.GetGoodBenefitByGoodResponse{
		Info: setting,
	}, nil
}
