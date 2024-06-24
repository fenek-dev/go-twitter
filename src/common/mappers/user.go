package mappers

import (
	proto "github.com/fenek-dev/go-twitter/proto/protogen"
	"github.com/fenek-dev/go-twitter/src/common/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserModelToProtoTweet(user *models.User) *proto.User {
	return &proto.User{
		Username:    user.Username,
		PassHash:    user.PassHash,
		Description: user.Description,
		CreatedAt:   &timestamppb.Timestamp{Seconds: user.CreatedAt.Unix()},
		UpdatedAt:   &timestamppb.Timestamp{Seconds: user.UpdatedAt.Unix()},
	}
}
