package handlers

import (
	"github.com/fenek-dev/go-twitter/src/cache/internal/storage/pg"
)

type Handlers struct {
	db *pg.Postgres
}

func New(db *pg.Postgres) *Handlers {
	return &Handlers{
		db: db,
	}
}
