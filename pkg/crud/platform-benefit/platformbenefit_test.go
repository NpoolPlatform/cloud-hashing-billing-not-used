package platformbenefit

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

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

func assertPlatformBenefit(t *testing.T, actual, expected *npool.PlatformBenefit) {
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.BenefitAccountID, expected.BenefitAccountID)
	assert.Equal(t, actual.Amount, expected.Amount)
	assert.Equal(t, actual.LastBenefitTimestamp, expected.LastBenefitTimestamp)
	assert.Equal(t, actual.ChainTransactionID, expected.ChainTransactionID)
	assert.Equal(t, actual.PlatformTransactionID, expected.PlatformTransactionID)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	platformBenefit := npool.PlatformBenefit{
		GoodID:                uuid.New().String(),
		BenefitAccountID:      uuid.New().String(),
		Amount:                0.13,
		LastBenefitTimestamp:  uint32(time.Now().Unix()),
		ChainTransactionID:    uuid.New().String(),
		PlatformTransactionID: uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreatePlatformBenefitRequest{
		Info: &platformBenefit,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertPlatformBenefit(t, resp.Info, &platformBenefit)
	}

	resp1, err := GetByGood(context.Background(), &npool.GetPlatformBenefitsByGoodRequest{
		GoodID: platformBenefit.GoodID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp1.Infos), 1)
	}

	resp2, err := Get(context.Background(), &npool.GetPlatformBenefitRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertPlatformBenefit(t, resp2.Info, &platformBenefit)
	}

	resp3, err := GetLatestByGood(context.Background(), &npool.GetLatestPlatformBenefitByGoodRequest{
		GoodID: platformBenefit.GoodID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertPlatformBenefit(t, resp3.Info, &platformBenefit)
	}

	platformBenefit.LastBenefitTimestamp++
	platformBenefit.ChainTransactionID = uuid.New().String()

	time.Sleep(2 * time.Second)

	resp4, err := Create(context.Background(), &npool.CreatePlatformBenefitRequest{
		Info: &platformBenefit,
	})
	assert.Nil(t, err)

	resp5, err := GetLatestByGood(context.Background(), &npool.GetLatestPlatformBenefitByGoodRequest{
		GoodID: platformBenefit.GoodID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp5.Info.ID, resp4.Info.ID)
		assertPlatformBenefit(t, resp5.Info, &platformBenefit)
	}
}
