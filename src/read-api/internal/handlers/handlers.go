package handlers

import "github.com/fenek-dev/go-twitter/src/read-api/internal/storage"

type Handlers struct {
	db *storage.Storage
}

func New(db *storage.Storage) *Handlers {
	return &Handlers{
		db: db,
	}
}
