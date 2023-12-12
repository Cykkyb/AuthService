package app

import (
	grpcapp "authService/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(log *slog.Logger, port int, tokenTTL time.Duration) *App {

	gRPCApp := grpcapp.NewApp(log, port)

	return &App{
		GRPCServer: gRPCApp,
	}

}
