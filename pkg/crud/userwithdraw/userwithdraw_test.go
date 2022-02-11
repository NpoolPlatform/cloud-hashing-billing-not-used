package userwithdraw

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

func assertUserWithdraw(t *testing.T, actual, expected *npool.UserWithdraw) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.CoinTypeID, expected.CoinTypeID)
	assert.Equal(t, actual.AccountID, expected.AccountID)
	assert.Equal(t, actual.Name, expected.Name)
	assert.Equal(t, actual.Message, expected.Message)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	userWithdraw := npool.UserWithdraw{
		AppID:      uuid.New().String(),
		UserID:     uuid.New().String(),
		CoinTypeID: uuid.New().String(),
		AccountID:  uuid.New().String(),
		Name:       uuid.New().String(),
		Message:    uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateUserWithdrawRequest{
		Info: &userWithdraw,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertUserWithdraw(t, resp.Info, &userWithdraw)
	}

	userWithdraw.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateUserWithdrawRequest{
		Info: &userWithdraw,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertUserWithdraw(t, resp1.Info, &userWithdraw)
	}

	resp2, err := GetByAccount(context.Background(), &npool.GetUserWithdrawByAccountRequest{
		AccountID: userWithdraw.AccountID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertUserWithdraw(t, resp2.Info, &userWithdraw)
	}

	resp3, err := Get(context.Background(), &npool.GetUserWithdrawRequest{
		ID: resp2.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertUserWithdraw(t, resp3.Info, &userWithdraw)
	}

	resp4, err := GetByAppUser(context.Background(), &npool.GetUserWithdrawsByAppUserRequest{
		AppID:  userWithdraw.AppID,
		UserID: userWithdraw.UserID,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(resp4.Infos), 0)
	}
}
