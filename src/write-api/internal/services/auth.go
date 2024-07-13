package services

import (
	"context"

	ssov1 "github.com/fenek-dev/go-twitter/proto/protogen"
)

func (s *Services) Register(ctx context.Context, username, password string) (string, error) {
	ctx, span := s.tracer.Start(ctx, "Register service")
	defer span.End()
	res, err := s.sso.Register(ctx, &ssov1.RegisterRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		return "", err
	}

	return res.Token, nil
}

func (s *Services) Login(ctx context.Context, username, password string) (string, error) {
	ctx, span := s.tracer.Start(ctx, "Login service")
	defer span.End()
	res, err := s.sso.Login(ctx, &ssov1.LoginRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		return "", err
	}

	return res.Token, nil
}
