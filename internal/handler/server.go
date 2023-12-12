package handler

import (
	"context"
	"github.com/Cykkyb/proto/gen/go/auth"
	"google.golang.org/grpc"
)

type serverAPI struct {
	auth.UnimplementedAuthServer
}

func RegisterServerAPI(gRPC *grpc.Server) {
	auth.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *auth.LoginRequest,
) (*auth.LoginResponse, error) {
	return &auth.LoginResponse{
		Token: "test",
	}, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *auth.RegisterRequest,
) (*auth.RegisterResponse, error) {
	panic("test")
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *auth.IsAdminRequest,
) (*auth.IsAdminResponse, error) {
	panic("test")
}
