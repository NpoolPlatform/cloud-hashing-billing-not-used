package api

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/cloud-hashing-billing/message/npool"
)

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

func TestCoinAccountTransactionCRUD(t *testing.T) { //nolint
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
		State:                 "wait",
	}
	firstCreateInfo := npool.CreateCoinAccountTransactionResponse{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateCoinAccountTransactionRequest{
			Info: &coinAccountTransaction,
		}).
		Post("http://localhost:34759/v1/create/coin/account/transaction")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.UUID{}.String())
			assertCoinAccountTransaction(t, firstCreateInfo.Info, &coinAccountTransaction)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetCoinAccountTransactionRequest{
			ID: firstCreateInfo.Info.ID,
		}).
		Post("http://localhost:34759/v1/get/coin/account/transaction")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetCoinAccountTransactionResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assertCoinAccountTransaction(t, info.Info, &coinAccountTransaction)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetCoinAccountTransactionsByCoinAccountRequest{
			CoinTypeID: firstCreateInfo.Info.CoinTypeID,
			AddressID:  firstCreateInfo.Info.FromAddressID,
		}).
		Post("http://localhost:34759/v1/get/coin/account/transactions/by/coin/account")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetCoinAccountTransactionsByCoinAccountResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, len(info.Infos), 1)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetCoinAccountTransactionsByCoinAccountRequest{
			CoinTypeID: firstCreateInfo.Info.CoinTypeID,
			AddressID:  firstCreateInfo.Info.ToAddressID,
		}).
		Post("http://localhost:34759/v1/get/coin/account/transactions/by/coin/account")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetCoinAccountTransactionsByCoinAccountResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, len(info.Infos), 1)
		}
	}

	coinAccountTransaction.State = "paying"
	coinAccountTransaction.ID = firstCreateInfo.Info.ID

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.UpdateCoinAccountTransactionRequest{
			Info: &coinAccountTransaction,
		}).
		Post("http://localhost:34759/v1/update/coin/account/transaction")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.UpdateCoinAccountTransactionResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assertCoinAccountTransaction(t, info.Info, &coinAccountTransaction)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.DeleteCoinAccountTransactionRequest{
			ID: firstCreateInfo.Info.ID,
		}).
		Post("http://localhost:34759/v1/delete/coin/account/transaction")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.DeleteCoinAccountTransactionResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assertCoinAccountTransaction(t, info.Info, &coinAccountTransaction)
		}
	}
}
