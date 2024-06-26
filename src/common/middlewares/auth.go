package middlewares

import (
	"context"
	"net/http"

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

		token, err := r.Cookie(common.COOKIE_TOKEN_NAME)
		if err != nil {
			common.SendResponse(w, http.StatusUnauthorized, "Missing auth token", nil)
			return
		}

		res, err := a.sso.Verify(r.Context(), &ssov1.VerifyRequest{Token: token.Value})
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
