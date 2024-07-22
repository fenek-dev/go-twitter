package handlers

import (
	"net/http"

	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
)

func (h *Handlers) FindUserById(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "user.handler.FindUserById")
	defer span.End()
	id := c.Param("id")

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
