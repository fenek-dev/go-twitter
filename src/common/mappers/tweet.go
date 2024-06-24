package mappers

import (
	proto "github.com/fenek-dev/go-twitter/proto/protogen"
	"github.com/fenek-dev/go-twitter/src/common/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TweetModelToProtoTweet(tweet *models.Tweet) *proto.Tweet {
	timestamppb.New(tweet.CreatedAt)
	return &proto.Tweet{
		Id:        tweet.ID,
		Content:   tweet.Content,
		Username:  tweet.Username,
		CreatedAt: timestamppb.New(tweet.CreatedAt),
		UpdatedAt: timestamppb.New(tweet.UpdatedAt),
	}
}

func ProtoTweetToModel(tweet *proto.Tweet) *models.Tweet {
	return &models.Tweet{
		ID:        tweet.Id,
		Content:   tweet.Content,
		Username:  tweet.Username,
		CreatedAt: tweet.CreatedAt.AsTime(),
		UpdatedAt: tweet.UpdatedAt.AsTime(),
	}
}
