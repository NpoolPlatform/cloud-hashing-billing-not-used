// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	crud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/platform-benefit"     //nolint
	mw "github.com/NpoolPlatform/cloud-hashing-billing/pkg/middleware/platform-benefit" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreatePlatformBenefit(ctx context.Context, in *npool.CreatePlatformBenefitRequest) (*npool.CreatePlatformBenefitResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create platform benefit error: %v", err)
		return &npool.CreatePlatformBenefitResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetPlatformBenefitsByGood(ctx context.Context, in *npool.GetPlatformBenefitsByGoodRequest) (*npool.GetPlatformBenefitsByGoodResponse, error) {
	resp, err := crud.GetByGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good platform benefit error: %v", err)
		return &npool.GetPlatformBenefitsByGoodResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetPlatformBenefits(ctx context.Context, in *npool.GetPlatformBenefitsRequest) (*npool.GetPlatformBenefitsResponse, error) {
	resp, err := crud.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good platform benefit error: %v", err)
		return &npool.GetPlatformBenefitsResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetLatestPlatformBenefitByGood(ctx context.Context, in *npool.GetLatestPlatformBenefitByGoodRequest) (*npool.GetLatestPlatformBenefitByGoodResponse, error) {
	resp, err := crud.GetLatestByGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good latest platform benefit error: %v", err)
		return &npool.GetLatestPlatformBenefitByGoodResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetPlatformBenefit(ctx context.Context, in *npool.GetPlatformBenefitRequest) (*npool.GetPlatformBenefitResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get platform benefit error: %v", err)
		return &npool.GetPlatformBenefitResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetPlatformBenefitDetail(ctx context.Context, in *npool.GetPlatformBenefitDetailRequest) (*npool.GetPlatformBenefitDetailResponse, error) {
	resp, err := mw.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get platform benefit detail error: %v", err)
		return &npool.GetPlatformBenefitDetailResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
