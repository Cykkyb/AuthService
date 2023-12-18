package service

import (
	"authService/internal/repository"
	"context"
	"github.com/Cykkyb/proto/gen/go/auth"
	"log/slog"
)

type Auth interface {
	Login(ctx context.Context, email, password string, appId int) (*auth.LoginResponse, error)
	Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error)
	IsAdmin(ctx context.Context, req *auth.IsAdminRequest) (*auth.IsAdminResponse, error)
}

type Service struct {
	Auth
}

func NewService(repo repository.Auth, log *slog.Logger) *Service {
	return &Service{
		Auth: NewAuthService(repo, log),
	}
}
