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

	err := validateLoginRequest(req)
	if err != nil {
		return nil, err
	}

	token, err := s.auth.Login(ctx, req.GetEmail(), req.GetPassword(), int(req.GetAppId()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &authproto.LoginResponse{
		Token: token,
	}, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *authproto.RegisterRequest,
) (*authproto.RegisterResponse, error) {

	err := validateRegisterRequest(req)
	if err != nil {
		return nil, err
	}

	userId, err := s.auth.Register(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &authproto.RegisterResponse{
		UserId: userId,
	}, nil

}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *authproto.IsAdminRequest,
) (*authproto.IsAdminResponse, error) {

	err := validateIsAdminRequest(req)
	if err != nil {
		return nil, err
	}

	isAdmin, err := s.auth.IsAdmin(ctx, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &authproto.IsAdminResponse{
		IsAdmin: isAdmin,
	}, nil
}

func validateLoginRequest(req *authproto.LoginRequest) error {

	if req.Email == "" {
		return status.Error(codes.InvalidArgument, "missing email")
	}

	if req.Password == "" {
		return status.Error(codes.InvalidArgument, "missing password")
	}

	if req.AppId == 0 {
		return status.Error(codes.InvalidArgument, "missing appId")
	}

	return nil
}

func validateRegisterRequest(req *authproto.RegisterRequest) error {

	if req.Email == "" {
		return status.Error(codes.InvalidArgument, "missing email")
	}

	if req.Password == "" {
		return status.Error(codes.InvalidArgument, "missing password")
	}

	return nil
}

func validateIsAdminRequest(req *authproto.IsAdminRequest) error {

	if req.UserId == 0 {
		return status.Error(codes.InvalidArgument, "missing userId")
	}

	return nil
}
