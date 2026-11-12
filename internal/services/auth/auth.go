package auth

import (
	"context"
	"log/slog"
	"time"
)

type Auth struct {
	log      *slog.Logger
	tokenTTL time.Duration
}

func New(log *slog.Logger, tokenTTL time.Duration) *Auth {
	return &Auth{log: log, tokenTTL: tokenTTL}
}

func (a *Auth) Login(ctx context.Context, email, password string, appID int) (string, error) {
	// TODO: бизнес-логика
	return "1234567890", nil
}

func (a *Auth) Register(ctx context.Context, email, password string) (int64, error) {
	// TODO: бизнес-логика
	return 0, nil
}
