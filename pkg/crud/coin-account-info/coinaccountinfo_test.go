package coinaccountinfo

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

func assertCoinAccount(t *testing.T, actual, expected *npool.CoinAccountInfo) {
	assert.Equal(t, actual.CoinTypeID, expected.CoinTypeID)
	assert.Equal(t, actual.Address, expected.Address)
	assert.Equal(t, actual.GeneratedBy, expected.GeneratedBy)
	assert.Equal(t, actual.UsedFor, expected.UsedFor)
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.UserID, expected.UserID)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	coinAccount := npool.CoinAccountInfo{
		CoinTypeID:  uuid.New().String(),
		AppID:       uuid.New().String(),
		UserID:      uuid.New().String(),
		GeneratedBy: "user",
		UsedFor:     "benefit",
	}

	resp, err := Create(context.Background(), &npool.CreateCoinAccountRequest{
		Info: &coinAccount,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assert.True(t, resp.Info.Idle)
		assertCoinAccount(t, resp.Info, &coinAccount)
	}

	resp1, err := Get(context.Background(), &npool.GetCoinAccountRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assert.True(t, resp1.Info.Idle)
		assertCoinAccount(t, resp1.Info, &coinAccount)
	}

	resp2, err := GetCoinAccountsByAppUser(context.Background(), &npool.GetCoinAccountsByAppUserRequest{
		AppID:  coinAccount.AppID,
		UserID: coinAccount.UserID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp2.Infos), 1)
	}
}
