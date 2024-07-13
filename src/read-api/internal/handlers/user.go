package handlers

import (
	"net/http"

	proto "github.com/fenek-dev/go-twitter/proto/protogen"
	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/models"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) Me(c *gin.Context) {
	ctx, span := h.tr.Start(c.Request.Context(), "read.handler.Me")
	defer span.End()
	user, ok := ctx.Value(common.REQUEST_CTX_USER).(models.User)
	if !ok {
		common.SendResponse(c.Writer, http.StatusInternalServerError, "Something gone wrong", nil)
		return
	}

	common.SendResponse(c.Writer, http.StatusOK, "ok", user)
}

func (h *Handlers) FindUserById(c *gin.Context) {
	id := c.Param("id")
	ctx, span := h.tr.Start(c.Request.Context(), "read.handler.FindUserById")
	defer span.End()

	if id == "" {
		common.SendResponse(c.Writer, http.StatusBadRequest, "incorrect_id", nil)
		return
	}

	tweet, err := h.db.FindUserById(ctx, &proto.FindUserByIdRequest{Id: id})
	if err != nil {
		common.SendResponse(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	common.SendResponse(c.Writer, http.StatusOK, "ok", tweet)
}
