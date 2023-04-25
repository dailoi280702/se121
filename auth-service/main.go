package main

import (
	"context"

	"github.com/dailoi280702/se121/auth_service/internal/service/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServiceServer struct {
	auth.UnimplementedAuthServiceServer
}

func (s *AuthServiceServer) SignIn(context.Context, *auth.SignInRes) (*auth.SignInReq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}

func (s *AuthServiceServer) SignUp(context.Context, *auth.SignUpReq) (*auth.SignUpRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}

func (s *AuthServiceServer) Refresh(context.Context, *auth.RefreshReq) (*auth.RefreshRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}

func (s *AuthServiceServer) SignOut(context.Context, *auth.SignOutReq) (*auth.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}

func main() {
	print("hello")
}
