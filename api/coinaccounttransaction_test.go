package api

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/cloud-hashing-billing/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/test-init" //nolint

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coin-account-info"        //nolint
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coin-account-transaction" //nolint
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
		Post("http://localhost:50030/v1/create/coin/account/transaction")
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
		Post("http://localhost:50030/v1/get/coin/account/transaction")
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
		Post("http://localhost:50030/v1/get/coin/account/transactions/by/coin/account")
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
		Post("http://localhost:50030/v1/get/coin/account/transactions/by/coin/account")
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
		Post("http://localhost:50030/v1/update/coin/account/transaction")
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
		Post("http://localhost:50030/v1/delete/coin/account/transaction")
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

func TestGetCoinAccountTransactionDetail(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	coinTypeID := uuid.New().String()

	account1 := npool.CoinAccountInfo{
		CoinTypeID:  coinTypeID,
		Address:     uuid.New().String(),
		GeneratedBy: "user",
		UsedFor:     "withdraw",
		AppID:       uuid.New().String(),
		UserID:      uuid.New().String(),
	}
	account2 := npool.CoinAccountInfo{
		CoinTypeID:  coinTypeID,
		Address:     uuid.New().String(),
		GeneratedBy: "platform",
		UsedFor:     "benefit",
		AppID:       uuid.UUID{}.String(),
		UserID:      uuid.UUID{}.String(),
	}

	resp1, err := coinaccountinfo.Create(context.Background(), &npool.CreateCoinAccountRequest{
		Info: &account1,
	})
	assert.Nil(t, err)

	resp2, err := coinaccountinfo.Create(context.Background(), &npool.CreateCoinAccountRequest{
		Info: &account2,
	})
	assert.Nil(t, err)

	transaction := npool.CoinAccountTransaction{
		UserID:                resp1.Info.UserID,
		AppID:                 resp1.Info.AppID,
		FromAddressID:         resp1.Info.ID,
		ToAddressID:           resp2.Info.ID,
		CoinTypeID:            coinTypeID,
		Amount:                0.13,
		Message:               "test transaction",
		ChainTransactionID:    "adfjklasjfdlksajf",
		PlatformTransactionID: uuid.New().String(),
	}
	resp3, err := coinaccounttransaction.Create(context.Background(), &npool.CreateCoinAccountTransactionRequest{
		Info: &transaction,
	})
	assert.Nil(t, err)

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetCoinAccountTransactionDetailRequest{
			ID: resp3.Info.ID,
		}).
		Post("http://localhost:50030/v1/get/coin/account/transaction/detail")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetCoinAccountTransactionDetailResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Detail.ID, resp3.Info.ID)
			assert.Equal(t, info.Detail.UserID, resp3.Info.UserID)
			assert.Equal(t, info.Detail.AppID, resp3.Info.AppID)

			assert.Equal(t, info.Detail.FromAddress.ID, resp1.Info.ID)
			assert.Equal(t, info.Detail.FromAddress.CoinTypeID, resp1.Info.CoinTypeID)
			assert.Equal(t, info.Detail.FromAddress.Address, resp1.Info.Address)
			assert.Equal(t, info.Detail.FromAddress.GeneratedBy, resp1.Info.GeneratedBy)
			assert.Equal(t, info.Detail.FromAddress.UsedFor, resp1.Info.UsedFor)
			assert.Equal(t, info.Detail.FromAddress.Idle, resp1.Info.Idle)
			assert.Equal(t, info.Detail.FromAddress.AppID, resp1.Info.AppID)
			assert.Equal(t, info.Detail.FromAddress.UserID, resp1.Info.UserID)

			assert.Equal(t, info.Detail.ToAddress.ID, resp2.Info.ID)
			assert.Equal(t, info.Detail.ToAddress.CoinTypeID, resp2.Info.CoinTypeID)
			assert.Equal(t, info.Detail.ToAddress.Address, resp2.Info.Address)
			assert.Equal(t, info.Detail.ToAddress.GeneratedBy, resp2.Info.GeneratedBy)
			assert.Equal(t, info.Detail.ToAddress.UsedFor, resp2.Info.UsedFor)
			assert.Equal(t, info.Detail.ToAddress.Idle, resp2.Info.Idle)
			assert.Equal(t, info.Detail.ToAddress.AppID, resp2.Info.AppID)
			assert.Equal(t, info.Detail.ToAddress.UserID, resp2.Info.UserID)

			assert.Equal(t, info.Detail.Amount, resp3.Info.Amount)
			assert.Equal(t, info.Detail.Message, resp3.Info.Message)
			assert.Equal(t, info.Detail.ChainTransactionID, resp3.Info.ChainTransactionID)
			assert.Equal(t, info.Detail.PlatformTransactionID, resp3.Info.PlatformTransactionID)
			assert.Equal(t, info.Detail.State, resp3.Info.State)
		}
	}
}
