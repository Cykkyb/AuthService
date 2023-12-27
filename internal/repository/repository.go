package repository

import (
	"authService/internal/entity"
	"context"
	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	Register(ctx context.Context, email, passwordHash string) (int32, error)
	IsAdmin(ctx context.Context, userId int32) (bool, error)
	GetUser(ctx context.Context, email string) (*entity.User, error)
	GetApp(ctx context.Context, appId int) (*entity.App, error)
}

type Repository struct {
	AuthRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthRepository: NewAuthPostgres(db),
	}
}
