package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_ssov1 "github.com/fenek-dev/go-twitter/proto/protogen/mocks"
	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandlers_Me(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := mock_ssov1.NewMockCacheServiceClient(ctrl)

	handlers := New(mock)
	handler := http.HandlerFunc(handlers.Me)

	t.Run("Something gone wrong", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/api/v1/me", nil)
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, http.StatusInternalServerError, "handler returned wrong status code")

		res := common.NewResponse(http.StatusInternalServerError, "Something gone wrong", nil)

		expected, err := json.Marshal(res)
		if err != nil {
			t.Fatalf("could not marshal response: %v", err)
		}

		assert.Equal(t, rr.Body.String(), string(expected), "handler returned unexpected body")
	})

	t.Run("Failed to find tweet", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/api/v1/me", nil)
		rr := httptest.NewRecorder()

		user := models.User{
			Username:    "test",
			Description: "test description",
		}

		ctx := context.WithValue(req.Context(), common.REQUEST_CTX_USER, user)
		req = req.WithContext(ctx)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, http.StatusOK, "handler returned wrong status code")

		res := common.NewResponse(http.StatusOK, "ok", user)

		expected, err := json.Marshal(res)
		if err != nil {
			t.Fatalf("could not marshal response: %v", err)
		}

		assert.Equal(t, rr.Body.String(), string(expected), "handler returned unexpected body")
	})

}
