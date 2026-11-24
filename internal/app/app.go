package app

import (
	grpcapp "auth-grpc-service/internal/app/grpc"
	authsvc "auth-grpc-service/internal/services/auth"
	"auth-grpc-service/internal/storage/sqlite"
	"fmt"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration) (*App, error) {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		return nil, fmt.Errorf("initialize storage: %w", err)
	}
	authService := authsvc.New(storage, storage, storage, log, tokenTTL)
	grpcApp := grpcapp.New(authService, log, grpcPort)
	return &App{
		GRPCSrv: grpcApp,
	}, nil
}
