// +build !codeanalysis

package api

import (
	"context"

	crud "github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/goodbenefit"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGoodBenefit(ctx context.Context, in *npool.CreateGoodBenefitRequest) (*npool.CreateGoodBenefitResponse, error) {
	resp, err := crud.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create good benefit error: %v", err)
		return &npool.CreateGoodBenefitResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateGoodBenefit(ctx context.Context, in *npool.UpdateGoodBenefitRequest) (*npool.UpdateGoodBenefitResponse, error) {
	resp, err := crud.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update good benefit error: %v", err)
		return &npool.UpdateGoodBenefitResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetGoodBenefit(ctx context.Context, in *npool.GetGoodBenefitRequest) (*npool.GetGoodBenefitResponse, error) {
	resp, err := crud.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good benefit error: %v", err)
		return &npool.GetGoodBenefitResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetGoodBenefitByGood(ctx context.Context, in *npool.GetGoodBenefitByGoodRequest) (*npool.GetGoodBenefitByGoodResponse, error) {
	resp, err := crud.GetByGood(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get good benefit error: %v", err)
		return &npool.GetGoodBenefitByGoodResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
