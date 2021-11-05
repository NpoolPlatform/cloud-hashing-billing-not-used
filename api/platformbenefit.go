package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-billing/message/npool"

	crud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/platform-benefit"     //nolint
	mw "github.com/NpoolPlatform/cloud-hashing-billing/pkg/middleware/platform-benefit" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Create(ctx context.Context, in *npool.CreatePlatformBenefitRequest) (*npool.CreatePlatformBenefitResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create platform benefit error: %w", err)
		return &npool.CreatePlatformBenefitResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func GetByGood(ctx context.Context, in *npool.GetPlatformBenefitsByGoodRequest) (*npool.GetPlatformBenefitsByGoodResponse, error) {
	resp, err := crud.GetByGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create get good platform benefit error: %w", err)
		return &npool.GetPlatformBenefitsByGoodResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func Get(ctx context.Context, in *npool.GetPlatformBenefitRequest) (*npool.GetPlatformBenefitResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create platform benefit error: %w", err)
		return &npool.GetPlatformBenefitResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func GetDetail(ctx context.Context, in *npool.GetPlatformBenefitDetailRequest) (*npool.GetPlatformBenefitDetailResponse, error) {
	resp, err := mw.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create platform benefit detail error: %w", err)
		return &npool.GetPlatformBenefitDetailResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
