package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/goodpayment"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGoodPayment(ctx context.Context, in *npool.CreateGoodPaymentRequest) (*npool.CreateGoodPaymentResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create good payment error: %v", err)
		return &npool.CreateGoodPaymentResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateGoodPayment(ctx context.Context, in *npool.UpdateGoodPaymentRequest) (*npool.UpdateGoodPaymentResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update good payment error: %v", err)
		return &npool.UpdateGoodPaymentResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetGoodPayment(ctx context.Context, in *npool.GetGoodPaymentRequest) (*npool.GetGoodPaymentResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good payment error: %v", err)
		return &npool.GetGoodPaymentResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetGoodPaymentByAccount(ctx context.Context, in *npool.GetGoodPaymentByAccountRequest) (*npool.GetGoodPaymentByAccountResponse, error) {
	resp, err := crud.GetByAccount(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good payment error: %v", err)
		return &npool.GetGoodPaymentByAccountResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetGoodPaymentsByGood(ctx context.Context, in *npool.GetGoodPaymentsByGoodRequest) (*npool.GetGoodPaymentsByGoodResponse, error) {
	resp, err := crud.GetByGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good payment error: %v", err)
		return &npool.GetGoodPaymentsByGoodResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetIdleGoodPaymentsByGood(ctx context.Context, in *npool.GetIdleGoodPaymentsByGoodRequest) (*npool.GetIdleGoodPaymentsByGoodResponse, error) {
	resp, err := crud.GetIdleByGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good payment error: %v", err)
		return &npool.GetIdleGoodPaymentsByGoodResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
