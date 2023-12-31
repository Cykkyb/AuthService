package jwt

import (
	"authService/internal/entity"
	"testing"
	"time"
)

func TestGenerateAccessToken(t *testing.T) {
	user := &entity.User{
		Id: 1,
	}
	app := &entity.App{
		Id:   1,
		Name: "TestApp",
	}

	duration := time.Hour

	token, err := GenerateAccessToken(user, app, duration)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if token == "" {
		t.Errorf("expected non-empty token, got empty token")
	}
}
