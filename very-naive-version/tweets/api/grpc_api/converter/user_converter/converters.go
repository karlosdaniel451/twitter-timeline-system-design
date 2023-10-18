package user_converter

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"tweets/api/grpc_api/protobuf/users_service"
	"tweets/domain/models"
	"tweets/utils"
)

func FromSignUpRequestToProtobufModel(request *users_service.SignUpRequest) *users_service.User {
	return &users_service.User{
		Name:            request.GetName(),
		UserName:        request.GetUserName(),
		Email:           request.GetEmail(),
		Description:     request.GetDescription(),
		Location:        request.GetLocation(),
		PinnedTweet:     request.GetPinnedTweet(),
		ProfileImageUrl: request.GetProfileImageUrl(),
		Protected:       request.GetProtected(),
		PublicMetrics:   nil,
		Url:             request.GetUrl(),
		Verified:        request.GetVerified(),
	}
}

func FromProtobufModelToDatabaseModel(user *users_service.User) *models.User {
	return &models.User{
		Name:            utils.ValueToPointer[string](user.GetName()),
		Username:        utils.ValueToPointer[string](user.GetUserName()),
		Email:           utils.ValueToPointer[string](user.GetEmail()),
		Description:     utils.ValueToPointer[string](user.GetDescription()),
		Location:        utils.ValueToPointer[string](user.GetLocation()),
		PinnedTweet:     nil,
		ProfileImageUrl: utils.ValueToPointer[string](user.GetProfileImageUrl()),
		Protected:       utils.ValueToPointer[bool](user.GetProtected()),
		PublicMetrics: &models.PublicMetrics{
			FollowersCount: user.GetPublicMetrics().GetFollowersCount(),
			FollowingCount: user.GetPublicMetrics().GetFollowingCount(),
			TweetCount:     user.GetPublicMetrics().GetTweetCount(),
			ListedCount:    user.GetPublicMetrics().GetListedCount(),
		},
		Url:             utils.ValueToPointer[string](user.GetUrl()),
		Verified:        utils.ValueToPointer[bool](user.GetVerified()),
		MostRecentTweet: nil,
		Followers:       nil,
		FollowerIds:     nil,
		Followees:       nil,
		FolloweeIds:     nil,
		Tweets:          nil,
	}
}

// Convert a database model to a Protobuf model.
func FromDatabaseModelToProtobufModel(user *models.User) *users_service.User {
	var pinnedTweetId, mostRecentTweetId string

	if user.PinnedTweet == nil {
		pinnedTweetId = ""
	} else {
		pinnedTweetId = user.PinnedTweet.Id.String()
	}

	if user.MostRecentTweet == nil {
		mostRecentTweetId = ""
	} else {
		mostRecentTweetId = user.MostRecentTweet.Id.String()
	}

	return &users_service.User{
		Id:              user.Id.String(),
		Name:            *user.Name,
		UserName:        *user.Username,
		Email:           *user.Email,
		Description:     *user.Description,
		Location:        *user.Location,
		PinnedTweet:     pinnedTweetId,
		ProfileImageUrl: *user.ProfileImageUrl,
		Protected:       *user.Protected,
		PublicMetrics: &users_service.PublicMetrics{
			FollowersCount: user.PublicMetrics.FollowersCount,
			FollowingCount: user.PublicMetrics.FollowingCount,
			TweetCount:     user.PublicMetrics.TweetCount,
			ListedCount:    user.PublicMetrics.ListedCount,
		},
		Url:             *user.Url,
		Verified:        *user.Verified,
		MostRecentTweet: mostRecentTweetId,
		FollowedUserIds: utils.MapStringer[uuid.UUID](user.FolloweeIds),
		FollowerUserIds: utils.MapStringer[uuid.UUID](user.FollowerIds),
		CreatedAt:       timestamppb.New(user.CreatedAt),
		UpdatedAt:       timestamppb.New(user.UpdatedAt),
		DeletedAt:       timestamppb.New(user.DeletedAt.Time),
	}
}

// Convert a slice of database models to a slice of Protobuf models.
func FromDatabaseModelsToProtobufModels(users []*models.User) []*users_service.User {
	convertedUsers := make([]*users_service.User, 0, len(users))
	for _, user := range users {
		convertedUsers = append(convertedUsers, FromDatabaseModelToProtobufModel(user))
	}

	return convertedUsers
}

func FromProtobufModelsToDatabaseModels(users []*users_service.User) []*models.User {
	convertedUsers := make([]*models.User, 0, len(users))
	for _, user := range users {
		convertedUsers = append(convertedUsers, FromProtobufModelToDatabaseModel(user))
	}

	return convertedUsers
}
