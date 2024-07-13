package services

import (
	ssov1 "github.com/fenek-dev/go-twitter/proto/protogen"

	"go.opentelemetry.io/otel/trace"
)

type Services struct {
	sso   ssov1.AuthServiceClient
	cache ssov1.CacheServiceClient

	tracer trace.Tracer
}

func New(sso ssov1.AuthServiceClient, cache ssov1.CacheServiceClient,
	tracer trace.Tracer) *Services {
	return &Services{
		sso:    sso,
		cache:  cache,
		tracer: tracer,
	}
}
