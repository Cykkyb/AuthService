package repository

import (
	"authService/internal/entity"
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user not found")
	ErrAppNotFound  = errors.New("wrong password")
)

type AuthPostgres struct {
	db *sqlx.DB
}

type userToken struct {
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
	UserId       int    `json:"user_id" db:"id"`
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

func (a *AuthPostgres) GetUser(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	query := fmt.Sprintf(`select * from %s where email=$1`, usersTable)
	err := a.db.Get(&user, query, email)

	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (a *AuthPostgres) Register(ctx context.Context, email, passwordHash string) (int32, error) {
	query := fmt.Sprintf(`insert into %s (email, password_hash) values ($1, $2) returning id`, usersTable)
	var userId int32
	err := a.db.Get(&userId, query, email, passwordHash)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (a *AuthPostgres) IsAdmin(ctx context.Context, userId int32) (bool, error) {
	var isAdmin bool

	query := fmt.Sprintf(`select is_admin from %s where id=$1`, usersTable)
	err := a.db.Get(&isAdmin, query, userId)

	if err != nil {
		return false, err
	}

	return isAdmin, nil
}

func (r *AuthPostgres) GetApp(ctx context.Context, appId int) (*entity.App, error) {
	var app entity.App

	query := fmt.Sprintf(`select * from %s where id=$1`, appTable)
	err := r.db.Get(&app, query, appId)
	if err != nil {
		return &app, err
	}

	return &app, nil
}

func (r *AuthPostgres) SetRefreshToken(token string, userId int) error {
	query := fmt.Sprintf(`update %s set refresh_token=$1, updated_at=$2 where id=$3`, usersTable)

	if _, err := r.db.Exec(query, token, time.Now(), userId); err != nil {
		return err
	}

	return nil
}

func (r *AuthPostgres) CheckRefreshToken(refreshToken string) (entity.User, error) {
	var user entity.User

	query := fmt.Sprintf(`select id from %s where refresh_token=$1`, usersTable)
	err := r.db.Get(&user, query, refreshToken)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
