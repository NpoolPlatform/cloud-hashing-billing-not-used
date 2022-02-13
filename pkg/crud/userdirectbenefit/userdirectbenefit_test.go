package userdirectbenefit

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

func assertUserDirectBenefit(t *testing.T, actual, expected *npool.UserDirectBenefit) {
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.CoinTypeID, expected.CoinTypeID)
	assert.Equal(t, actual.AccountID, expected.AccountID)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	userDirectBenefit := npool.UserDirectBenefit{
		AppID:      uuid.New().String(),
		UserID:     uuid.New().String(),
		CoinTypeID: uuid.New().String(),
		AccountID:  uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateUserDirectBenefitRequest{
		Info: &userDirectBenefit,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{}.String())
		assertUserDirectBenefit(t, resp.Info, &userDirectBenefit)
	}

	userDirectBenefit.ID = resp.Info.ID

	resp1, err := Update(context.Background(), &npool.UpdateUserDirectBenefitRequest{
		Info: &userDirectBenefit,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertUserDirectBenefit(t, resp1.Info, &userDirectBenefit)
	}

	resp2, err := GetByAccount(context.Background(), &npool.GetUserDirectBenefitByAccountRequest{
		AccountID: userDirectBenefit.AccountID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp2.Info.ID, resp.Info.ID)
		assertUserDirectBenefit(t, resp2.Info, &userDirectBenefit)
	}

	resp3, err := Get(context.Background(), &npool.GetUserDirectBenefitRequest{
		ID: resp2.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp3.Info.ID, resp.Info.ID)
		assertUserDirectBenefit(t, resp3.Info, &userDirectBenefit)
	}

	resp4, err := GetByAppUser(context.Background(), &npool.GetUserDirectBenefitsByAppUserRequest{
		AppID:  userDirectBenefit.AppID,
		UserID: userDirectBenefit.UserID,
	})
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(resp4.Infos), 0)
	}
}
