package handlers

import (
	"net/http"

	proto "github.com/fenek-dev/go-twitter/proto/protogen"
	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) FindTweetById(c *gin.Context) {
	id := c.Param("id")
	ctx, span := h.tr.Start(c.Request.Context(), "read.handler.FindTweetById")
	defer span.End()

	if id == "" {
		common.SendResponse(c.Writer, http.StatusBadRequest, "incorrect_id", nil)
		return
	}

	tweet, err := h.db.FindTweetById(ctx, &proto.FindTweetByIdRequest{Id: id})
	if err != nil {
		common.SendResponse(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendResponse(c.Writer, http.StatusOK, "ok", tweet)
}
