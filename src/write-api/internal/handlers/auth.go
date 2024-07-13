package handlers

import (
	"net/http"

	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/write-api/internal/dto"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
)

func (h *Handlers) Register(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "write.handler.Register")
	defer span.End()
	var data dto.RegisterDto

	err := c.BindJSON(&data)
	if err != nil {
		common.SendResponse(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	span.SetAttributes(attribute.String("username", data.Username))

	token, err := h.service.Register(ctx, data.Username, data.Password)
	if err != nil || token == "" {
		common.SendResponse(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	tokenCookie := createTokenCookie(token)
	http.SetCookie(c.Writer, tokenCookie)

	common.SendResponse(c.Writer, http.StatusCreated, "ok", nil)
}

func (h *Handlers) Login(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "write.handler.Login")
	defer span.End()
	var data dto.LoginDto

	err := c.BindJSON(&data)
	if err != nil {
		common.SendResponse(c.Writer, http.StatusBadRequest, err.Error(), nil)
		return
	}

	span.SetAttributes(attribute.String("username", data.Username))

	token, err := h.service.Login(ctx, data.Username, data.Password)
	if err != nil || token == "" {
		common.SendResponse(c.Writer, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	tokenCookie := createTokenCookie(token)
	http.SetCookie(c.Writer, tokenCookie)

	common.SendResponse(c.Writer, http.StatusCreated, "ok", token)
}

func createTokenCookie(token string) *http.Cookie {
	return &http.Cookie{
		Name:     common.COOKIE_TOKEN_NAME,
		Value:    token,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
}
