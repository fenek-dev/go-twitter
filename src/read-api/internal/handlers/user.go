package handlers

import (
	"net/http"

	proto "github.com/fenek-dev/go-twitter/proto/protogen"
	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/models"
)

func (h *Handlers) Me(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(common.REQUEST_CTX_USER).(models.User)
	if !ok {
		common.SendResponse(w, http.StatusInternalServerError, "Something gone wrong", nil)
	}

	common.SendResponse(w, http.StatusOK, "ok", user)
}

func (h *Handlers) FindUserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		common.SendResponse(w, http.StatusBadRequest, "Incorrect id", nil)
		return
	}

	tweet, err := h.db.FindUserById(r.Context(), &proto.FindUserByIdRequest{Id: id})
	if err != nil {
		common.SendResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendResponse(w, http.StatusOK, "ok", tweet)
}
