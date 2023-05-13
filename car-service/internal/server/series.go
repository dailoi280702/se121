package server

import (
	"context"

	"github.com/dailoi280702/se121/car-service/pkg/car"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *carSerivceServer) GetSeries(context.Context, *car.GetSeriesReq) (*car.Series, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSeries not implemented")
}

func (s *carSerivceServer) CreateSeries(context.Context, *car.CreateSeriesReq) (*car.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSeries not implemented")
}

func (s *carSerivceServer) UpdateSeries(context.Context, *car.UpdateSeriesReq) (*car.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSeries not implemented")
}

func (s *carSerivceServer) SearchForSeries(context.Context, *car.SearchForSeriesReq) (*car.SearchForSeriesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchForCar not implemented")
}
