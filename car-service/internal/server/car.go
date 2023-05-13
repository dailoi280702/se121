package server

import (
	"context"

	"github.com/dailoi280702/se121/car-service/pkg/car"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *carSerivceServer) GetCar(context.Context, *car.GetCarReq) (*car.Car, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCar not implemented")
}

func (s *carSerivceServer) CreateCar(context.Context, *car.CreateCarReq) (*car.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCar not implemented")
}

func (s *carSerivceServer) UpdateCar(context.Context, *car.UpdateCarReq) (*car.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCar not implemented")
}

func (s *carSerivceServer) DeleteCar(context.Context, *car.DeleteCarReq) (*car.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCar not implemented")
}

func (s *carSerivceServer) SearchForCar(context.Context, *car.SearchForCarReq) (*car.SearchForCarRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchForCar not implemented")
}

func (s *carSerivceServer) GetCarMetadata(context.Context, *car.Empty) (*car.GetCarMetadataRes, error) {
	brands, err := getAllBrandFromDb(s.db)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "car service err: %v", err)
	}
	series, err := getAllSeriesFromDb(s.db)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "car service err: %v", err)
	}
	fuelTypes, err := getAllFuelTypesFromDb(s.db)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "car service err: %v", err)
	}
	transmissions, err := getAllTransmissionFromDb(s.db)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "car service err: %v", err)
	}
	res := car.GetCarMetadataRes{
		Brands:       brands,
		Series:       series,
		FuelType:     fuelTypes,
		Transmission: transmissions,
	}
	return &res, nil
}
