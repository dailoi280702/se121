package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"github.com/dailoi280702/se121/car_service/pkg/car"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var serverAdress = flag.String("server address", "[::]:5050", "server address of car serivce")

type carSerivceServer struct {
	car.UnimplementedCarServiceServer
}

func NewServer() *carSerivceServer {
	return &carSerivceServer{}
}

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

func (s *carSerivceServer) SearchForCar(*car.SearchForCarReq, car.CarService_SearchForCarServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchForCar not implemented")
}

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

func (s *carSerivceServer) GetSeries(context.Context, *car.GetSeriesReq) (*car.Series, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSeries not implemented")
}

func (s *carSerivceServer) CreateSeries(context.Context, *car.CreateSeriesReq) (*car.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSeries not implemented")
}

func (s *carSerivceServer) UpdateSeries(context.Context, *car.UpdateSeriesReq) (*car.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSeries not implemented")
}

func (s *carSerivceServer) SearchForSeries(*car.SearchForSeriesReq, car.CarService_SearchForSeriesServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchForSeries not implemented")
}

func main() {
	lis, err := net.Listen("tcp", *serverAdress)
	if err != nil {
		fmt.Printf("failed to listen on %s: %v", *serverAdress, err)
	}
	server := grpc.NewServer()
	car.RegisterCarServiceServer(server, NewServer())
	if err = server.Serve(lis); err != nil {
		fmt.Printf("failed to serve car server: %v", err)
	}
}
