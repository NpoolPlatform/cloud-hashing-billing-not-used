package userbenefit

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

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

func assertUserBenefit(t *testing.T, actual, expected *npool.UserBenefit) {
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.Amount, expected.Amount)
	assert.Equal(t, actual.LastBenefitTimestamp, expected.LastBenefitTimestamp)
	assert.Equal(t, actual.OrderID, expected.OrderID)
}

func TestCRUD(t *testing.T) { //nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	userBenefit := npool.UserBenefit{
		GoodID:               uuid.New().String(),
		AppID:                uuid.New().String(),
		UserID:               uuid.New().String(),
		Amount:               0.13,
		LastBenefitTimestamp: uint32(time.Now().Unix()),
		OrderID:              uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateUserBenefitRequest{
		Info: &userBenefit,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertUserBenefit(t, resp.Info, &userBenefit)
	}

	resp1, err := GetByAppUser(context.Background(), &npool.GetUserBenefitsByAppUserRequest{
		AppID:  userBenefit.AppID,
		UserID: userBenefit.UserID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp1.Infos), 1)
	}

	resp2, err := GetByApp(context.Background(), &npool.GetUserBenefitsByAppRequest{
		AppID: userBenefit.AppID,
	})
	if assert.Nil(t, err) {
		assert.Positive(t, len(resp2.Infos))
	}

	resp3, err := GetLatestByGood(context.Background(), &npool.GetLatestUserBenefitByGoodRequest{
		GoodID: userBenefit.GoodID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, userBenefit.LastBenefitTimestamp, resp3.Info.LastBenefitTimestamp)
		assertUserBenefit(t, &userBenefit, resp3.Info)
	}
}
