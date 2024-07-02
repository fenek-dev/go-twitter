package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	ssov1 "github.com/fenek-dev/go-twitter/proto/protogen"
	mock_ssov1 "github.com/fenek-dev/go-twitter/proto/protogen/mocks"
	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHandlers_FindTweetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := mock_ssov1.NewMockCacheServiceClient(ctrl)

	handlers := New(mock)
	handler := http.HandlerFunc(handlers.FindTweetById)

	t.Run("Without id", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/api/v1/tweet/", nil)
		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, http.StatusBadRequest, "handler returned wrong status code")

		res := common.Response{
			Code:    http.StatusBadRequest,
			Message: "incorrect_id",
			Data:    nil,
		}

		expected, err := json.Marshal(res)
		if err != nil {
			t.Fatalf("could not marshal response: %v", err)
		}

		assert.Equal(t, rr.Body.String(), string(expected), "handler returned unexpected body")
	})

	t.Run("Failed to find tweet", func(t *testing.T) {

		id := "2"
		req := httptest.NewRequest(http.MethodGet, "/api/v1/tweet/"+id, nil)
		rr := httptest.NewRecorder()
		req.SetPathValue("id", id)

		responseError := errors.New("failed to find tweet")
		mock.EXPECT().FindTweetById(gomock.Any(), gomock.Eq(&ssov1.FindTweetByIdRequest{Id: id})).Return(nil, responseError)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, http.StatusInternalServerError, "handler returned wrong status code")

		res := common.Response{
			Code:    http.StatusInternalServerError,
			Message: responseError.Error(),
			Data:    nil,
		}

		expected, err := json.Marshal(res)
		if err != nil {
			t.Fatalf("could not marshal response: %v", err)
		}

		assert.Equal(t, rr.Body.String(), string(expected), "handler returned unexpected body")
	})

	t.Run("Success", func(t *testing.T) {
		id := "1"
		req := httptest.NewRequest(http.MethodGet, "/api/v1/tweet/"+id, nil)
		rr := httptest.NewRecorder()
		req.SetPathValue("id", id)
		response := &ssov1.FindTweetByIdResponse{
			Tweet: &ssov1.Tweet{
				Id:       id,
				Content:  "content",
				Username: "username",
			},
		}

		mock.EXPECT().FindTweetById(gomock.Any(), gomock.Eq(&ssov1.FindTweetByIdRequest{Id: id})).Return(response, nil)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, http.StatusOK, "handler returned wrong status code")

		res := common.Response{
			Code:    http.StatusOK,
			Message: "ok",
			Data:    response,
		}

		expected, err := json.Marshal(res)
		if err != nil {
			t.Fatalf("could not marshal response: %v", err)
		}

		assert.Equal(t, rr.Body.String(), string(expected), "handler returned unexpected body")

	})
}
