package tweet_converter

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"tweets/api/grpc_api/protobuf/tweets_service"
	"tweets/domain/models"
	"tweets/utils"
)

func FromPostTweetRequestToDatabaseModel(
	request *tweets_service.PostTweetRequest,
) *models.Tweet {

	// Required fields that require conversion.
	var (
		userId = &uuid.Nil
	)

	*userId = uuid.MustParse(request.GetUserId())

	// Nullable fields that require conversion.
	var (
		repliesTo *uuid.UUID
		quoteTo   *uuid.UUID
	)

	if request.GetRepliesTo() == "" {
		repliesTo = nil
	} else {
		*repliesTo = uuid.MustParse(request.GetRepliesTo())
	}

	if request.GetQuoteTo() == "" {
		quoteTo = nil
	} else {
		*quoteTo = uuid.MustParse(request.GetQuoteTo())
	}

	return &models.Tweet{
		Text:      utils.ValueToPointer[string](request.GetText()),
		UserId:    userId,
		RepliesTo: repliesTo,
		QuoteTo:   quoteTo,
	}
}

// Convert a database model to a Protobuf model.
func FromDatabaseModelToProtobufModel(tweet *models.Tweet) *tweets_service.Tweet {
	var repliesTo, quoteTo string

	if tweet.RepliesTo == nil {
		repliesTo = ""
	} else {
		repliesTo = tweet.RepliesTo.String()
	}

	if tweet.QuoteTo == nil {
		quoteTo = ""
	} else {
		quoteTo = tweet.QuoteTo.String()
	}

	return &tweets_service.Tweet{
		Id:        tweet.Id.String(),
		Text:      *tweet.Text,
		UserId:    tweet.UserId.String(),
		RepliesTo: repliesTo,
		QuoteTo:   quoteTo,
		CreatedAt: timestamppb.New(tweet.CreatedAt),
	}
}

// Convert a slice of database models to a slice of Protobuf models.
func FromDatabaseModelsToProtobufModels(tweets []*models.Tweet) []*tweets_service.Tweet {
	convertedTweets := make([]*tweets_service.Tweet, 0, len(tweets))

	// Required fields that require conversion.
	var (
		userId string
	)

	// Nullable fields that require conversion.
	var (
		repliesTo string
		quoteTo   string
	)

	for _, tweet := range tweets {
		if tweet.UserId != nil {
			userId = tweet.UserId.String()
		}
		if tweet.RepliesTo != nil {
			repliesTo = tweet.RepliesTo.String()
		}
		if tweet.QuoteTo != nil {
			quoteTo = tweet.QuoteTo.String()
		}

		convertedTweets = append(convertedTweets, &tweets_service.Tweet{
			Id:        tweet.Id.String(),
			Text:      *tweet.Text,
			UserId:    userId,
			RepliesTo: repliesTo,
			QuoteTo:   quoteTo,
			CreatedAt: timestamppb.New(tweet.CreatedAt),
		})
	}

	return convertedTweets
}
