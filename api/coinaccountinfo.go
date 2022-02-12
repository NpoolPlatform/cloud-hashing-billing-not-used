// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coin-account-info" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCoinAccount(ctx context.Context, in *npool.CreateCoinAccountRequest) (*npool.CreateCoinAccountResponse, error) {
	resp, err := coinaccountinfo.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create coin account info error: %v", err)
		return &npool.CreateCoinAccountResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetCoinAccount(ctx context.Context, in *npool.GetCoinAccountRequest) (*npool.GetCoinAccountResponse, error) {
	resp, err := coinaccountinfo.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin account info error: %v", err)
		return &npool.GetCoinAccountResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetCoinAccountByCoinAddress(ctx context.Context, in *npool.GetCoinAccountByCoinAddressRequest) (*npool.GetCoinAccountByCoinAddressResponse, error) {
	resp, err := coinaccountinfo.GetByCoinAddress(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin account info by coin address error: %v", err)
		return &npool.GetCoinAccountByCoinAddressResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetCoinAccounts(ctx context.Context, in *npool.GetCoinAccountsRequest) (*npool.GetCoinAccountsResponse, error) {
	resp, err := coinaccountinfo.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin accounts info error: %v", err)
		return &npool.GetCoinAccountsResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) DeleteCoinAccount(ctx context.Context, in *npool.DeleteCoinAccountRequest) (*npool.DeleteCoinAccountResponse, error) {
	resp, err := coinaccountinfo.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("delete coin account info error: %v", err)
		return &npool.DeleteCoinAccountResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
