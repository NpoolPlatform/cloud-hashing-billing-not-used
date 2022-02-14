package userwithdrawitem

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

func assertUserWithdrawItem(t *testing.T, actual, expected *npool.UserWithdrawItem) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.CoinTypeID, expected.CoinTypeID)
	assert.Equal(t, actual.WithdrawToAccountID, expected.WithdrawToAccountID)
	assert.Equal(t, actual.Amount, expected.Amount)
	assert.Equal(t, actual.PlatformTransactionID, expected.PlatformTransactionID)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	userWithdrawItem := npool.UserWithdrawItem{
		AppID:                 uuid.New().String(),
		UserID:                uuid.New().String(),
		CoinTypeID:            uuid.New().String(),
		WithdrawToAccountID:   uuid.New().String(),
		Amount:                1.0,
		PlatformTransactionID: uuid.UUID{}.String(),
	}

	resp, err := Create(context.Background(), &npool.CreateUserWithdrawItemRequest{
		Info: &userWithdrawItem,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertUserWithdrawItem(t, resp.Info, &userWithdrawItem)
	}

	userWithdrawItem.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateUserWithdrawItemRequest{
		Info: &userWithdrawItem,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertUserWithdrawItem(t, resp1.Info, &userWithdrawItem)
	}

	resp2, err := GetByAccount(context.Background(), &npool.GetUserWithdrawItemsByAccountRequest{
		AccountID: userWithdrawItem.WithdrawToAccountID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Infos[0].ID, resp.Info.ID)
		assertUserWithdrawItem(t, resp2.Infos[0], &userWithdrawItem)
	}

	resp3, err := Get(context.Background(), &npool.GetUserWithdrawItemRequest{
		ID: resp1.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertUserWithdrawItem(t, resp3.Info, &userWithdrawItem)
	}

	resp4, err := GetByAppUser(context.Background(), &npool.GetUserWithdrawItemsByAppUserRequest{
		AppID:  userWithdrawItem.AppID,
		UserID: userWithdrawItem.UserID,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(resp4.Infos), 0)
	}
}
