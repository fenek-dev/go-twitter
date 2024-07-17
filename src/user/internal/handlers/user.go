package handlers

import (
	"net/http"

	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/models"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
)

func (h *Handlers) Me(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "read.handler.Me")
	defer span.End()
	user, ok := ctx.Value(common.REQUEST_CTX_USER).(models.User)
	if !ok {
		common.SendResponse(c.Writer, http.StatusInternalServerError, "Something gone wrong", nil)
		return
	}
	span.SetAttributes(attribute.String("id", user.Username))

	common.SendResponse(c.Writer, http.StatusOK, "ok", user)
}

func (h *Handlers) FindUserById(c *gin.Context) {
	id := c.Param("id")
	ctx, span := h.tracer.Start(c.Request.Context(), "read.handler.FindUserById")
	defer span.End()

	span.SetAttributes(attribute.String("id", id))

	if id == "" {
		common.SendResponse(c.Writer, http.StatusBadRequest, "incorrect_id", nil)
		return
	}

	user, err := h.service.FindUserById(ctx, id)
	if err != nil {
		common.SendResponse(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	span.SetAttributes(attribute.String("username", user.Username))

	common.SendResponse(c.Writer, http.StatusOK, "ok", user)
}
