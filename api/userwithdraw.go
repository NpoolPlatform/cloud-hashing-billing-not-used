// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/userwithdraw"
	mw "github.com/NpoolPlatform/cloud-hashing-billing/pkg/middleware/withdraw"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUserWithdraw(ctx context.Context, in *npool.CreateUserWithdrawRequest) (*npool.CreateUserWithdrawResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create user withdraw error: %v", err)
		return &npool.CreateUserWithdrawResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateUserWithdraw(ctx context.Context, in *npool.UpdateUserWithdrawRequest) (*npool.UpdateUserWithdrawResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update user withdraw error: %v", err)
		return &npool.UpdateUserWithdrawResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserWithdraw(ctx context.Context, in *npool.GetUserWithdrawRequest) (*npool.GetUserWithdrawResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user withdraw error: %v", err)
		return &npool.GetUserWithdrawResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserWithdrawByAccount(ctx context.Context, in *npool.GetUserWithdrawByAccountRequest) (*npool.GetUserWithdrawByAccountResponse, error) {
	resp, err := crud.GetByAccount(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user withdraw error: %v", err)
		return &npool.GetUserWithdrawByAccountResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserWithdrawsByAppUser(ctx context.Context, in *npool.GetUserWithdrawsByAppUserRequest) (*npool.GetUserWithdrawsByAppUserResponse, error) {
	resp, err := crud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user withdraw error: %v", err)
		return &npool.GetUserWithdrawsByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserWithdrawsByAppUserCoin(ctx context.Context, in *npool.GetUserWithdrawsByAppUserCoinRequest) (*npool.GetUserWithdrawsByAppUserCoinResponse, error) {
	resp, err := crud.GetByAppUserCoin(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user withdraw error: %v", err)
		return &npool.GetUserWithdrawsByAppUserCoinResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserWithdrawsByOtherAppUser(ctx context.Context, in *npool.GetUserWithdrawsByOtherAppUserRequest) (*npool.GetUserWithdrawsByOtherAppUserResponse, error) {
	resp, err := crud.GetByAppUser(ctx, &npool.GetUserWithdrawsByAppUserRequest{
		AppID:  in.GetTargetAppID(),
		UserID: in.GetTargetUserID(),
	})
	if err != nil {
		logger.Sugar().Errorf("get user withdraw error: %v", err)
		return &npool.GetUserWithdrawsByOtherAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetUserWithdrawsByOtherAppUserResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetUserWithdrawAccount(ctx context.Context, in *npool.GetUserWithdrawAccountRequest) (*npool.GetUserWithdrawAccountResponse, error) {
	resp, err := mw.GetUserWithdrawAccount(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user withdraw account error: %v", err)
		return &npool.GetUserWithdrawAccountResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserWithdrawAccountsByAppUser(ctx context.Context, in *npool.GetUserWithdrawAccountsByAppUserRequest) (*npool.GetUserWithdrawAccountsByAppUserResponse, error) {
	resp, err := mw.GetUserWithdrawAccountsByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user withdraw account error: %v", err)
		return &npool.GetUserWithdrawAccountsByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserWithdrawAccountsByOtherAppUser(ctx context.Context, in *npool.GetUserWithdrawAccountsByOtherAppUserRequest) (*npool.GetUserWithdrawAccountsByOtherAppUserResponse, error) {
	resp, err := s.GetUserWithdrawAccountsByAppUser(ctx, &npool.GetUserWithdrawAccountsByAppUserRequest{
		AppID:  in.GetTargetAppID(),
		UserID: in.GetTargetUserID(),
	})
	if err != nil {
		logger.Sugar().Errorf("get user withdraw account error: %v", err)
		return &npool.GetUserWithdrawAccountsByOtherAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetUserWithdrawAccountsByOtherAppUserResponse{
		Infos: resp.Infos,
	}, nil
}
