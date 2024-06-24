package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/write-api/internal/dto"
)

func (h *Handlers) CreateTweet(w http.ResponseWriter, r *http.Request) {
	var data *dto.CreateDto

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		common.SendResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	tweet, err := h.service.CreateTweet(r.Context(), data.Username, data.Content)
	if err != nil {
		common.SendResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendResponse(w, http.StatusCreated, "ok", tweet)
}

func (h *Handlers) UpdateTweet(w http.ResponseWriter, r *http.Request) {
	var data *dto.UpdateDto

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		common.SendResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	tweet, err := h.service.UpdateTweet(r.Context(), data.Id, data.Content)
	if err != nil {
		common.SendResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendResponse(w, http.StatusOK, "ok", tweet)
}

func (h *Handlers) DeleteTweet(w http.ResponseWriter, r *http.Request) {
	var data *dto.DeleteDto

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		common.SendResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	id, err := h.service.DeleteTweet(r.Context(), data.Id)
	if err != nil {
		common.SendResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendResponse(w, http.StatusOK, "ok", id)
}
