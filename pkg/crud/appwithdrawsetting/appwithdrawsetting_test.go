package appwithdrawsetting

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

func assertAppWithdrawSetting(t *testing.T, actual, expected *npool.AppWithdrawSetting) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.CoinTypeID, expected.CoinTypeID)
	assert.Equal(t, actual.WithdrawAutoReviewCoinAmount, expected.WithdrawAutoReviewCoinAmount)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	setting := npool.AppWithdrawSetting{
		AppID:                        uuid.New().String(),
		CoinTypeID:                   uuid.New().String(),
		WithdrawAutoReviewCoinAmount: 100,
	}

	resp, err := Create(context.Background(), &npool.CreateAppWithdrawSettingRequest{
		Info: &setting,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertAppWithdrawSetting(t, resp.Info, &setting)
	}

	setting.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateAppWithdrawSettingRequest{
		Info: &setting,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertAppWithdrawSetting(t, resp1.Info, &setting)
	}

	resp2, err := GetByAppCoin(context.Background(), &npool.GetAppWithdrawSettingByAppCoinRequest{
		AppID:      setting.AppID,
		CoinTypeID: setting.CoinTypeID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertAppWithdrawSetting(t, resp2.Info, &setting)
	}

	resp3, err := Get(context.Background(), &npool.GetAppWithdrawSettingRequest{
		ID: resp2.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertAppWithdrawSetting(t, resp3.Info, &setting)
	}
}
