package handlers

import (
	"log/slog"

	"github.com/fenek-dev/go-twitter/src/write-api/internal/services"
)

type Handlers struct {
	log     *slog.Logger
	service *services.Services
}

func New(service *services.Services, log *slog.Logger) *Handlers {
	return &Handlers{
		log:     log,
		service: service,
	}
}
