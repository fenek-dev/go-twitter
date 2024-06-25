package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/write-api/internal/dto"
)

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	var data dto.RegisterDto

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		common.SendResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	usrname, err := h.service.Register(r.Context(), data.Username, data.Password)
	if err != nil || usrname == "" {
		common.SendResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendResponse(w, http.StatusCreated, "ok", usrname)
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	var data dto.LoginDto

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		common.SendResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	token, err := h.service.Login(r.Context(), data.Username, data.Password)
	if err != nil || token == "" {
		common.SendResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendResponse(w, http.StatusCreated, "ok", token)
}
