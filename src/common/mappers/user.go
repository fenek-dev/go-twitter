package mappers

import (
	proto "github.com/fenek-dev/go-twitter/proto/protogen"
	"github.com/fenek-dev/go-twitter/src/common/models"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func UserModelToProtoUser(user *models.User) *proto.User {
	return &proto.User{
		Username:    user.Username,
		Description: user.Description,
		CreatedAt:   timestamppb.New(user.CreatedAt),
		UpdatedAt:   timestamppb.New(user.UpdatedAt),
	}
}

func ProtoUserToModel(user *proto.User) *models.User {
	return &models.User{
		Username:    user.Username,
		Description: user.Description,
		CreatedAt:   user.CreatedAt.AsTime(),
		UpdatedAt:   user.UpdatedAt.AsTime(),
	}
}

func ClaimsToUserModel(claims jwt.MapClaims) *models.User {

	createdAt, _ := time.Parse(time.RFC3339, claims["created_at"].(string))
	updatedAt, _ := time.Parse(time.RFC3339, claims["updated_at"].(string))

	return &models.User{
		Id:          claims["id"].(string),
		Username:    claims["username"].(string),
		Description: claims["description"].(string),
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
