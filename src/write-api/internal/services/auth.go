package services

import (
	"context"

	ssov1 "github.com/fenek-dev/go-twitter/src/sso/protogen"
)

func (s *Services) Register(ctx context.Context, username, password string) (string, error) {
	res, err := s.sso.Register(ctx, &ssov1.RegisterRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		return "", err
	}

	return res.Username, nil
}

func (s *Services) Login(ctx context.Context, username, password string) (string, error) {
	res, err := s.sso.Login(ctx, &ssov1.LoginRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		return "", err
	}

	return res.Token, nil
}
