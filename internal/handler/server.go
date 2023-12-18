package handler

import (
	"authService/internal/service"
	"context"
	authproto "github.com/Cykkyb/proto/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	authproto.UnimplementedAuthServer
	auth service.Auth
}

func RegisterServerAPI(gRPC *grpc.Server, auth service.Auth) {
	authproto.RegisterAuthServer(gRPC, &serverAPI{auth: auth})
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *authproto.LoginRequest,
) (*authproto.LoginResponse, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "missing email")
	}

	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "missing password")
	}

	if req.AppId == 0 {
		return nil, status.Error(codes.InvalidArgument, "missing appId")
	}

	s.auth.Login(ctx, req.GetEmail(), req.GetPassword(), int(req.GetAppId()))

	return &authproto.LoginResponse{
		Token: "test",
	}, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *authproto.RegisterRequest,
) (*authproto.RegisterResponse, error) {
	panic("test")
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *authproto.IsAdminRequest,
) (*authproto.IsAdminResponse, error) {
	panic("test")
}
