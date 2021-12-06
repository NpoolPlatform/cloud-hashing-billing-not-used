package platformbenefit

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-billing/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/test-init" //nolint

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coin-account-info" //nolint
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/platform-benefit"  //nolint

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

func TestGet(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	account1 := npool.CoinAccountInfo{
		CoinTypeID:  uuid.New().String(),
		Address:     uuid.New().String(),
		GeneratedBy: "user",
		UsedFor:     "withdraw",
		AppID:       uuid.New().String(),
		UserID:      uuid.New().String(),
	}

	resp1, err := coinaccountinfo.Create(context.Background(), &npool.CreateCoinAccountRequest{
		Info: &account1,
	})
	assert.Nil(t, err)

	benefit := npool.PlatformBenefit{
		GoodID:               uuid.New().String(),
		BenefitAccountID:     resp1.Info.ID,
		Amount:               0.13,
		LastBenefitTimestamp: uint32(time.Now().Unix()) + 1000,
		ChainTransactionID:   "adfjklasjfdlksajf",
	}
	resp2, err := platformbenefit.Create(context.Background(), &npool.CreatePlatformBenefitRequest{
		Info: &benefit,
	})
	assert.Nil(t, err)

	resp3, err := Get(context.Background(), &npool.GetPlatformBenefitDetailRequest{
		ID: resp2.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Detail.ID, resp2.Info.ID)
		assert.Equal(t, resp3.Detail.GoodID, resp2.Info.GoodID)

		assert.Equal(t, resp3.Detail.BenefitAddress.ID, resp1.Info.ID)
		assert.Equal(t, resp3.Detail.BenefitAddress.CoinTypeID, resp1.Info.CoinTypeID)
		assert.Equal(t, resp3.Detail.BenefitAddress.Address, resp1.Info.Address)
		assert.Equal(t, resp3.Detail.BenefitAddress.GeneratedBy, resp1.Info.GeneratedBy)
		assert.Equal(t, resp3.Detail.BenefitAddress.UsedFor, resp1.Info.UsedFor)
		assert.Equal(t, resp3.Detail.BenefitAddress.Idle, resp1.Info.Idle)
		assert.Equal(t, resp3.Detail.BenefitAddress.AppID, resp1.Info.AppID)
		assert.Equal(t, resp3.Detail.BenefitAddress.UserID, resp1.Info.UserID)

		assert.Equal(t, resp3.Detail.Amount, resp2.Info.Amount)
		assert.Equal(t, resp3.Detail.ChainTransactionID, resp2.Info.ChainTransactionID)
	}
}
