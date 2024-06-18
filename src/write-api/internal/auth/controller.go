package auth

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/fenek-dev/go-twitter/src/write-api/internal/auth/dto"
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

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	var data dto.RegisterDto

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usrname, err := c.service.Register(r.Context(), data.Username, data.Password)
	if err != nil || usrname == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	var data dto.LoginDto

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usrname, err := c.service.Login(r.Context(), data.Username, data.Password)
	if err != nil || usrname == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
