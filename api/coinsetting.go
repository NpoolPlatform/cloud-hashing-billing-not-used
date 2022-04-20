package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coinsetting"
	mw "github.com/NpoolPlatform/cloud-hashing-billing/pkg/middleware/coinsetting"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCoinSetting(ctx context.Context, in *npool.CreateCoinSettingRequest) (*npool.CreateCoinSettingResponse, error) {
	resp, err := mw.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create coin setting error: %v", err)
		return &npool.CreateCoinSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateCoinSetting(ctx context.Context, in *npool.UpdateCoinSettingRequest) (*npool.UpdateCoinSettingResponse, error) {
	resp, err := mw.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update coin setting error: %v", err)
		return &npool.UpdateCoinSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetCoinSetting(ctx context.Context, in *npool.GetCoinSettingRequest) (*npool.GetCoinSettingResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin setting error: %v", err)
		return &npool.GetCoinSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetCoinSettingByCoin(ctx context.Context, in *npool.GetCoinSettingByCoinRequest) (*npool.GetCoinSettingByCoinResponse, error) {
	resp, err := crud.GetByCoin(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin setting error: %v", err)
		return &npool.GetCoinSettingByCoinResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetCoinSettings(ctx context.Context, in *npool.GetCoinSettingsRequest) (*npool.GetCoinSettingsResponse, error) {
	resp, err := crud.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin setting error: %v", err)
		return &npool.GetCoinSettingsResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
