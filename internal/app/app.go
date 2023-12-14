package app

import (
	grpcapp "grpc/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, DBURI string, TokenTTL time.Duration) *App {
	grpcApp := grpcapp.New(log, grpcPort)

	return &App{GRPCServer: grpcApp}
}
