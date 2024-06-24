package mappers

import (
	proto "github.com/fenek-dev/go-twitter/proto/protogen"
	"github.com/fenek-dev/go-twitter/src/common/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TweetModelToProtoTweet(tweet *models.Tweet) *proto.Tweet {
	return &proto.Tweet{
		Id:        tweet.ID,
		Content:   tweet.Content,
		Username:  tweet.Username,
		CreatedAt: &timestamppb.Timestamp{Seconds: tweet.CreatedAt.Unix()},
		UpdatedAt: &timestamppb.Timestamp{Seconds: tweet.UpdatedAt.Unix()},
	}
}
