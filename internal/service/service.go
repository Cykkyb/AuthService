package service

import (
	"authService/internal/repository"
	"context"
	"log/slog"
	"time"
)

type Auth interface {
	Login(ctx context.Context, email, password string, appId int) (string, error)
	Register(ctx context.Context, email, password string) (int32, error)
	IsAdmin(ctx context.Context, userId int32) (bool, error)
}

type Service struct {
	Auth
}

func NewService(repo repository.AuthRepository, log *slog.Logger, tokenTTl time.Duration) *Service {
	return &Service{
		Auth: NewAuthService(repo, log, tokenTTl),
	}
}
