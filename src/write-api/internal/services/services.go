package services

import (
	proto "github.com/fenek-dev/go-twitter/proto/protogen"
	ssov1 "github.com/fenek-dev/go-twitter/proto/protogen"
)

type Services struct {
	sso   ssov1.AuthServiceClient
	cache proto.CacheServiceClient
}

func New(sso ssov1.AuthServiceClient, cache proto.CacheServiceClient) *Services {
	return &Services{
		sso:   sso,
		cache: cache,
	}
}
