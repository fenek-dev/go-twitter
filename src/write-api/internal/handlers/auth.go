package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/write-api/internal/dto"
	"go.opentelemetry.io/otel/attribute"
)

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	ctx, span := h.tracer.Start(r.Context(), "Register")
	defer span.End()
	var data dto.RegisterDto

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		common.SendResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	userAttr := attribute.String("username", data.Username)
	span.SetAttributes(userAttr)

	token, err := h.service.Register(ctx, data.Username, data.Password)
	if err != nil || token == "" {
		common.SendResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	tokenCookie := createTokenCookie(token)
	http.SetCookie(w, tokenCookie)

	common.SendResponse(w, http.StatusCreated, "ok", nil)
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	ctx, span := h.tracer.Start(r.Context(), "Login")
	defer span.End()
	var data dto.LoginDto

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		common.SendResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	userAttr := attribute.String("username", data.Username)
	span.SetAttributes(userAttr)

	token, err := h.service.Login(ctx, data.Username, data.Password)
	if err != nil || token == "" {
		common.SendResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	tokenCookie := createTokenCookie(token)
	http.SetCookie(w, tokenCookie)

	common.SendResponse(w, http.StatusCreated, "ok", token)
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
