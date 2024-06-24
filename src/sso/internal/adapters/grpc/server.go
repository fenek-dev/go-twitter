package grpc

import (
	"context"

	ssov1 "github.com/fenek-dev/go-twitter/proto/protogen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServiceServer
	auth Auth
}

type Auth interface {
	Login(
		ctx context.Context,
		username string,
		password string,
	) (token string, err error)
	RegisterNewUser(
		ctx context.Context,
		username string,
		password string,
	) (usrname string, err error)
}

func Register(gRPCServer *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServiceServer(gRPCServer, &serverAPI{auth: auth})
}

func (s *serverAPI) Login(
	ctx context.Context,
	in *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	if in.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "username is required")
	}

	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	token, err := s.auth.Login(ctx, in.GetUsername(), in.GetPassword())
	if err != nil {
		// if errors.Is(err, auth.ErrInvalidCredentials) {
		// 	return nil, status.Error(codes.InvalidArgument, "invalid email or password")
		// }

		return nil, status.Error(codes.Internal, "failed to login")
	}

	return &ssov1.LoginResponse{Token: token}, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	in *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	if in.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "username is required")
	}

	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	urname, err := s.auth.RegisterNewUser(ctx, in.GetUsername(), in.GetPassword())
	if err != nil {
		// if errors.Is(err, storage.ErrUserExists) {
		// 	return nil, status.Error(codes.AlreadyExists, "user already exists")
		// }

		return nil, status.Error(codes.Internal, "failed to register user")
	}

	return &ssov1.RegisterResponse{Username: urname}, nil
}
