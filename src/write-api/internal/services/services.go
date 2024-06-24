package services

import ssov1 "github.com/fenek-dev/go-twitter/src/sso/protogen"

type Services struct {
	sso ssov1.AuthClient
}

func New(sso ssov1.AuthClient) *Services {
	return &Services{
		sso: sso,
	}
}
