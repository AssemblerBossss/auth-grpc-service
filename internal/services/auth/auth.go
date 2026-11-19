package auth

import (
	"auth-grpc-service/internal/domain/models"
	"auth-grpc-service/internal/lib/logger/sl"
	"context"
	"fmt"
	"log/slog"
	"time"

	"golang.org/x/crypto/bcrypt"
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

func New(
	userSaver UserSaver,
	userProvider UserProvider,
	appProvider AppProvider,
	log *slog.Logger,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		usrSave:     userSaver,
		usrProvider: userProvider,
		appProvider: appProvider,
		log:         log,
		tokenTTL:    tokenTTL,
	}
}

// Login checks if user given credentials n the system
func (a *Auth) Login(ctx context.Context, email, password string, appID int32) (string, error) {
	const op = "Auth.Login"

	log := a.log.With(slog.String("op", op), slog.String("email", email))
	log.Info("attempting to login user")

	user, err := a.usrProvider.User(ctx, email)
	if err != nil {

	}

	return "1234567890", nil
}

func (a *Auth) Register(ctx context.Context, email, password string) (int64, error) {
	const op = "auth.RegisterNewUser"
	log := a.log.With(slog.String("op", op), slog.String("email", email))

	log.Info("Registering user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Failed to hash password", sl.Err(err))
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := a.usrSave.SaveUser(ctx, email, passHash)
	if err != nil {
		log.Error("Failed to save user", sl.Err(err))

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (a *Auth) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	return false, nil
}
