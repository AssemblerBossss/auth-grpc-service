package grpcapp

import (
	authgrpc "auth-grpc-service/internal/grpc/auth"
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

// New creates new gRPC server app
func New(log *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()
	authgrpc.Register(gRPCServer)
	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (s *App) MustRun() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}

func (s *App) Run() error {
	const op = "grpcapp.Run"
	log := s.log.With(slog.String("op", "op"))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("started gRPC server", slog.String("addr", l.Addr().String()))
	return s.gRPCServer.Serve(l)
}

func (s *App) Stop() {
	const op = "grpcapp.Stop"

	s.log.With(slog.String("op", "op")).Info("stopped gRPC server", slog.Int("port", s.port))
	s.gRPCServer.GracefulStop()

}
