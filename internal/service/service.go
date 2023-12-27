package service

import (
	"authService/internal/repository"
	"context"
	"github.com/Cykkyb/proto/gen/go/auth"
	"log/slog"
	"time"
)

type Auth interface {
	Login(ctx context.Context, email, password string, appId int) (string, error)
	Register(ctx context.Context, email, password string) (int32, error)
	IsAdmin(ctx context.Context, req *auth.IsAdminRequest) (bool, error)
}

type Service struct {
	Auth
}

func NewService(repo repository.AuthRepository, log *slog.Logger, tokenTTl time.Duration) *Service {
	return &Service{
		Auth: NewAuthService(repo, log, tokenTTl),
	}
}
