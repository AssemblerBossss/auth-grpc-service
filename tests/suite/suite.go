package suite

import (
	"auth-grpc-service/internal/config"
	"context"
	"net"
	"strconv"
	"testing"

	ssov1 "github.com/AssemblerBossss/auth-grpc-api/pkg/v1/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const configPath = "../config/local_tests.yaml"

type Suite struct {
	*testing.T
	Cfg        *config.Config   // Конфигурация приложения
	AuthClient ssov1.AuthClient // Клиент для взаимодействия с gRPC-сервером Auth
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoadByPath(configPath)

	// Основной родительский контекст
	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	grpcAddress := net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.Port))

	cc, err := grpc.NewClient(
		grpcAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("grpc server connection failed %v", err)
	}
	authClient := ssov1.NewAuthClient(cc)

	return ctx, &Suite{
		T:          t,
		Cfg:        cfg,
		AuthClient: authClient,
	}

}
