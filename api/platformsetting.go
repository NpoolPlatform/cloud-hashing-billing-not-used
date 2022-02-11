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
		return &npool.CreatePlatformSettingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) CreatePlatformSettingForOtherApp(ctx context.Context, in *npool.CreatePlatformSettingForOtherAppRequest) (*npool.CreatePlatformSettingForOtherAppResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()

	resp, err := crud.Create(ctx, &npool.CreatePlatformSettingRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorf("create platform setting for other app error: %v", err)
		return &npool.CreatePlatformSettingForOtherAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return &npool.CreatePlatformSettingForOtherAppResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) GetPlatformSetting(ctx context.Context, in *npool.GetPlatformSettingRequest) (*npool.GetPlatformSettingResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get platform setting error: %v", err)
		return &npool.GetPlatformSettingResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetPlatformSettingByApp(ctx context.Context, in *npool.GetPlatformSettingByAppRequest) (*npool.GetPlatformSettingByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get platform setting error: %v", err)
		return &npool.GetPlatformSettingByAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return resp, nil
}

func (s *Server) GetPlatformSettingByOtherApp(ctx context.Context, in *npool.GetPlatformSettingByOtherAppRequest) (*npool.GetPlatformSettingByOtherAppResponse, error) {
	resp, err := crud.GetByApp(ctx, &npool.GetPlatformSettingByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("get platform setting error: %v", err)
		return &npool.GetPlatformSettingByOtherAppResponse{}, status.Error(codes.Internal, "internal server error")
	}
	return &npool.GetPlatformSettingByOtherAppResponse{
		Info: resp.Info,
	}, nil
}
