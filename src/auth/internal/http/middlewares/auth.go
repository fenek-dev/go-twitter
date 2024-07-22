package middlewares

import (
	"github.com/fenek-dev/go-twitter/src/auth/internal/services"
	"net/http"

	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/gin-gonic/gin"
)

func Auth(srv *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := c.Cookie(common.COOKIE_TOKEN_NAME)
		if err != nil {
			common.SendResponse(c.Writer, http.StatusUnauthorized, "Missing auth token", nil)
			return
		}

		user, err := srv.Verify(c.Request.Context(), token)
		if err != nil {
			common.SendResponse(c.Writer, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.Set(common.REQUEST_CTX_USER, user)
		c.Next()
	}
}
