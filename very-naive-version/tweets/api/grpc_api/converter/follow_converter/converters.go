package follow_converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"tweets/api/grpc_api/protobuf/users_service"
	"tweets/domain/models"
)

// FromDatabaseModelToProtobufModel Convert a database model to a Protobuf model.
func FromDatabaseModelToProtobufModel(follow *models.Follow) *users_service.Follow {
	return &users_service.Follow{
		FollowerId: follow.FollowerId.String(),
		FolloweeId: follow.FolloweeId.String(),
		CreatedAt:  timestamppb.New(follow.CreatedAt),
		UpdatedAt:  timestamppb.New(follow.UpdatedAt),
		DeletedAt:  timestamppb.New(follow.DeletedAt.Time),
	}
}
