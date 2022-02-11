package coinsetting

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

func assertCoinSetting(t *testing.T, actual, expected *npool.CoinSetting) {
	assert.Equal(t, actual.CoinTypeID, expected.CoinTypeID)
	assert.Equal(t, actual.WarmAccountCoinAmount, expected.WarmAccountCoinAmount)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	platformSetting := npool.CoinSetting{
		CoinTypeID:            uuid.New().String(),
		WarmAccountCoinAmount: 100,
	}

	resp, err := Create(context.Background(), &npool.CreateCoinSettingRequest{
		Info: &platformSetting,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertCoinSetting(t, resp.Info, &platformSetting)
	}

	platformSetting.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateCoinSettingRequest{
		Info: &platformSetting,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertCoinSetting(t, resp1.Info, &platformSetting)
	}

	resp2, err := GetByCoin(context.Background(), &npool.GetCoinSettingByCoinRequest{
		CoinTypeID: platformSetting.CoinTypeID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertCoinSetting(t, resp2.Info, &platformSetting)
	}

	resp3, err := Get(context.Background(), &npool.GetCoinSettingRequest{
		ID: resp2.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertCoinSetting(t, resp3.Info, &platformSetting)
	}
}
