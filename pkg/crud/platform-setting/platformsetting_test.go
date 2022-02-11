package platformsetting

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/test-init" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

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
	assert.Equal(t, actual.WarmAccountUSDAmount, expected.WarmAccountUSDAmount)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	platformSetting := npool.PlatformSetting{
		WarmAccountUSDAmount: 100,
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
}
