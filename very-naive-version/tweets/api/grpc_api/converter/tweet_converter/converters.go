package tweet_converter

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"tweets/api/grpc_api/protobuf/tweets_service"
	"tweets/domain/models"
)

// Convert a database model to a Protobuf model.
func FromDatabaseModelToProtobufModel(tweet *models.Tweet) *tweets_service.Tweet {
	var repliesTo, quoteTo string

	if *tweet.RepliesTo == uuid.Nil {
		repliesTo = ""
	} else {
		repliesTo = tweet.RepliesTo.String()
	}

	if *tweet.QuoteTo == uuid.Nil {
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
	for _, tweet := range tweets {
		convertedTweets = append(convertedTweets, &tweets_service.Tweet{
			Id:        tweet.Id.String(),
			Text:      *tweet.Text,
			UserId:    tweet.UserId.String(),
			RepliesTo: tweet.RepliesTo.String(),
			QuoteTo:   tweet.QuoteTo.String(),
			CreatedAt: timestamppb.New(tweet.CreatedAt),
		})
	}

	return convertedTweets
}
