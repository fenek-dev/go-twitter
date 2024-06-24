package grpc

import (
	"context"

	proto "github.com/fenek-dev/go-twitter/proto/protogen"
	"github.com/fenek-dev/go-twitter/src/cache/internal/storage/pg"
	"github.com/fenek-dev/go-twitter/src/cache/internal/storage/redis"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	proto.UnimplementedCacheServiceServer
	storage *pg.Postgres
	redis   *redis.Redis
}

func Register(gRPCServer *grpc.Server, storage *pg.Postgres, redis *redis.Redis) {
	proto.RegisterCacheServiceServer(gRPCServer, &serverAPI{storage: storage, redis: redis})
}

func (s *serverAPI) CreateTweet(context.Context, *proto.CreateTweetRequest) (*proto.CreateTweetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTweet not implemented")
}
func (s *serverAPI) DeleteTweet(context.Context, *proto.DeleteTweetRequest) (*proto.DeleteTweetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTweet not implemented")
}
func (s *serverAPI) FindTweetById(context.Context, *proto.FindTweetByIdRequest) (*proto.FindTweetByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTweetById not implemented")
}
func (s *serverAPI) FindUserById(context.Context, *proto.FindUserByIdRequest) (*proto.FindUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserById not implemented")
}
func (s *serverAPI) SaveUser(context.Context, *proto.SaveUserRequest) (*proto.SaveUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUser not implemented")
}
func (s *serverAPI) UpdateTweet(context.Context, *proto.UpdateTweetRequest) (*proto.UpdateTweetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTweet not implemented")
}
