package handlers

import (
	proto "github.com/fenek-dev/go-twitter/proto/protogen"
)

type Handlers struct {
	db proto.CacheServiceClient
}

func New(db proto.CacheServiceClient) *Handlers {
	return &Handlers{
		db: db,
	}
}
