package grpc

import (
	"context"

	proto "github.com/fenek-dev/go-twitter/proto/protogen"
	"github.com/fenek-dev/go-twitter/src/cache/internal/storage/pg"
	"github.com/fenek-dev/go-twitter/src/cache/internal/storage/redis"
	"github.com/fenek-dev/go-twitter/src/common/mappers"

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

func (s *serverAPI) CreateTweet(ctx context.Context, in *proto.CreateTweetRequest) (*proto.CreateTweetResponse, error) {
	tweet, err := s.storage.CreateTweet(ctx, in.Username, in.Content)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &proto.CreateTweetResponse{Tweet: mappers.TweetModelToProtoTweet(&tweet)}, nil
}

func (s *serverAPI) DeleteTweet(ctx context.Context, in *proto.DeleteTweetRequest) (*proto.DeleteTweetResponse, error) {
	err := s.storage.DeleteTweet(ctx, in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &proto.DeleteTweetResponse{Id: in.Id}, nil
}

func (s *serverAPI) FindTweetById(ctx context.Context, in *proto.FindTweetByIdRequest) (*proto.FindTweetByIdResponse, error) {
	tweet, err := s.storage.FindTweetById(ctx, in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &proto.FindTweetByIdResponse{Tweet: mappers.TweetModelToProtoTweet(&tweet)}, nil
}

func (s *serverAPI) FindUserById(ctx context.Context, in *proto.FindUserByIdRequest) (*proto.FindUserByIdResponse, error) {
	user, err := s.storage.FindUserById(ctx, in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &proto.FindUserByIdResponse{User: mappers.UserModelToProtoUser(&user)}, nil
}

func (s *serverAPI) SaveUser(ctx context.Context, in *proto.SaveUserRequest) (*proto.SaveUserResponse, error) {
	id, err := s.storage.SaveUser(ctx, in.Username, in.PassHash)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &proto.SaveUserResponse{UserId: id}, nil
}

func (s *serverAPI) UpdateTweet(ctx context.Context, in *proto.UpdateTweetRequest) (*proto.UpdateTweetResponse, error) {
	tweet, err := s.storage.UpdateTweet(ctx, in.Id, in.Content)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &proto.UpdateTweetResponse{Tweet: mappers.TweetModelToProtoTweet(&tweet)}, nil
}
