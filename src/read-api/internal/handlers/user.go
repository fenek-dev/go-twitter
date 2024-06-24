package handlers

import (
	"net/http"

	proto "github.com/fenek-dev/go-twitter/proto/protogen"
	"github.com/fenek-dev/go-twitter/src/common"
)

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
