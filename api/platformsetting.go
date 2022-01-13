package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	crud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/platform-setting"     //nolint
	mw "github.com/NpoolPlatform/cloud-hashing-billing/pkg/middleware/platform-setting" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreatePlatformSetting(ctx context.Context, in *npool.CreatePlatformSettingRequest) (*npool.CreatePlatformSettingResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create platform setting error: %v", err)
		return &npool.CreatePlatformSettingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetPlatformSettingByGood(ctx context.Context, in *npool.GetPlatformSettingByGoodRequest) (*npool.GetPlatformSettingByGoodResponse, error) {
	resp, err := crud.GetByGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get platform setting error: %v", err)
		return &npool.GetPlatformSettingByGoodResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetPlatformSettingDetail(ctx context.Context, in *npool.GetPlatformSettingDetailRequest) (*npool.GetPlatformSettingDetailResponse, error) {
	resp, err := mw.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get platform setting detail error: %v", err)
		return &npool.GetPlatformSettingDetailResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}
