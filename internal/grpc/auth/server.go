package auth

import (
	"auth-grpc-service/internal/services/auth"
	"context"
	"errors"

	ssov1 "github.com/AssemblerBossss/auth-grpc-api/pkg/v1/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	auth Auth
}
type Auth interface {
	Login(ctx context.Context, email string, password string, appID int32) (token string, err error)
	Register(ctx context.Context, email string, password string) (userID int64, err error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

func Register(gRPC *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{auth: auth})
}

const (
	emptyValue = 0
)

func (s *serverAPI) Login(
	ctx context.Context,
	in *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	if err := validateLogin(in); err != nil {
		return nil, err
	}
	token, err := s.auth.Login(ctx, in.GetEmail(), in.GetPassword(), in.GetAppId())
	if err != nil {
		if errors.Is(err, auth.ErrInvalidCredentials) {
			return nil, status.Error(codes.InvalidArgument, "invalid credentials")
		}
		if errors.Is(err, auth.ErrInvalidAppID) {
			return nil, status.Error(codes.NotFound, "application not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &ssov1.LoginResponse{Token: token}, err
}

func (s *serverAPI) Register(
	ctx context.Context,
	in *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	if err := validateRegister(in); err != nil {
		return nil, err
	}
	userID, err := s.auth.Register(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
		if errors.Is(err, auth.ErrUserExists) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssov1.RegisterResponse{UserId: userID}, err
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	in *ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponse, error) {
	if err := validateIsAdmin(in); err != nil {
		return nil, err
	}
	isAdmin, err := s.auth.IsAdmin(ctx, in.GetUserId())
	if err != nil {
		if errors.Is(err, auth.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssov1.IsAdminResponse{IsAdmin: isAdmin}, nil
}

func validateLogin(in *ssov1.LoginRequest) error {
	if in.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "missing email")
	}
	if in.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "missing password")
	}
	if in.GetAppId() == emptyValue {
		return status.Error(codes.InvalidArgument, "app_id is required")
	}
	return nil
}

func validateRegister(in *ssov1.RegisterRequest) error {
	if in.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "missing email")
	}
	if in.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "missing password")
	}
	return nil
}
func validateIsAdmin(in *ssov1.IsAdminRequest) error {
	if in.GetUserId() == emptyValue {
		return status.Error(codes.InvalidArgument, "user_id is required")
	}
	return nil
}
