package auth

import (
	"log/slog"
)

type Controller struct {
	log     *slog.Logger
	service *Service
}

func NewController(log *slog.Logger, service *Service) *Controller {
	return &Controller{
		log:     log,
		service: service,
	}
}
