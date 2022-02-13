package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/platform-setting"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreatePlatformSetting(ctx context.Context, in *npool.CreatePlatformSettingRequest) (*npool.CreatePlatformSettingResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create platform setting error: %v", err)
		return &npool.CreatePlatformSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdatePlatformSetting(ctx context.Context, in *npool.UpdatePlatformSettingRequest) (*npool.UpdatePlatformSettingResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update platform setting error: %v", err)
		return &npool.UpdatePlatformSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetPlatformSetting(ctx context.Context, in *npool.GetPlatformSettingRequest) (*npool.GetPlatformSettingResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get platform setting error: %v", err)
		return &npool.GetPlatformSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
