package platformsetting

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-billing/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coin-account-info" //nolint
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/platform-setting"  //nolint

	"golang.org/x/xerrors"
)

func Get(ctx context.Context, in *npool.GetPlatformSettingDetailRequest) (*npool.GetPlatformSettingDetailResponse, error) {
	setting, err := platformsetting.Get(ctx, &npool.GetPlatformSettingRequest{
		ID: in.GetID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get platform setting: %v", err)
	}

	benefitAddress, err := coinaccountinfo.Get(ctx, &npool.GetCoinAccountRequest{
		ID: setting.Info.BenefitAccountID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get benefit account: %v", err)
	}

	platformOfflineAddress, err := coinaccountinfo.Get(ctx, &npool.GetCoinAccountRequest{
		ID: setting.Info.PlatformOfflineAccountID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get platform offline account: %v", err)
	}

	userOnlineAddress, err := coinaccountinfo.Get(ctx, &npool.GetCoinAccountRequest{
		ID: setting.Info.UserOnlineAccountID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get user online account: %v", err)
	}

	userOfflineAddress, err := coinaccountinfo.Get(ctx, &npool.GetCoinAccountRequest{
		ID: setting.Info.UserOfflineAccountID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get user offline account: %v", err)
	}

	return &npool.GetPlatformSettingDetailResponse{
		Detail: &npool.PlatformSettingDetail{
			ID:                     setting.Info.ID,
			GoodID:                 setting.Info.GoodID,
			BenefitAddress:         benefitAddress.Info,
			PlatformOfflineAddress: platformOfflineAddress.Info,
			UserOnlineAddress:      userOnlineAddress.Info,
			UserOfflineAddress:     userOfflineAddress.Info,
			BenefitIntervalHours:   setting.Info.BenefitIntervalHours,
		},
	}, nil
}
