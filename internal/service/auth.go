package service

import (
	"authService/internal/repository"
	"context"
	"github.com/Cykkyb/proto/gen/go/auth"
	"log/slog"
)

type AuthService struct {
	log        *slog.Logger
	repository repository.Auth
}

func NewAuthService(repo repository.Auth, log *slog.Logger) *AuthService {
	return &AuthService{
		log:        log,
		repository: repo,
	}
}

func (s *AuthService) Login(ctx context.Context, email, password string, appId int) (*auth.LoginResponse, error) {
	return nil, nil
}

func (s *AuthService) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	return nil, nil
}

func (s *AuthService) IsAdmin(ctx context.Context, req *auth.IsAdminRequest) (*auth.IsAdminResponse, error) {
	return nil, nil
}
