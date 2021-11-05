package platformsetting

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-billing/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/test-init" //nolint

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

func assertPlatformSetting(t *testing.T, actual, expected *npool.PlatformSetting) {
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.BenefitAccountID, expected.BenefitAccountID)
	assert.Equal(t, actual.PlatformOfflineAccountID, expected.PlatformOfflineAccountID)
	assert.Equal(t, actual.UserOnlineAccountID, expected.UserOnlineAccountID)
	assert.Equal(t, actual.UserOfflineAccountID, expected.UserOfflineAccountID)
	assert.Equal(t, actual.BenefitIntervalHours, expected.BenefitIntervalHours)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	platformSetting := npool.PlatformSetting{
		GoodID:                   uuid.New().String(),
		BenefitAccountID:         uuid.New().String(),
		PlatformOfflineAccountID: uuid.New().String(),
		UserOnlineAccountID:      uuid.New().String(),
		UserOfflineAccountID:     uuid.New().String(),
		BenefitIntervalHours:     24,
	}

	resp, err := Create(context.Background(), &npool.CreatePlatformSettingRequest{
		Info: &platformSetting,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertPlatformSetting(t, resp.Info, &platformSetting)
	}

	platformSetting.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdatePlatformSettingRequest{
		Info: &platformSetting,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertPlatformSetting(t, resp1.Info, &platformSetting)
	}

	resp2, err := GetByGood(context.Background(), &npool.GetPlatformSettingByGoodRequest{
		GoodID: platformSetting.GoodID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertPlatformSetting(t, resp2.Info, &platformSetting)
	}
}
