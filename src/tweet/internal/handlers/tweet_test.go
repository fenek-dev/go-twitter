package handlers

import (
	"context"
	"errors"
	"github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/models"
	common_test "github.com/fenek-dev/go-twitter/src/common/test"
	"github.com/fenek-dev/go-twitter/src/tweet/internal/dto"
	"github.com/fenek-dev/go-twitter/src/tweet/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers_CreateTweet(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tweet := &models.Tweet{
		Username: "test",
		Content:  "very important tweet",
	}

	type srvResponse struct {
		tweet *models.Tweet
		err   error
	}

	tests := []struct {
		name        string
		want        string
		status      int
		ctxUser     *models.User
		body        *dto.CreateDto
		srvResponse srvResponse
	}{
		{
			name:   "create tweet",
			want:   common_test.NewStringResponse(http.StatusCreated, "ok", tweet),
			status: http.StatusCreated,
			ctxUser: &models.User{
				Username: "test",
			},
			body: &dto.CreateDto{
				Username: tweet.Username,
				Content:  tweet.Content,
			},
			srvResponse: srvResponse{
				tweet: tweet,
				err:   nil,
			},
		},
		{
			name:        "can not get user from ctx",
			want:        common_test.NewStringResponse(http.StatusInternalServerError, ErrCanNotGetUser.Error(), nil),
			status:      http.StatusInternalServerError,
			ctxUser:     nil,
			body:        &dto.CreateDto{},
			srvResponse: srvResponse{},
		},
		{
			name:   "error from service",
			want:   common_test.NewStringResponse(http.StatusInternalServerError, "error from service", nil),
			status: http.StatusInternalServerError,
			ctxUser: &models.User{
				Username: "test",
			},
			body: &dto.CreateDto{
				Username: tweet.Username,
				Content:  tweet.Content,
			},
			srvResponse: srvResponse{
				tweet: nil,
				err:   errors.New("error from service"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			service := mocks.NewService(t)
			log := common_test.NewTestLogger(t)
			tracer, shutdown := common_test.NewTestTracer(t)
			defer shutdown(ctx)

			service.On("CreateTweet", mock.Anything, mock.Anything, mock.Anything).Return(tt.srvResponse.tweet, tt.srvResponse.err).Maybe()

			h := &Handlers{
				service: service,
				log:     log,
				tracer:  tracer,
			}

			w := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(w)
			body := common_test.NewBodyBuffer(t, &tt.body)
			c.Request = httptest.NewRequest(http.MethodPost, "/tweet", body)

			c.Set(common.REQUEST_CTX_USER, tt.ctxUser)

			h.CreateTweet(c)

			responseData, err := io.ReadAll(w.Body)
			if err != nil {
				t.Fatalf("Error reading response body: %v", err)
			}

			assert.Equal(t, tt.status, w.Code)
			assert.Equal(t, tt.want, string(responseData))

		})
	}
}

