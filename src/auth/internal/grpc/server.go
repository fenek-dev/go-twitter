package grpc

import (
	"context"

	proto "github.com/fenek-dev/go-twitter/proto/protogen"
	"github.com/fenek-dev/go-twitter/src/common/mappers"
	"github.com/fenek-dev/go-twitter/src/common/models"
	"go.opentelemetry.io/otel/trace"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	proto.UnimplementedAuthServiceServer
	auth Auth
	tr   trace.Tracer
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
	) (token string, err error)
	Verify(
		ctx context.Context,
		token string,
	) (user *models.User, err error)
}

func Register(gRPCServer *grpc.Server, auth Auth, tr trace.Tracer) {
	proto.RegisterAuthServiceServer(gRPCServer, &serverAPI{auth: auth, tr: tr})
}

func (s *serverAPI) Login(
	ctx context.Context,
	in *proto.LoginRequest,
) (*proto.LoginResponse, error) {
	ctx, span := s.tr.Start(ctx, "sso.grpc.Login")
	defer span.End()
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

	return &proto.LoginResponse{Token: token}, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	in *proto.RegisterRequest,
) (*proto.RegisterResponse, error) {
	ctx, span := s.tr.Start(ctx, "sso.grpc.Register")
	defer span.End()
	if in.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "username is required")
	}

	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	token, err := s.auth.RegisterNewUser(ctx, in.GetUsername(), in.GetPassword())
	if err != nil {
		// if errors.Is(err, storage.ErrUserExists) {
		// 	return nil, status.Error(codes.AlreadyExists, "user already exists")
		// }

		return nil, status.Error(codes.Internal, "failed to register user")
	}

	return &proto.RegisterResponse{Token: token}, nil
}

func (s *serverAPI) Verify(
	ctx context.Context,
	in *proto.VerifyRequest,
) (*proto.VerifyResponse, error) {
	ctx, span := s.tr.Start(ctx, "sso.grpc.Verify")
	defer span.End()
	user, err := s.auth.Verify(ctx, in.Token)

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to verify user")
	}

	protoUser := mappers.UserModelToProtoUser(user)

	return &proto.VerifyResponse{User: protoUser}, nil
}
