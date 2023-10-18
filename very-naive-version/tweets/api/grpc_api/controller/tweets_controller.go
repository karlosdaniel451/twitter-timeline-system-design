package controller

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"tweets/api/grpc_api/converter/tweet_converter"
	tweets_service2 "tweets/api/grpc_api/protobuf/tweets_service"
	"tweets/domain/models"
	"tweets/errs"
	"tweets/usecase"
	"tweets/utils"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TweetsController struct {
	tweets_service2.UnimplementedTweetsServiceServer
	tweetsUseCase usecase.TweetsUseCase
}

func NewTweetsController(tweetsUseCase usecase.TweetsUseCase) TweetsController {
	return TweetsController{tweetsUseCase: tweetsUseCase}
}

func (controller *TweetsController) PostTweet(
	ctx context.Context,
	request *tweets_service2.PostTweetRequest,
) (*tweets_service2.PostTweetResponse, error) {

	userId, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("invalid user_id uuid: %s", err),
		)
	}

	var repliesTo uuid.UUID
	if request.GetRepliesTo() != "" {
		repliesTo, err = uuid.Parse(request.GetUserId())
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument, fmt.Sprintf("invalid replies_to uuid: %s", err),
			)
		}
	}

	var quoteTo uuid.UUID
	if request.GetRepliesTo() != "" {
		quoteTo, err = uuid.Parse(request.GetUserId())
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument, fmt.Sprintf("invalid quote_to uuid: %s", err),
			)
		}
	}

	tweet := models.Tweet{
		Text:      utils.ValueToPointer[string](request.GetText()),
		UserId:    &userId,
		RepliesTo: &repliesTo,
		QuoteTo:   &quoteTo,
	}

	createdTweet, err := controller.tweetsUseCase.CreateTweet(&tweet)
	if err != nil {
		slog.Error(
			"internal error when creating tweet",
			"error", err,
			"request", request.String(),
		)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	response := tweets_service2.PostTweetResponse{
		Tweet: tweet_converter.FromDatabaseModelToProtobufModel(createdTweet),
	}

	return &response, nil
}

func (controller *TweetsController) DeleteTweetById(
	ctx context.Context,
	request *tweets_service2.DeleteTweetByIdRequest,
) (*tweets_service2.DeleteTweetByIdResponse, error) {

	response := tweets_service2.DeleteTweetByIdResponse{}
	tweetId, err := uuid.Parse(request.TweetId)
	if err != nil {
		return &response, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("invalid tweet_id uuid: %s", err),
		)
	}

	err = controller.tweetsUseCase.DeleteTweetById(tweetId)
	if err != nil {
		var errNotFound *errs.ErrorNotFound
		if errors.As(err, &errNotFound) {
			return nil, status.Errorf(
				codes.NotFound, err.Error(),
			)
		}

		slog.Error(
			"internal error when creating tweet",
			"error", err,
			"request", request.String(),
		)
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	return &response, nil
}

func (controller *TweetsController) GetTweetById(
	ctx context.Context,
	request *tweets_service2.GetTweetByIdRequest,
) (*tweets_service2.GetTweetByIdResponse, error) {

	tweetId, err := uuid.Parse(request.TweetId)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("invalid tweet_id uuid: %s", err),
		)
	}

	tweet, err := controller.tweetsUseCase.GetTweetById(tweetId)
	if err != nil {
		var errNotFound *errs.ErrorNotFound
		if errors.As(err, &errNotFound) {
			return nil, status.Errorf(
				codes.NotFound, err.Error(),
			)
		}

		slog.Error(
			"internal error when getting tweet by id",
			"error", err, "request",
			request.String(),
		)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	response := tweets_service2.GetTweetByIdResponse{
		Tweet: tweet_converter.FromDatabaseModelToProtobufModel(tweet),
	}

	return &response, nil
}

func (controller *TweetsController) GetAllTweets(
	ctx context.Context,
	request *tweets_service2.GetAllTweetsRequest,
) (*tweets_service2.GetAllTweetsResponse, error) {

	allTweets, err := controller.tweetsUseCase.GetAllTweets()
	if err != nil {
		slog.Error(
			"internal error when getting all tweets",
			"error", err, "request",
			request.String(),
		)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	response := tweets_service2.GetAllTweetsResponse{
		Tweets: tweet_converter.FromDatabaseModelsToProtobufModels(allTweets),
	}

	return &response, nil
}

func (controller *TweetsController) GetTweetsOfUser(
	ctx context.Context,
	request *tweets_service2.GetTweetsOfUserRequest,
) (*tweets_service2.GetTweetsOfUserResponse, error) {

	userId, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("invalid user_id uuid: %s", err),
		)
	}

	tweets, err := controller.tweetsUseCase.GetTweetOfUser(userId)
	if err != nil {
		var errNotFound *errs.ErrorNotFound
		if errors.As(err, &errNotFound) {
			return nil, status.Errorf(
				codes.NotFound, err.Error(),
			)
		}
		slog.Error(
			"internal error when getting tweets of user",
			"error", err,
			"request", request.String(),
		)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	response := tweets_service2.GetTweetsOfUserResponse{
		Tweets: tweet_converter.FromDatabaseModelsToProtobufModels(tweets),
	}

	return &response, nil
}
