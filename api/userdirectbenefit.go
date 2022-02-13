package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/userdirectbenefit"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUserDirectBenefit(ctx context.Context, in *npool.CreateUserDirectBenefitRequest) (*npool.CreateUserDirectBenefitResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create user direct benefit error: %v", err)
		return &npool.CreateUserDirectBenefitResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateUserDirectBenefit(ctx context.Context, in *npool.UpdateUserDirectBenefitRequest) (*npool.UpdateUserDirectBenefitResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update user direct benefit error: %v", err)
		return &npool.UpdateUserDirectBenefitResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserDirectBenefit(ctx context.Context, in *npool.GetUserDirectBenefitRequest) (*npool.GetUserDirectBenefitResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user direct benefit error: %v", err)
		return &npool.GetUserDirectBenefitResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserDirectBenefitByAccount(ctx context.Context, in *npool.GetUserDirectBenefitByAccountRequest) (*npool.GetUserDirectBenefitByAccountResponse, error) {
	resp, err := crud.GetByAccount(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user direct benefit error: %v", err)
		return &npool.GetUserDirectBenefitByAccountResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserDirectBenefitsByAppUser(ctx context.Context, in *npool.GetUserDirectBenefitsByAppUserRequest) (*npool.GetUserDirectBenefitsByAppUserResponse, error) {
	resp, err := crud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user direct benefit error: %v", err)
		return &npool.GetUserDirectBenefitsByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserDirectBenefitsByOtherAppUser(ctx context.Context, in *npool.GetUserDirectBenefitsByOtherAppUserRequest) (*npool.GetUserDirectBenefitsByOtherAppUserResponse, error) {
	resp, err := crud.GetByAppUser(ctx, &npool.GetUserDirectBenefitsByAppUserRequest{
		AppID:  in.GetTargetAppID(),
		UserID: in.GetTargetUserID(),
	})
	if err != nil {
		logger.Sugar().Errorf("get user direct benefit error: %v", err)
		return &npool.GetUserDirectBenefitsByOtherAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetUserDirectBenefitsByOtherAppUserResponse{
		Infos: resp.Infos,
	}, nil
}
