package middlewares

import (
	"net/http"

	ssov1 "github.com/fenek-dev/go-twitter/proto/protogen"
	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/mappers"
	"github.com/fenek-dev/go-twitter/src/common/models"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	sso ssov1.AuthServiceClient
}

func NewAuthMiddleware(sso ssov1.AuthServiceClient) *Auth {
	return &Auth{sso: sso}
}

func (a *Auth) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := c.Cookie(common.COOKIE_TOKEN_NAME)
		if err != nil {
			common.SendResponse(c.Writer, http.StatusUnauthorized, "Missing auth token", nil)
			return
		}

		res, err := a.sso.Verify(c.Request.Context(), &ssov1.VerifyRequest{Token: token})
		if err != nil {
			common.SendResponse(c.Writer, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		user := mappers.ProtoUserToModel(res.User)

		c.Set(common.REQUEST_CTX_USER, user)
		c.Next()
	}
}

func UserFromCtx(c *gin.Context) (*models.User, bool) {
	u, ok := c.Get(common.REQUEST_CTX_USER)
	if !ok {
		return nil, ok
	}
	user, ok := u.(*models.User)
	return user, ok
}
