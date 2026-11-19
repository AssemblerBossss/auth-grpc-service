package auth

import (
	"auth-grpc-service/internal/domain/models"
	"context"
	"log/slog"
	"time"
)

type Auth struct {
	usrSave     UserSaver
	usrProvider UserProvider
	appProvider AppProvider
	log         *slog.Logger
	tokenTTL    time.Duration
}

type UserSaver interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (uid int64, err error)
}

type UserProvider interface {
	User(cxt context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

type AppProvider interface {
	App(cxt context.Context, appID string) (models.App, error)
}

func New(userSaver UserSaver, userProvider UserProvider, appProvider AppProvider, log *slog.Logger, tokenTTL time.Duration) *Auth {
	return &Auth{usrSave: userSaver, usrProvider: userProvider, appProvider: appProvider, log: log, tokenTTL: tokenTTL}
}

func (a *Auth) Login(ctx context.Context, email, password string, appID int32) (string, error) {
	// TODO: бизнес-логика
	return "1234567890", nil
}

func (a *Auth) Register(ctx context.Context, email, password string) (int64, error) {
	// TODO: бизнес-логика
	return 0, nil
}

func (a *Auth) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	return false, nil
}
