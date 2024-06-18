package tweets

import (
	"net/http"

	"github.com/fenek-dev/go-twitter/src/common"
)

type Controller struct {
	repository *Repository
}

func NewController(repository *Repository) *Controller {
	return &Controller{
		repository: repository,
	}
}

func (c *Controller) FindById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		common.SendResponse(w, http.StatusBadRequest, "Incorrect id", nil)
		return
	}

	tweet, err := c.repository.FindById(r.Context(), id)
	if err != nil {
		common.SendResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendResponse(w, http.StatusOK, "ok", tweet)
}
