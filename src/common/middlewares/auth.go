package middlewares

import (
	"context"
	"net/http"
	"strings"

	ssov1 "github.com/fenek-dev/go-twitter/proto/protogen"
	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/mappers"
)

type Auth struct {
	sso ssov1.AuthServiceClient
}

func NewAuthMiddleware(sso ssov1.AuthServiceClient) *Auth {
	return &Auth{sso: sso}
}

func (a *Auth) Handle(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")

		if len(t) != 2 {
			common.SendResponse(w, http.StatusUnauthorized, "Missing auth token", nil)
			return
		}
		token := t[1]

		res, err := a.sso.Verify(r.Context(), &ssov1.VerifyRequest{Token: token})
		if err != nil {
			common.SendResponse(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		user := mappers.ProtoUserToModel(res.User)

		ctx := context.WithValue(r.Context(), common.REQUEST_CTX_USER, user)

		req := r.WithContext(ctx)

		f(w, req)
	})
}
