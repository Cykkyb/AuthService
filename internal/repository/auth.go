package repository

import (
	"context"
	"github.com/Cykkyb/proto/gen/go/auth"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type AuthPostgres struct {
	log *slog.Logger
	db  *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB, log *slog.Logger) *AuthPostgres {
	return &AuthPostgres{
		log: log,
		db:  db,
	}
}

func (a *AuthPostgres) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	return nil, nil
}

func (a *AuthPostgres) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {

	return nil, nil
}

func (a *AuthPostgres) IsAdmin(ctx context.Context, req *auth.IsAdminRequest) (*auth.IsAdminResponse, error) {
	return nil, nil
}
