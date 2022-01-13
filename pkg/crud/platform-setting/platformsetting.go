package platformsetting

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/platformsetting"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validatePlatformSetting(info *npool.PlatformSetting) error {
	_, err := uuid.Parse(info.GetGoodID())
	if err != nil {
		return xerrors.Errorf("invalid good id: %v", err)
	}
	_, err = uuid.Parse(info.GetBenefitAccountID())
	if err != nil {
		return xerrors.Errorf("invalid benefit account id: %v", err)
	}
	_, err = uuid.Parse(info.GetPlatformOfflineAccountID())
	if err != nil {
		return xerrors.Errorf("invalid platform offline account id: %v", err)
	}
	_, err = uuid.Parse(info.GetUserOnlineAccountID())
	if err != nil {
		return xerrors.Errorf("invalid user online account id: %v", err)
	}
	_, err = uuid.Parse(info.GetUserOfflineAccountID())
	if err != nil {
		return xerrors.Errorf("invalid user offline id: %v", err)
	}
	return nil
}

func dbRowToPlatformSetting(row *ent.PlatformSetting) *npool.PlatformSetting {
	return &npool.PlatformSetting{
		ID:                       row.ID.String(),
		GoodID:                   row.GoodID.String(),
		BenefitAccountID:         row.BenefitAccountID.String(),
		PlatformOfflineAccountID: row.PlatformOfflineAccountID.String(),
		UserOnlineAccountID:      row.UserOnlineAccountID.String(),
		UserOfflineAccountID:     row.UserOfflineAccountID.String(),
		BenefitIntervalHours:     row.BenefitIntervalHours,
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
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetBenefitAccountID(uuid.MustParse(in.GetInfo().GetBenefitAccountID())).
		SetPlatformOfflineAccountID(uuid.MustParse(in.GetInfo().GetPlatformOfflineAccountID())).
		SetUserOnlineAccountID(uuid.MustParse(in.GetInfo().GetUserOnlineAccountID())).
		SetUserOfflineAccountID(uuid.MustParse(in.GetInfo().GetUserOfflineAccountID())).
		SetBenefitIntervalHours(in.GetInfo().GetBenefitIntervalHours()).
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
		SetBenefitAccountID(uuid.MustParse(in.GetInfo().GetBenefitAccountID())).
		SetPlatformOfflineAccountID(uuid.MustParse(in.GetInfo().GetPlatformOfflineAccountID())).
		SetUserOnlineAccountID(uuid.MustParse(in.GetInfo().GetUserOnlineAccountID())).
		SetUserOfflineAccountID(uuid.MustParse(in.GetInfo().GetUserOfflineAccountID())).
		SetBenefitIntervalHours(in.GetInfo().GetBenefitIntervalHours()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update platform setting: %v", err)
	}

	return &npool.UpdatePlatformSettingResponse{
		Info: dbRowToPlatformSetting(info),
	}, nil
}

func GetByGood(ctx context.Context, in *npool.GetPlatformSettingByGoodRequest) (*npool.GetPlatformSettingByGoodResponse, error) {
	goodID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
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
				platformsetting.GoodID(goodID),
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

	return &npool.GetPlatformSettingByGoodResponse{
		Info: setting,
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
	if len(infos) == 0 {
		return nil, xerrors.Errorf("empty query result")
	}

	return &npool.GetPlatformSettingResponse{
		Info: dbRowToPlatformSetting(infos[0]),
	}, nil
}
