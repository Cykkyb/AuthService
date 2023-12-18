package grpcapp

import (
	"authService/internal/handler"
	"authService/internal/repository"
	"authService/internal/service"
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewApp(log *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()

	db, err := repository.ConnectDb(repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "postgres",
		DBname:   "postgres",
		SSL:      "disable",
	})
	if err != nil {
		log.Error("Failed to connect db", err)
	}

	repository := repository.NewRepository(db, log)
	service := service.NewService(repository, log)
	handler.RegisterServerAPI(gRPCServer, service)

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
