package handlers

import (
	"log/slog"

	"github.com/fenek-dev/go-twitter/src/write-api/internal/services"
	"github.com/fenek-dev/go-twitter/src/write-api/internal/storage"
)

type Handlers struct {
	log     *slog.Logger
	db      *storage.Storage
	service *services.Services
}

func New(db *storage.Storage, service *services.Services, log *slog.Logger) *Handlers {
	return &Handlers{
		db:      db,
		log:     log,
		service: service,
	}
}
