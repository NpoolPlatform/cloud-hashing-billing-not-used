package platformbenefit

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-billing/message/npool"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/test-init" //nolint

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
	assert.Equal(t, actual.ChainTransactionID, expected.ChainTransactionID)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	platformBenefit := npool.PlatformBenefit{
		GoodID:             uuid.New().String(),
		BenefitAccountID:   uuid.New().String(),
		Amount:             0.13,
		ChainTransactionID: uuid.New().String(),
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
}
