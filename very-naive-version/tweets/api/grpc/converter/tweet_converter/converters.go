package tweet_converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"tweets/api/grpc/controller/protobuf/tweets_service"
	"tweets/domain/models"
)

// Convert a database model to a Protobuf model.
func FromDatabaseModelToProtobufModel(tweet *models.Tweet) *tweets_service.Tweet {
	return &tweets_service.Tweet{
		Id:        tweet.Id.String(),
		Text:      tweet.Text,
		UserId:    tweet.UserId.String(),
		RepliesTo: tweet.RepliesTo.String(),
		QuoteTo:   tweet.QuoteTo.String(),
		CreatedAt: timestamppb.New(tweet.CreatedAt),
	}
}

// Convert a slice of database models to a slice of Protobuf models.
func FromDatabaseModelsToProtobufModels(tweets []*models.Tweet) []*tweets_service.Tweet {
	convertedTweets := make([]*tweets_service.Tweet, 0, len(tweets))
	for _, tweet := range tweets {
		convertedTweets = append(convertedTweets, &tweets_service.Tweet{
			Id:        tweet.Id.String(),
			Text:      tweet.Text,
			UserId:    tweet.UserId.String(),
			RepliesTo: tweet.RepliesTo.String(),
			QuoteTo:   tweet.QuoteTo.String(),
			CreatedAt: timestamppb.New(tweet.CreatedAt),
		})
	}

	return convertedTweets
}
