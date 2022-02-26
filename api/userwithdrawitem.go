// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/userwithdrawitem"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUserWithdrawItem(ctx context.Context, in *npool.CreateUserWithdrawItemRequest) (*npool.CreateUserWithdrawItemResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create user withdraw item error: %v", err)
		return &npool.CreateUserWithdrawItemResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateUserWithdrawItem(ctx context.Context, in *npool.UpdateUserWithdrawItemRequest) (*npool.UpdateUserWithdrawItemResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update user withdraw item error: %v", err)
		return &npool.UpdateUserWithdrawItemResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserWithdrawItem(ctx context.Context, in *npool.GetUserWithdrawItemRequest) (*npool.GetUserWithdrawItemResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user withdraw item error: %v", err)
		return &npool.GetUserWithdrawItemResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserWithdrawItemsByAccount(ctx context.Context, in *npool.GetUserWithdrawItemsByAccountRequest) (*npool.GetUserWithdrawItemsByAccountResponse, error) {
	resp, err := crud.GetByAccount(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user withdraw item error: %v", err)
		return &npool.GetUserWithdrawItemsByAccountResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserWithdrawItemsByAppUser(ctx context.Context, in *npool.GetUserWithdrawItemsByAppUserRequest) (*npool.GetUserWithdrawItemsByAppUserResponse, error) {
	resp, err := crud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user withdraw item error: %v", err)
		return &npool.GetUserWithdrawItemsByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserWithdrawItemsByAppUserCoinWithdrawType(ctx context.Context, in *npool.GetUserWithdrawItemsByAppUserCoinWithdrawTypeRequest) (*npool.GetUserWithdrawItemsByAppUserCoinWithdrawTypeResponse, error) {
	resp, err := crud.GetByAppUserCoinWithdrawType(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user withdraw item error: %v", err)
		return &npool.GetUserWithdrawItemsByAppUserCoinWithdrawTypeResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserWithdrawItemsByOtherAppUser(ctx context.Context, in *npool.GetUserWithdrawItemsByOtherAppUserRequest) (*npool.GetUserWithdrawItemsByOtherAppUserResponse, error) {
	resp, err := crud.GetByAppUser(ctx, &npool.GetUserWithdrawItemsByAppUserRequest{
		AppID:  in.GetTargetAppID(),
		UserID: in.GetTargetUserID(),
	})
	if err != nil {
		logger.Sugar().Errorf("get user withdraw item error: %v", err)
		return &npool.GetUserWithdrawItemsByOtherAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetUserWithdrawItemsByOtherAppUserResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetUserWithdrawItems(ctx context.Context, in *npool.GetUserWithdrawItemsRequest) (*npool.GetUserWithdrawItemsResponse, error) {
	resp, err := crud.GetAll(ctx, &npool.GetUserWithdrawItemsRequest{})
	if err != nil {
		logger.Sugar().Errorf("get user withdraw item error: %v", err)
		return &npool.GetUserWithdrawItemsResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
