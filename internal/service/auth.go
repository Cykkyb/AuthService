package service

import (
	"authService/internal/lib/jwt"
	"authService/internal/repository"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
)

const (
	salt = "gsgfsd@#bb__21F"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type AuthService struct {
	log        *slog.Logger
	repository repository.AuthRepository
	tokenTTl   time.Duration
}

func NewAuthService(repo repository.AuthRepository, log *slog.Logger, tokenTTl time.Duration) *AuthService {
	return &AuthService{
		log:        log,
		repository: repo,
		tokenTTl:   tokenTTl,
	}
}

func (s *AuthService) Login(ctx context.Context, email, password string, appId int) (string, error) {
	s.log.Info("Login", slog.String("email", email))

	user, err := s.repository.GetUser(ctx, email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			s.log.Warn("user not found", slog.String("email", email))

			return "", ErrInvalidCredentials
		}

		s.log.Error("failed to getUser", slog.String("email", email))
		return "", err
	}
	err = checkPassword(user.PasswordHash, password)
	if err != nil {
		s.log.Error("wrong password", slog.String("email", email))

		return "", ErrInvalidCredentials
	}

	app, err := s.repository.GetApp(ctx, appId)
	if err != nil {
		s.log.Error("failed to get app", slog.Int("appId", appId), slog.String("Error", err.Error()))
		return "", err
	}

	token, err := jwt.GenerateAccessToken(user, app, s.tokenTTl)
	if err != nil {
		s.log.Error("failed to generate token", slog.String("email", err.Error()))
		return "", err
	}

	return token, nil
}

func (s *AuthService) Register(ctx context.Context, email, password string) (int32, error) {
	s.log.Info("Register", slog.String("email", email))

	passwordHash := generatePasswordHash(password)
	userId, err := s.repository.Register(ctx, email, passwordHash)
	if err != nil {
		if errors.Is(err, repository.ErrUserExists) {
			s.log.Warn("user already exists", slog.String("email", email))
			return 0, repository.ErrUserExists
		}

		s.log.Error("failed to register user", slog.String("email", email), slog.String("Error", err.Error()))
		return 0, err
	}
	return userId, nil
}

func (s *AuthService) IsAdmin(ctx context.Context, userId int32) (bool, error) {
	s.log.Info("IsAdmin", slog.Int("id", int(userId)))

	isAdmin, err := s.repository.IsAdmin(ctx, userId)
	if err != nil {
		s.log.Error("failed to get isAdmin", slog.Int("id", int(userId)))
		return false, err
	}

	return isAdmin, nil
}

func generatePasswordHash(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)

	return string(hashedPassword)
}

func checkPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password+salt))
}
