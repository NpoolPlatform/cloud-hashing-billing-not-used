package userpaymentbalance

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

func assertUserPaymentBalance(t *testing.T, actual, expected *npool.UserPaymentBalance) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.PaymentID, expected.PaymentID)
	assert.Equal(t, actual.Amount, expected.Amount)
	assert.Equal(t, actual.CoinTypeID, expected.CoinTypeID)
	assert.Equal(t, actual.CoinUSDCurrency, expected.CoinUSDCurrency)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	userPaymentBalance := npool.UserPaymentBalance{
		AppID:           uuid.New().String(),
		UserID:          uuid.New().String(),
		PaymentID:       uuid.New().String(),
		CoinTypeID:      uuid.New().String(),
		Amount:          1.0,
		CoinUSDCurrency: 1.0,
	}

	resp, err := Create(context.Background(), &npool.CreateUserPaymentBalanceRequest{
		Info: &userPaymentBalance,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertUserPaymentBalance(t, resp.Info, &userPaymentBalance)
	}

	userPaymentBalance.ID = resp.Info.ID

	resp2, err := GetByApp(context.Background(), &npool.GetUserPaymentBalancesByAppRequest{
		AppID: userPaymentBalance.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Infos[0].ID, resp.Info.ID)
		assertUserPaymentBalance(t, resp2.Infos[0], &userPaymentBalance)
	}

	resp3, err := Get(context.Background(), &npool.GetUserPaymentBalanceRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertUserPaymentBalance(t, resp3.Info, &userPaymentBalance)
	}

	resp4, err := GetByAppUser(context.Background(), &npool.GetUserPaymentBalancesByAppUserRequest{
		AppID:  userPaymentBalance.AppID,
		UserID: userPaymentBalance.UserID,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(resp4.Infos), 0)
	}
}
