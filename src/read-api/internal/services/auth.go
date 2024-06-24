package auth

import (
	"context"

	ssov1 "github.com/fenek-dev/go-twitter/proto/protogen"
)

type Service struct {
	sso ssov1.AuthServiceClient
}

func NewService(sso ssov1.AuthServiceClient) *Service {
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
