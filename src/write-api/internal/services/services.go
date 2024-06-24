package services

import ssov1 "github.com/fenek-dev/go-twitter/proto/protogen"

type Services struct {
	sso ssov1.AuthServiceClient
}

func New(sso ssov1.AuthServiceClient) *Services {
	return &Services{
		sso: sso,
	}
}
