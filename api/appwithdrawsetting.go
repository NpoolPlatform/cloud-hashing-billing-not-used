package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/appwithdrawsetting"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppWithdrawSetting(ctx context.Context, in *npool.CreateAppWithdrawSettingRequest) (*npool.CreateAppWithdrawSettingResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create app withdraw setting error: %v", err)
		return &npool.CreateAppWithdrawSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateAppWithdrawSetting(ctx context.Context, in *npool.UpdateAppWithdrawSettingRequest) (*npool.UpdateAppWithdrawSettingResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update app withdraw setting error: %v", err)
		return &npool.UpdateAppWithdrawSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppWithdrawSetting(ctx context.Context, in *npool.GetAppWithdrawSettingRequest) (*npool.GetAppWithdrawSettingResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app withdraw setting error: %v", err)
		return &npool.GetAppWithdrawSettingResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppWithdrawSettingsByApp(ctx context.Context, in *npool.GetAppWithdrawSettingsByAppRequest) (*npool.GetAppWithdrawSettingsByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app withdraw setting error: %v", err)
		return &npool.GetAppWithdrawSettingsByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppWithdrawSettingByAppCoin(ctx context.Context, in *npool.GetAppWithdrawSettingByAppCoinRequest) (*npool.GetAppWithdrawSettingByAppCoinResponse, error) {
	resp, err := crud.GetByAppCoin(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get app withdraw setting error: %v", err)
		return &npool.GetAppWithdrawSettingByAppCoinResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppWithdrawSettingsByOtherApp(ctx context.Context, in *npool.GetAppWithdrawSettingsByOtherAppRequest) (*npool.GetAppWithdrawSettingsByOtherAppResponse, error) {
	resp, err := crud.GetByApp(ctx, &npool.GetAppWithdrawSettingsByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("get app withdraw setting error: %v", err)
		return &npool.GetAppWithdrawSettingsByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAppWithdrawSettingsByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}
