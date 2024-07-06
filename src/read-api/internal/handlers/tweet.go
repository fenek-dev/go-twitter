package handlers

import (
	"net/http"

	proto "github.com/fenek-dev/go-twitter/proto/protogen"
	"github.com/fenek-dev/go-twitter/src/common"
)

func (h *Handlers) FindTweetById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	ctx, span := h.tr.Start(r.Context(), "read.FindTweetById")
	defer span.End()

	if id == "" {
		common.SendResponse(w, http.StatusBadRequest, "incorrect_id", nil)
		return
	}

	tweet, err := h.db.FindTweetById(ctx, &proto.FindTweetByIdRequest{Id: id})
	if err != nil {
		common.SendResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendResponse(w, http.StatusOK, "ok", tweet)
}
