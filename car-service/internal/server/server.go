package server

import "github.com/dailoi280702/se121/car_service/pkg/car"

type carSerivceServer struct {
	car.UnimplementedCarServiceServer
}

func NewServer() *carSerivceServer {
	return &carSerivceServer{}
}
