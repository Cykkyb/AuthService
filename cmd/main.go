package main

import (
	"authService/internal/app"
	"authService/internal/config"
	"authService/internal/lib/logger"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoadConfig()

	log := initLogger()
	log.Info("Config loaded",
		slog.String("env", cfg.Env),
		slog.Int("port", cfg.GRPC.Port),
	)

	application := app.NewApp(log, cfg.GRPC.Port, cfg.TokenTTL)

	go application.GRPCServer.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
}

func initLogger() *slog.Logger {
	opts := logger.PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}
	handler := logger.NewPrettyHandler(os.Stdout, opts)
	log := slog.New(handler)

	return log
}
