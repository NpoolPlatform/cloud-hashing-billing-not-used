// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/cloud-hashing-billing/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coin-account-transaction"          //nolint
	mw "github.com/NpoolPlatform/cloud-hashing-billing/pkg/middleware/coin-account-transaction" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCoinAccountTransaction(ctx context.Context, in *npool.CreateCoinAccountTransactionRequest) (*npool.CreateCoinAccountTransactionResponse, error) {
	resp, err := coinaccounttransaction.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create coin account transaction error: %v", err)
		return &npool.CreateCoinAccountTransactionResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetCoinAccountTransaction(ctx context.Context, in *npool.GetCoinAccountTransactionRequest) (*npool.GetCoinAccountTransactionResponse, error) {
	resp, err := coinaccounttransaction.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin account transaction error: %v", err)
		return &npool.GetCoinAccountTransactionResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetCoinAccountTransactionsByCoinAccount(ctx context.Context, in *npool.GetCoinAccountTransactionsByCoinAccountRequest) (*npool.GetCoinAccountTransactionsByCoinAccountResponse, error) {
	resp, err := coinaccounttransaction.GetCoinAccountTransactionsByCoinAccount(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin account transaction by coin account error: %v", err)
		return &npool.GetCoinAccountTransactionsByCoinAccountResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetCoinAccountTransactionsByState(ctx context.Context, in *npool.GetCoinAccountTransactionsByStateRequest) (*npool.GetCoinAccountTransactionsByStateResponse, error) {
	resp, err := coinaccounttransaction.GetCoinAccountTransactionsByState(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin account transaction by state error: %v", err)
		return &npool.GetCoinAccountTransactionsByStateResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetCoinAccountTransactionsByCoin(ctx context.Context, in *npool.GetCoinAccountTransactionsByCoinRequest) (*npool.GetCoinAccountTransactionsByCoinResponse, error) {
	resp, err := coinaccounttransaction.GetCoinAccountTransactionsByCoin(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin account transaction by coin error: %v", err)
		return &npool.GetCoinAccountTransactionsByCoinResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) UpdateCoinAccountTransaction(ctx context.Context, in *npool.UpdateCoinAccountTransactionRequest) (*npool.UpdateCoinAccountTransactionResponse, error) {
	resp, err := coinaccounttransaction.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update coin account transaction error: %v", err)
		return &npool.UpdateCoinAccountTransactionResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetCoinAccountTransactionDetail(ctx context.Context, in *npool.GetCoinAccountTransactionDetailRequest) (*npool.GetCoinAccountTransactionDetailResponse, error) {
	resp, err := mw.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin account transaction detail error: %v", err)
		return &npool.GetCoinAccountTransactionDetailResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) DeleteCoinAccountTransaction(ctx context.Context, in *npool.DeleteCoinAccountTransactionRequest) (*npool.DeleteCoinAccountTransactionResponse, error) {
	resp, err := coinaccounttransaction.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("delete coin account transaction error: %v", err)
		return &npool.DeleteCoinAccountTransactionResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
