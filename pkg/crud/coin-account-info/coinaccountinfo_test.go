package coinaccountinfo

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

func assertCoinAccount(t *testing.T, actual, expected *npool.CoinAccountInfo) {
	assert.Equal(t, actual.CoinTypeID, expected.CoinTypeID)
	assert.Equal(t, actual.Address, expected.Address)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	coinAccount := npool.CoinAccountInfo{
		CoinTypeID:             uuid.New().String(),
		Address:                uuid.New().String(),
		PlatformHoldPrivateKey: false,
	}

	resp, err := Create(context.Background(), &npool.CreateCoinAccountRequest{
		Info: &coinAccount,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertCoinAccount(t, resp.Info, &coinAccount)
	}

	resp1, err := Get(context.Background(), &npool.GetCoinAccountRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertCoinAccount(t, resp1.Info, &coinAccount)
	}

	resp3, err := Delete(context.Background(), &npool.DeleteCoinAccountRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertCoinAccount(t, resp3.Info, &coinAccount)
	}
}
