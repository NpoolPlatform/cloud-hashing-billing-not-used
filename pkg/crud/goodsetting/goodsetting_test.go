package goodsetting

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

func assertGoodSetting(t *testing.T, actual, expected *npool.GoodSetting) {
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.WarmAccountCoinAmount, expected.WarmAccountCoinAmount)
	assert.Equal(t, actual.WarmAccountUSDAmount, expected.WarmAccountUSDAmount)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	goodSetting := npool.GoodSetting{
		GoodID:                uuid.New().String(),
		WarmAccountCoinAmount: 100,
		WarmAccountUSDAmount:  100,
	}

	resp, err := Create(context.Background(), &npool.CreateGoodSettingRequest{
		Info: &goodSetting,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertGoodSetting(t, resp.Info, &goodSetting)
	}

	goodSetting.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateGoodSettingRequest{
		Info: &goodSetting,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertGoodSetting(t, resp1.Info, &goodSetting)
	}

	resp2, err := GetByGood(context.Background(), &npool.GetGoodSettingByGoodRequest{
		GoodID: goodSetting.GoodID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertGoodSetting(t, resp2.Info, &goodSetting)
	}

	resp3, err := Get(context.Background(), &npool.GetGoodSettingRequest{
		ID: resp2.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertGoodSetting(t, resp3.Info, &goodSetting)
	}
}