func TestHandlers_DeleteTweet(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tweet := &models.Tweet{
		ID:       "2",
		Username: "test",
		Content:  "very important tweet",
	}

	type srvResponse struct {
		id  string
		err error
	}

	tests := []struct {
		name        string
		want        string
		status      int
		ctxUser     *models.User
		body        *dto.DeleteDto
		srvResponse srvResponse
	}{
		{
			name:   "delete tweet",
			want:   common_test.NewStringResponse(http.StatusOK, "ok", tweet.ID),
			status: http.StatusOK,
			ctxUser: &models.User{
				Username: "test",
			},
			body: &dto.DeleteDto{
				Id: tweet.ID,
			},
			srvResponse: srvResponse{
				id:  tweet.ID,
				err: nil,
			},
		},
		{
			name:   "error from service",
			want:   common_test.NewStringResponse(http.StatusInternalServerError, "error from service", nil),
			status: http.StatusInternalServerError,
			ctxUser: &models.User{
				Username: "test",
			},
			body: &dto.DeleteDto{
				Id: tweet.ID,
			},
			srvResponse: srvResponse{
				id:  "",
				err: errors.New("error from service"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			service := mocks.NewService(t)
			log := common_test.NewTestLogger(t)
			tracer, shutdown := common_test.NewTestTracer(t)
			defer shutdown(ctx)

			service.On("DeleteTweet", mock.Anything, mock.Anything).Return(tt.srvResponse.id, tt.srvResponse.err).Maybe()

			h := &Handlers{
				service: service,
				log:     log,
				tracer:  tracer,
			}

			w := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(w)
			body := common_test.NewBodyBuffer(t, &tt.body)
			c.Request = httptest.NewRequest(http.MethodDelete, "/tweet", body)

			c.Set(common.REQUEST_CTX_USER, tt.ctxUser)

			h.DeleteTweet(c)

			responseData, err := io.ReadAll(w.Body)
			if err != nil {
				t.Fatalf("Error reading response body: %v", err)
			}

			assert.Equal(t, tt.status, w.Code)
			assert.Equal(t, tt.want, string(responseData))

		})
	}
}

func TestHandlers_FindTweetById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tweet := &models.Tweet{
		ID:       "2",
		Username: "test",
		Content:  "very important tweet",
	}

	type srvResponse struct {
		tweet *models.Tweet
		err   error
	}

	tests := []struct {
		name        string
		want        string
		status      int
		ctxUser     *models.User
		id          string
		srvResponse srvResponse
	}{
		{
			name:    "create tweet",
			want:    common_test.NewStringResponse(http.StatusOK, "ok", tweet),
			status:  http.StatusOK,
			ctxUser: &models.User{},
			id:      tweet.ID,
			srvResponse: srvResponse{
				tweet: tweet,
				err:   nil,
			},
		},
		{
			name:   "error from service",
			want:   common_test.NewStringResponse(http.StatusInternalServerError, "error from service", nil),
			status: http.StatusInternalServerError,
			ctxUser: &models.User{
				Username: "test",
			},
			id: tweet.ID,
			srvResponse: srvResponse{
				tweet: nil,
				err:   errors.New("error from service"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			service := mocks.NewService(t)
			log := common_test.NewTestLogger(t)
			tracer, shutdown := common_test.NewTestTracer(t)
			defer shutdown(ctx)

			service.On("FindTweetById", mock.Anything, mock.Anything, mock.Anything).Return(tt.srvResponse.tweet, tt.srvResponse.err).Maybe()

			h := &Handlers{
				service: service,
				log:     log,
				tracer:  tracer,
			}

			w := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(w)
			c.AddParam("id", tt.id)
			c.Request = httptest.NewRequest(http.MethodGet, "/tweet/"+tt.id, nil)

			c.Set(common.REQUEST_CTX_USER, tt.ctxUser)

			h.FindTweetById(c)

			responseData, err := io.ReadAll(w.Body)
			if err != nil {
				t.Fatalf("Error reading response body: %v", err)
			}

			assert.Equal(t, tt.status, w.Code)
			assert.Equal(t, tt.want, string(responseData))

		})
	}
}

func TestHandlers_UpdateTweet(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tweet := &models.Tweet{
		ID:       "2",
		Username: "test",
		Content:  "very important tweet",
	}

	type srvResponse struct {
		tweet *models.Tweet
		err   error
	}

	tests := []struct {
		name        string
		want        string
		status      int
		ctxUser     *models.User
		body        *dto.UpdateDto
		srvResponse srvResponse
	}{
		{
			name:   "create tweet",
			want:   common_test.NewStringResponse(http.StatusOK, "ok", tweet),
			status: http.StatusOK,
			ctxUser: &models.User{
				Username: "test",
			},
			body: &dto.UpdateDto{
				Id:      tweet.ID,
				Content: tweet.Content,
			},
			srvResponse: srvResponse{
				tweet: tweet,
				err:   nil,
			},
		},
		{
			name:   "error from service",
			want:   common_test.NewStringResponse(http.StatusInternalServerError, "error from service", nil),
			status: http.StatusInternalServerError,
			ctxUser: &models.User{
				Username: "test",
			},
			body: &dto.UpdateDto{
				Id:      tweet.ID,
				Content: tweet.Content,
			},
			srvResponse: srvResponse{
				tweet: nil,
				err:   errors.New("error from service"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			service := mocks.NewService(t)
			log := common_test.NewTestLogger(t)
			tracer, shutdown := common_test.NewTestTracer(t)
			defer shutdown(ctx)

			service.On("UpdateTweet", mock.Anything, mock.Anything, mock.Anything).Return(tt.srvResponse.tweet, tt.srvResponse.err).Maybe()

			h := &Handlers{
				service: service,
				log:     log,
				tracer:  tracer,
			}

			w := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(w)
			body := common_test.NewBodyBuffer(t, &tt.body)
			c.Request = httptest.NewRequest(http.MethodPost, "/tweet", body)

			c.Set(common.REQUEST_CTX_USER, tt.ctxUser)

			h.UpdateTweet(c)

			responseData, err := io.ReadAll(w.Body)
			if err != nil {
				t.Fatalf("Error reading response body: %v", err)
			}

			assert.Equal(t, tt.status, w.Code)
			assert.Equal(t, tt.want, string(responseData))

		})
	}
}
