package mappers

import (
	proto "github.com/fenek-dev/go-twitter/proto/protogen"
	"github.com/fenek-dev/go-twitter/src/common/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserModelToProtoUser(user *models.User) *proto.User {
	return &proto.User{
		Username:    user.Username,
		PassHash:    user.PassHash,
		Description: user.Description,
		CreatedAt:   timestamppb.New(user.CreatedAt),
		UpdatedAt:   timestamppb.New(user.UpdatedAt),
	}
}

func ProtoUserToModel(user *proto.User) *models.User {
	return &models.User{
		Username:    user.Username,
		PassHash:    user.PassHash,
		Description: user.Description,
		CreatedAt:   user.CreatedAt.AsTime(),
		UpdatedAt:   user.UpdatedAt.AsTime(),
	}
}
