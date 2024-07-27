package services

import (
	"context"
	"errors"
	"github.com/fenek-dev/go-twitter/src/common/models"
	common_test "github.com/fenek-dev/go-twitter/src/common/test"
	"github.com/fenek-dev/go-twitter/src/tweet/mocks"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func TestServices_CreateTweet(t *testing.T) {
	type args struct {
		username string
		content  string
	}
	type pgResponse struct {
		tweet *models.Tweet
		err   error
	}

	tweet := &models.Tweet{
		ID:       "1",
		Username: "test",
		Content:  "something very important",
	}

	tests := []struct {
		name         string
		args         args
		want         *models.Tweet
		wantErr      bool
		shouldCallPg bool
		pgResponse   pgResponse
	}{
		{
			name: "Success",
			args: args{
				username: "test",
				content:  "something very important",
			},
			want:         tweet,
			wantErr:      false,
			shouldCallPg: true,
			pgResponse: pgResponse{
				tweet: tweet,
				err:   nil,
			},
		},
		{
			name: "Error from pg",
			args: args{
				username: "test",
				content:  "something very important",
			},
			want:         nil,
			wantErr:      true,
			shouldCallPg: true,
			pgResponse: pgResponse{
				tweet: nil,
				err:   errors.New("error from pg"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			pg := mocks.NewPostgres(t)
			rdb := mocks.NewRedis(t)
			tracer, shutdown := common_test.NewTestTracer(t)
			defer shutdown(ctx)

			if tt.shouldCallPg {
				pg.On("CreateTweet", mock.Anything, tt.args.username, tt.args.content).Return(tt.pgResponse.tweet, tt.pgResponse.err)
			}

			s := &Services{
				pg:     pg,
				rdb:    rdb,
				tracer: tracer,
			}

			got, err := s.CreateTweet(ctx, tt.args.username, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTweet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTweet() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServices_FindTweetById(t *testing.T) {
	type args struct {
		id string
	}
	type pgResponse struct {
		tweet *models.Tweet
		err   error
	}

	tweet := &models.Tweet{
		ID:       "1",
		Username: "test",
		Content:  "something very important",
	}

	tests := []struct {
		name         string
		args         args
		want         *models.Tweet
		wantErr      bool
		shouldCallPg bool
		pgResponse   pgResponse
	}{
		{
			name: "Success",
			args: args{
				id: "1",
			},
			want:         tweet,
			wantErr:      false,
			shouldCallPg: true,
			pgResponse: pgResponse{
				tweet: tweet,
				err:   nil,
			},
		},
		{
			name: "Error from pg",
			args: args{
				id: "1",
			},
			want:         nil,
			wantErr:      true,
			shouldCallPg: true,
			pgResponse: pgResponse{
				tweet: nil,
				err:   errors.New("error from pg"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			pg := mocks.NewPostgres(t)
			rdb := mocks.NewRedis(t)
			tracer, shutdown := common_test.NewTestTracer(t)
			defer shutdown(ctx)
			if tt.shouldCallPg {
				pg.On("FindTweetById", mock.Anything, tt.args.id).Return(tt.pgResponse.tweet, tt.pgResponse.err)
			}

			s := &Services{
				pg:     pg,
				rdb:    rdb,
				tracer: tracer,
			}

			got, err := s.FindTweetById(ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindTweetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindTweetById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServices_UpdateTweet(t *testing.T) {
	type args struct {
		id      string
		content string
	}
	type pgResponse struct {
		tweet *models.Tweet
		err   error
	}

	tweet := &models.Tweet{
		ID:       "1",
		Username: "test",
		Content:  "something very important",
	}

	tests := []struct {
		name         string
		args         args
		want         *models.Tweet
		wantErr      bool
		shouldCallPg bool
		pgResponse   pgResponse
	}{
		{
			name: "Success",
			args: args{
				id:      "1",
				content: "something very important",
			},
			want:         tweet,
			wantErr:      false,
			shouldCallPg: true,
			pgResponse: pgResponse{
				tweet: tweet,
				err:   nil,
			},
		},
		{
			name: "Error from pg",
			args: args{
				id:      "1",
				content: "something very important",
			},
			want:         nil,
			wantErr:      true,
			shouldCallPg: true,
			pgResponse: pgResponse{
				tweet: nil,
				err:   errors.New("error from pg"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			pg := mocks.NewPostgres(t)
			rdb := mocks.NewRedis(t)
			tracer, shutdown := common_test.NewTestTracer(t)
			defer shutdown(ctx)
			if tt.shouldCallPg {
				pg.On("UpdateTweet", mock.Anything, tt.args.id, tt.args.content).Return(tt.pgResponse.tweet, tt.pgResponse.err)
			}

			s := &Services{
				pg:     pg,
				rdb:    rdb,
				tracer: tracer,
			}

			got, err := s.UpdateTweet(ctx, tt.args.id, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindTweetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindTweetById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServices_DeleteTweet(t *testing.T) {
	type args struct {
		id string
	}
	type pgResponse struct {
		id  string
		err error
	}
	tests := []struct {
		name         string
		args         args
		want         string
		wantErr      bool
		shouldCallPg bool
		pgResponse   pgResponse
	}{
		{
			name: "Success",
			args: args{
				id: "1",
			},
			want:         "1",
			wantErr:      false,
			shouldCallPg: true,
			pgResponse: pgResponse{
				err: nil,
			},
		},
		{
			name: "Error from pg",
			args: args{
				id: "1",
			},
			want:         "",
			wantErr:      true,
			shouldCallPg: true,
			pgResponse: pgResponse{
				err: errors.New("error from pg"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			pg := mocks.NewPostgres(t)
			rdb := mocks.NewRedis(t)
			tracer, shutdown := common_test.NewTestTracer(t)
			defer shutdown(ctx)
			if tt.shouldCallPg {
				pg.On("DeleteTweet", mock.Anything, tt.args.id).Return(tt.pgResponse.err)
			}

			s := &Services{
				pg:     pg,
				rdb:    rdb,
				tracer: tracer,
			}

			got, err := s.DeleteTweet(ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteTweet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeleteTweet() got = %v, want %v", got, tt.want)
			}
		})
	}
}
