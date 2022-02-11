package goodbenefit

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

func assertGoodBenefit(t *testing.T, actual, expected *npool.GoodBenefit) {
	assert.Equal(t, actual.GoodID, expected.GoodID)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	platformSetting := npool.GoodBenefit{
		GoodID:                   uuid.New().String(),
		BenefitAccountID:         uuid.New().String(),
		PlatformOfflineAccountID: uuid.New().String(),
		UserOfflineAccountID:     uuid.New().String(),
		UserOnlineAccountID:      uuid.New().String(),
		BenefitIntervalHours:     24,
	}

	resp, err := Create(context.Background(), &npool.CreateGoodBenefitRequest{
		Info: &platformSetting,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertGoodBenefit(t, resp.Info, &platformSetting)
	}

	platformSetting.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateGoodBenefitRequest{
		Info: &platformSetting,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertGoodBenefit(t, resp1.Info, &platformSetting)
	}

	resp2, err := GetByGood(context.Background(), &npool.GetGoodBenefitByGoodRequest{
		GoodID: platformSetting.GoodID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertGoodBenefit(t, resp2.Info, &platformSetting)
	}

	resp3, err := Get(context.Background(), &npool.GetGoodBenefitRequest{
		ID: resp2.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertGoodBenefit(t, resp3.Info, &platformSetting)
	}
}
