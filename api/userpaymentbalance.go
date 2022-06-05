// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/userpaymentbalance"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUserPaymentBalance(ctx context.Context, in *npool.CreateUserPaymentBalanceRequest) (*npool.CreateUserPaymentBalanceResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create user payment balance error: %v", err)
		return &npool.CreateUserPaymentBalanceResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateUserPaymentBalanceForOtherAppUser(ctx context.Context, in *npool.CreateUserPaymentBalanceForOtherAppUserRequest) (*npool.CreateUserPaymentBalanceForOtherAppUserResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()
	info.UserID = in.GetTargetUserID()

	resp, err := crud.Create(ctx, &npool.CreateUserPaymentBalanceRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorf("create user payment balance error: %v", err)
		return &npool.CreateUserPaymentBalanceForOtherAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.CreateUserPaymentBalanceForOtherAppUserResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) GetUserPaymentBalance(ctx context.Context, in *npool.GetUserPaymentBalanceRequest) (*npool.GetUserPaymentBalanceResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user payment balance error: %v", err)
		return &npool.GetUserPaymentBalanceResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserPaymentBalancesByApp(ctx context.Context, in *npool.GetUserPaymentBalancesByAppRequest) (*npool.GetUserPaymentBalancesByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user payment balance error: %v", err)
		return &npool.GetUserPaymentBalancesByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserPaymentBalancesByOtherApp(ctx context.Context, in *npool.GetUserPaymentBalancesByOtherAppRequest) (*npool.GetUserPaymentBalancesByOtherAppResponse, error) {
	resp, err := crud.GetByApp(ctx, &npool.GetUserPaymentBalancesByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("get user payment balance error: %v", err)
		return &npool.GetUserPaymentBalancesByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetUserPaymentBalancesByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetUserPaymentBalancesByAppUser(ctx context.Context, in *npool.GetUserPaymentBalancesByAppUserRequest) (*npool.GetUserPaymentBalancesByAppUserResponse, error) {
	resp, err := crud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user payment balance error: %v", err)
		return &npool.GetUserPaymentBalancesByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
