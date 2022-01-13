package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	crud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/user-benefit" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUserBenefit(ctx context.Context, in *npool.CreateUserBenefitRequest) (*npool.CreateUserBenefitResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create user benefit error: %v", err)
		return &npool.CreateUserBenefitResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetUserBenefitsByAppUser(ctx context.Context, in *npool.GetUserBenefitsByAppUserRequest) (*npool.GetUserBenefitsByAppUserResponse, error) {
	resp, err := crud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user benefit error: %v", err)
		return &npool.GetUserBenefitsByAppUserResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetLatestUserBenefitByGoodAppUser(ctx context.Context, in *npool.GetLatestUserBenefitByGoodAppUserRequest) (*npool.GetLatestUserBenefitByGoodAppUserResponse, error) {
	resp, err := crud.GetLatestByGoodAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get latest user benefit by good app user error: %v", err)
		return &npool.GetLatestUserBenefitByGoodAppUserResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetUserBenefitByApp(ctx context.Context, in *npool.GetUserBenefitsByAppRequest) (*npool.GetUserBenefitsByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user benefit error: %v", err)
		return &npool.GetUserBenefitsByAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
