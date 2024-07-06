package handlers

import (
	proto "github.com/fenek-dev/go-twitter/proto/protogen"
	"go.opentelemetry.io/otel/trace"
)

type Handlers struct {
	db proto.CacheServiceClient
	tr trace.Tracer
}

func New(db proto.CacheServiceClient, tr trace.Tracer) *Handlers {
	return &Handlers{
		db: db,
		tr: tr,
	}
}
