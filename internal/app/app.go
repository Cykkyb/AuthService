package app

import (
	"authService/internal/handler"
	"authService/internal/repository"
	"authService/internal/service"
	"fmt"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"time"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewApp(log *slog.Logger, port int, tokenTTL time.Duration, db *sqlx.DB) *App {
	gRPCServer := grpc.NewServer()

	repo := repository.NewRepository(db, log)
	services := service.NewService(repo, log)
	handler.RegisterServerAPI(gRPCServer, services)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", "grpc.Run", err)
	}

	a.log.Info("Starting grpc")

	if err = a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", "grpc.Run", err)
	}

	return nil
}

func (a *App) Stop() {
	a.log.Info("Stopping grpc")

	a.gRPCServer.GracefulStop()
}
