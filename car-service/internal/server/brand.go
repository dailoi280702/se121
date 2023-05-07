package server

import (
	"context"

	"github.com/dailoi280702/se121/car_service/pkg/car"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *carSerivceServer) GetBrand(context.Context, *car.GetBrandReq) (*car.Brand, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrand not implemented")
}

func (s *carSerivceServer) CreateBrand(context.Context, *car.CreateBrandReq) (*car.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBrand not implemented")
}

func (s *carSerivceServer) UpdateBrand(context.Context, *car.UpdateBrandReq) (*car.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBrand not implemented")
}

func (s *carSerivceServer) SearchForBrand(*car.SearchForBrandReq, car.CarService_SearchForBrandServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchForBrand not implemented")
}
