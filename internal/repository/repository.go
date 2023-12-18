package repository

import (
	"context"
	"github.com/Cykkyb/proto/gen/go/auth"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type Auth interface {
	Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error)
	Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error)
	IsAdmin(ctx context.Context, req *auth.IsAdminRequest) (*auth.IsAdminResponse, error)
}

type Repository struct {
	Auth
}

func NewRepository(db *sqlx.DB, log *slog.Logger) *Repository {
	return &Repository{
		Auth: NewAuthPostgres(db, log),
	}
}
