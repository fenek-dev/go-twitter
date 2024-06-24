package auth

import (
	"context"

	ssov1 "github.com/fenek-dev/go-twitter/src/sso/protogen"
)

type Service struct {
	sso ssov1.AuthClient
}

func NewService(sso ssov1.AuthClient) *Service {
	return &Service{
		sso: sso,
	}
}

func (s *Service) Verify(ctx context.Context, token string) (string, error) {
	// res, err := s.sso.Verify(ctx, &ssov1.VerifyRequest{
	// 	Token: token,
	// })

	// if err != nil {
	// 	return "", err
	// }

	// return res.Username, nil
	return "", nil
}
