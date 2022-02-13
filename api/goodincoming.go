// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/goodincoming"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGoodIncoming(ctx context.Context, in *npool.CreateGoodIncomingRequest) (*npool.CreateGoodIncomingResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create good incoming error: %v", err)
		return &npool.CreateGoodIncomingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateGoodIncoming(ctx context.Context, in *npool.UpdateGoodIncomingRequest) (*npool.UpdateGoodIncomingResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update good incoming error: %v", err)
		return &npool.UpdateGoodIncomingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetGoodIncoming(ctx context.Context, in *npool.GetGoodIncomingRequest) (*npool.GetGoodIncomingResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good incoming error: %v", err)
		return &npool.GetGoodIncomingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetGoodIncomingsByGood(ctx context.Context, in *npool.GetGoodIncomingsByGoodRequest) (*npool.GetGoodIncomingsByGoodResponse, error) {
	resp, err := crud.GetByGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good incomings error: %v", err)
		return &npool.GetGoodIncomingsByGoodResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetGoodIncomingByGoodCoin(ctx context.Context, in *npool.GetGoodIncomingByGoodCoinRequest) (*npool.GetGoodIncomingByGoodCoinResponse, error) {
	resp, err := crud.GetByGoodCoin(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good incoming error: %v", err)
		return &npool.GetGoodIncomingByGoodCoinResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
