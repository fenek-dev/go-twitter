package tweets

import (
	"encoding/json"
	"net/http"

	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/write-api/internal/tweets/dto"
)

type Controller struct {
	repository *Repository
}

func NewController(repository *Repository) *Controller {
	return &Controller{
		repository: repository,
	}
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var data *dto.CreateDto

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tweet, err := c.repository.Create(r.Context(), data.Username, data.Content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	common.SendResponse(w, http.StatusCreated, "ok", tweet)

}