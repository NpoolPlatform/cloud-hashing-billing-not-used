// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-billing/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coin-account-info" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCoinAccount(ctx context.Context, in *npool.CreateCoinAccountRequest) (*npool.CreateCoinAccountResponse, error) {
	resp, err := coinaccountinfo.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create coin account info error: %w", err)
		return &npool.CreateCoinAccountResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetCoinAccount(ctx context.Context, in *npool.GetCoinAccountRequest) (*npool.GetCoinAccountResponse, error) {
	resp, err := coinaccountinfo.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin account info error: %w", err)
		return &npool.GetCoinAccountResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetCoinAccountsByAppUser(ctx context.Context, in *npool.GetCoinAccountsByAppUserRequest) (*npool.GetCoinAccountsByAppUserResponse, error) {
	resp, err := coinaccountinfo.GetCoinAccountsByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin account info by app user error: %w", err)
		return &npool.GetCoinAccountsByAppUserResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) DeleteCoinAccount(ctx context.Context, in *npool.DeleteCoinAccountRequest) (*npool.DeleteCoinAccountResponse, error) {
	resp, err := coinaccountinfo.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("delete coin account info error: %w", err)
		return &npool.DeleteCoinAccountResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
