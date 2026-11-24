package app

import (
	grpcapp "auth-grpc-service/internal/app/grpc"
	authsvc "auth-grpc-service/internal/services/auth"
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
	tokenTTL time.Duration) *App {

	authService := authsvc.New(log, tokenTTL)
	grpcApp := grpcapp.New(authService, log, grpcPort)
	return &App{
		GRPCSrv: grpcApp,
	}

}
