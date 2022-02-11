package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/goodsetting"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGoodSetting(ctx context.Context, in *npool.CreateGoodSettingRequest) (*npool.CreateGoodSettingResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create good setting error: %v", err)
		return &npool.CreateGoodSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateGoodSetting(ctx context.Context, in *npool.UpdateGoodSettingRequest) (*npool.UpdateGoodSettingResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update good setting error: %v", err)
		return &npool.UpdateGoodSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetGoodSetting(ctx context.Context, in *npool.GetGoodSettingRequest) (*npool.GetGoodSettingResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good setting error: %v", err)
		return &npool.GetGoodSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetGoodSettingByGood(ctx context.Context, in *npool.GetGoodSettingByGoodRequest) (*npool.GetGoodSettingByGoodResponse, error) {
	resp, err := crud.GetByGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good setting error: %v", err)
		return &npool.GetGoodSettingByGoodResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
