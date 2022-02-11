package goodpayment

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

func assertGoodPayment(t *testing.T, actual, expected *npool.GoodPayment) {
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.PaymentCoinTypeID, expected.PaymentCoinTypeID)
	assert.Equal(t, actual.AccountID, expected.AccountID)
	assert.Equal(t, actual.Idle, expected.Idle)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	payment := npool.GoodPayment{
		GoodID:            uuid.New().String(),
		PaymentCoinTypeID: uuid.New().String(),
		AccountID:         uuid.New().String(),
		Idle:              true,
	}

	resp, err := Create(context.Background(), &npool.CreateGoodPaymentRequest{
		Info: &payment,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertGoodPayment(t, resp.Info, &payment)
	}

	payment.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateGoodPaymentRequest{
		Info: &payment,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertGoodPayment(t, resp1.Info, &payment)
	}

	resp2, err := GetByAccount(context.Background(), &npool.GetGoodPaymentByAccountRequest{
		AccountID: payment.AccountID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertGoodPayment(t, resp2.Info, &payment)
	}

	resp3, err := Get(context.Background(), &npool.GetGoodPaymentRequest{
		ID: resp2.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertGoodPayment(t, resp3.Info, &payment)
	}

	resp4, err := GetByGood(context.Background(), &npool.GetGoodPaymentsByGoodRequest{
		GoodID: payment.GoodID,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(resp4.Infos), 0)
	}
}
