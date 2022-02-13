package goodincoming

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

func assertGoodIncoming(t *testing.T, actual, expected *npool.GoodIncoming) {
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.CoinTypeID, expected.CoinTypeID)
	assert.Equal(t, actual.AccountID, expected.AccountID)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	goodIncoming := npool.GoodIncoming{
		GoodID:     uuid.New().String(),
		CoinTypeID: uuid.New().String(),
		AccountID:  uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateGoodIncomingRequest{
		Info: &goodIncoming,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertGoodIncoming(t, resp.Info, &goodIncoming)
	}

	goodIncoming.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateGoodIncomingRequest{
		Info: &goodIncoming,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertGoodIncoming(t, resp1.Info, &goodIncoming)
	}

	resp2, err := GetByGood(context.Background(), &npool.GetGoodIncomingsByGoodRequest{
		GoodID: goodIncoming.GoodID,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(resp2.Infos), 0)
	}

	resp3, err := Get(context.Background(), &npool.GetGoodIncomingRequest{
		ID: resp1.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertGoodIncoming(t, resp3.Info, &goodIncoming)
	}

	resp4, err := GetByGoodCoin(context.Background(), &npool.GetGoodIncomingByGoodCoinRequest{
		GoodID:     goodIncoming.GoodID,
		CoinTypeID: goodIncoming.CoinTypeID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp4.Info.ID, resp.Info.ID)
		assertGoodIncoming(t, resp4.Info, &goodIncoming)
	}
}
