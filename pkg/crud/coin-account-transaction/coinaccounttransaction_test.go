package coinaccounttransaction

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/const"     //nolint
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

func assertCoinAccountTransaction(t *testing.T, actual, expected *npool.CoinAccountTransaction) {
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.CoinTypeID, expected.CoinTypeID)
	assert.Equal(t, actual.FromAddressID, expected.FromAddressID)
	assert.Equal(t, actual.ToAddressID, expected.ToAddressID)
	assert.Equal(t, actual.Amount, expected.Amount)
	assert.Equal(t, actual.Message, expected.Message)
	assert.Equal(t, actual.ChainTransactionID, expected.ChainTransactionID)
	assert.Equal(t, actual.PlatformTransactionID, expected.PlatformTransactionID)
	assert.Equal(t, actual.State, expected.State)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	coinAccountTransaction := npool.CoinAccountTransaction{
		UserID:                uuid.New().String(),
		AppID:                 uuid.New().String(),
		CoinTypeID:            uuid.New().String(),
		FromAddressID:         uuid.New().String(),
		ToAddressID:           uuid.New().String(),
		PlatformTransactionID: uuid.New().String(),
		Amount:                1.3,
		Message:               "for transaction test",
		State:                 constant.CoinTransactionStateCreated,
	}

	resp, err := Create(context.Background(), &npool.CreateCoinAccountTransactionRequest{
		Info: &coinAccountTransaction,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertCoinAccountTransaction(t, resp.Info, &coinAccountTransaction)
	}

	resp1, err := Get(context.Background(), &npool.GetCoinAccountTransactionRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertCoinAccountTransaction(t, resp1.Info, &coinAccountTransaction)
	}

	resp2, err := GetCoinAccountTransactionsByCoinAccount(context.Background(), &npool.GetCoinAccountTransactionsByCoinAccountRequest{
		CoinTypeID: coinAccountTransaction.CoinTypeID,
		AddressID:  coinAccountTransaction.FromAddressID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp2.Infos), 1)
	}

	resp2, err = GetCoinAccountTransactionsByCoinAccount(context.Background(), &npool.GetCoinAccountTransactionsByCoinAccountRequest{
		CoinTypeID: coinAccountTransaction.CoinTypeID,
		AddressID:  coinAccountTransaction.ToAddressID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp2.Infos), 1)
	}

	resp3, err := GetCoinAccountTransactionsByState(context.Background(), &npool.GetCoinAccountTransactionsByStateRequest{
		State: constant.CoinTransactionStateCreated,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(resp3.Infos), 0)
	}

	resp4, err := GetCoinAccountTransactionsByCoin(context.Background(), &npool.GetCoinAccountTransactionsByCoinRequest{
		CoinTypeID: coinAccountTransaction.CoinTypeID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp4.Infos), 1)
	}

	coinAccountTransaction.State = "paying"
	coinAccountTransaction.ID = resp.Info.ID

	resp5, err := Update(context.Background(), &npool.UpdateCoinAccountTransactionRequest{
		Info: &coinAccountTransaction,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp5.Info.ID, resp.Info.ID)
		assertCoinAccountTransaction(t, resp5.Info, &coinAccountTransaction)
	}

	resp6, err := Delete(context.Background(), &npool.DeleteCoinAccountTransactionRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp6.Info.ID, resp.Info.ID)
		assertCoinAccountTransaction(t, resp6.Info, &coinAccountTransaction)
	}
}
