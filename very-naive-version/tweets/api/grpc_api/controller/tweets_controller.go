package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"tweets/api/grpc_api/converter/tweet_converter"
	tweetsService "tweets/api/grpc_api/protobuf/tweets_service"
	"tweets/errs"
	"tweets/usecase"
)

type TweetsController struct {
	tweetsService.UnimplementedTweetsServiceServer
	tweetsUseCase usecase.TweetsUseCase
}

func NewTweetsController(tweetsUseCase usecase.TweetsUseCase) TweetsController {
	return TweetsController{tweetsUseCase: tweetsUseCase}
}

func (controller *TweetsController) PostTweet(
	ctx context.Context,
	request *tweetsService.PostTweetRequest,
) (*tweetsService.PostTweetResponse, error) {

	if request.GetUserId() == "" {
		return nil, status.Errorf(
			codes.InvalidArgument, "user_id field is required",
		)
	}

	if request.Text == "" {
		return nil, status.Errorf(
			codes.InvalidArgument, "text field is required",
		)
	}

	_, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("invalid user_id uuid: %s", err),
		)
	}

	if request.GetRepliesTo() != "" {
		_, err := uuid.Parse(request.GetUserId())
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument, fmt.Sprintf("invalid replies_to uuid: %s", err),
			)
		}
	}

	if request.GetQuoteTo() != "" {
		_, err := uuid.Parse(request.GetUserId())
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument, fmt.Sprintf("invalid quote_to uuid: %s", err),
			)
		}
	}

	tweet := tweet_converter.FromPostTweetRequestToDatabaseModel(request)

	createdTweet, err := controller.tweetsUseCase.CreateTweet(tweet)
	if err != nil {
		slog.Error(
			"internal error when creating tweet",
			"error", err,
			"request", request.String(),
		)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	response := tweetsService.PostTweetResponse{
		Tweet: tweet_converter.FromDatabaseModelToProtobufModel(createdTweet),
	}

	return &response, nil
}

func (controller *TweetsController) DeleteTweetById(
	ctx context.Context,
	request *tweetsService.DeleteTweetByIdRequest,
) (*tweetsService.DeleteTweetByIdResponse, error) {

	response := tweetsService.DeleteTweetByIdResponse{}
	tweetId, err := uuid.Parse(request.TweetId)
	if err != nil {
		return &response, status.Errorf(
			codes.InvalidArgument, fmt.Sprintf("invalid tweet_id uuid: %s", err),
		)
	}

	err = controller.tweetsUseCase.DeleteTweetById(tweetId)
	if err != nil {
		if errors.As(err, &errs.ErrorNotFound{}) {
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
	request *tweetsService.GetTweetByIdRequest,
) (*tweetsService.GetTweetByIdResponse, error) {

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

	response := tweetsService.GetTweetByIdResponse{
		Tweet: tweet_converter.FromDatabaseModelToProtobufModel(tweet),
	}

	return &response, nil
}

func (controller *TweetsController) GetAllTweets(
	ctx context.Context,
	request *tweetsService.GetAllTweetsRequest,
) (*tweetsService.GetAllTweetsResponse, error) {

	allTweets, err := controller.tweetsUseCase.GetAllTweets()
	if err != nil {
		slog.Error(
			"internal error when getting all tweets",
			"error", err, "request",
			request.String(),
		)
		return nil, status.Errorf(codes.Internal, "internal error")
	}

	response := tweetsService.GetAllTweetsResponse{
		Tweets: tweet_converter.FromDatabaseModelsToProtobufModels(allTweets),
	}

	return &response, nil
}

func (controller *TweetsController) GetTweetsOfUser(
	ctx context.Context,
	request *tweetsService.GetTweetsOfUserRequest,
) (*tweetsService.GetTweetsOfUserResponse, error) {

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

	response := tweetsService.GetTweetsOfUserResponse{
		Tweets: tweet_converter.FromDatabaseModelsToProtobufModels(tweets),
	}

	return &response, nil
}
