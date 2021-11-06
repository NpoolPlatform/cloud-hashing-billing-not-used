package platformsetting

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-billing/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/test-init" //nolint

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coin-account-info" //nolint
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/platform-setting"  //nolint

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestGet(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	account1 := npool.CoinAccountInfo{
		CoinTypeID:  uuid.New().String(),
		Address:     uuid.New().String(),
		GeneratedBy: "platform",
		UsedFor:     "benefit",
		AppID:       uuid.New().String(),
		UserID:      uuid.New().String(),
	}

	resp1, err := coinaccountinfo.Create(context.Background(), &npool.CreateCoinAccountRequest{
		Info: &account1,
	})
	assert.Nil(t, err)

	setting := npool.PlatformSetting{
		GoodID:                   uuid.New().String(),
		BenefitAccountID:         resp1.Info.ID,
		PlatformOfflineAccountID: resp1.Info.ID,
		UserOnlineAccountID:      resp1.Info.ID,
		UserOfflineAccountID:     resp1.Info.ID,
		BenefitIntervalHours:     24,
	}
	resp2, err := platformsetting.Create(context.Background(), &npool.CreatePlatformSettingRequest{
		Info: &setting,
	})
	assert.Nil(t, err)

	resp3, err := Get(context.Background(), &npool.GetPlatformSettingDetailRequest{
		ID: resp2.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Detail.ID, resp2.Info.ID)
		assert.Equal(t, resp3.Detail.GoodID, resp2.Info.GoodID)

		assert.Equal(t, resp3.Detail.BenefitAddress.ID, resp1.Info.ID)
		assert.Equal(t, resp3.Detail.BenefitAddress.CoinTypeID, resp1.Info.CoinTypeID)
		assert.Equal(t, resp3.Detail.BenefitAddress.Address, resp1.Info.Address)
		assert.Equal(t, resp3.Detail.BenefitAddress.GeneratedBy, resp1.Info.GeneratedBy)
		assert.Equal(t, resp3.Detail.BenefitAddress.UsedFor, resp1.Info.UsedFor)
		assert.Equal(t, resp3.Detail.BenefitAddress.Idle, resp1.Info.Idle)
		assert.Equal(t, resp3.Detail.BenefitAddress.AppID, resp1.Info.AppID)
		assert.Equal(t, resp3.Detail.BenefitAddress.UserID, resp1.Info.UserID)

		assert.Equal(t, resp3.Detail.PlatformOfflineAddress.ID, resp1.Info.ID)
		assert.Equal(t, resp3.Detail.PlatformOfflineAddress.CoinTypeID, resp1.Info.CoinTypeID)
		assert.Equal(t, resp3.Detail.PlatformOfflineAddress.Address, resp1.Info.Address)
		assert.Equal(t, resp3.Detail.PlatformOfflineAddress.GeneratedBy, resp1.Info.GeneratedBy)
		assert.Equal(t, resp3.Detail.PlatformOfflineAddress.UsedFor, resp1.Info.UsedFor)
		assert.Equal(t, resp3.Detail.PlatformOfflineAddress.Idle, resp1.Info.Idle)
		assert.Equal(t, resp3.Detail.PlatformOfflineAddress.AppID, resp1.Info.AppID)
		assert.Equal(t, resp3.Detail.PlatformOfflineAddress.UserID, resp1.Info.UserID)

		assert.Equal(t, resp3.Detail.UserOnlineAddress.ID, resp1.Info.ID)
		assert.Equal(t, resp3.Detail.UserOnlineAddress.CoinTypeID, resp1.Info.CoinTypeID)
		assert.Equal(t, resp3.Detail.UserOnlineAddress.Address, resp1.Info.Address)
		assert.Equal(t, resp3.Detail.UserOnlineAddress.GeneratedBy, resp1.Info.GeneratedBy)
		assert.Equal(t, resp3.Detail.UserOnlineAddress.UsedFor, resp1.Info.UsedFor)
		assert.Equal(t, resp3.Detail.UserOnlineAddress.Idle, resp1.Info.Idle)
		assert.Equal(t, resp3.Detail.UserOnlineAddress.AppID, resp1.Info.AppID)
		assert.Equal(t, resp3.Detail.UserOnlineAddress.UserID, resp1.Info.UserID)

		assert.Equal(t, resp3.Detail.UserOfflineAddress.ID, resp1.Info.ID)
		assert.Equal(t, resp3.Detail.UserOfflineAddress.CoinTypeID, resp1.Info.CoinTypeID)
		assert.Equal(t, resp3.Detail.UserOfflineAddress.Address, resp1.Info.Address)
		assert.Equal(t, resp3.Detail.UserOfflineAddress.GeneratedBy, resp1.Info.GeneratedBy)
		assert.Equal(t, resp3.Detail.UserOfflineAddress.UsedFor, resp1.Info.UsedFor)
		assert.Equal(t, resp3.Detail.UserOfflineAddress.Idle, resp1.Info.Idle)
		assert.Equal(t, resp3.Detail.UserOfflineAddress.AppID, resp1.Info.AppID)
		assert.Equal(t, resp3.Detail.UserOfflineAddress.UserID, resp1.Info.UserID)

		assert.Equal(t, resp3.Detail.BenefitIntervalHours, resp2.Info.BenefitIntervalHours)
	}
}
