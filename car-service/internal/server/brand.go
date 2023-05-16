package server

import (
	"context"

	"github.com/dailoi280702/se121/car-service/pkg/car"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *carSerivceServer) GetBrand(ctx context.Context, req *car.GetBrandReq) (*car.Brand, error) {
	id := int(req.GetId())
	brand, err := dbGetBrandById(s.db, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while get car data from db: %v", err)
	}
	if brand == nil {
		return nil, status.Errorf(codes.NotFound, "car %d not exists", id)
	}
	return brand, nil
}

func (s *carSerivceServer) CreateBrand(context.Context, *car.CreateBrandReq) (*car.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBrand not implemented")
}

func (s *carSerivceServer) UpdateBrand(context.Context, *car.UpdateBrandReq) (*car.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBrand not implemented")
}

func (s *carSerivceServer) SearchForBrand(context.Context, *car.SearchReq) (*car.SearchForBrandRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchForCar not implemented")
}
