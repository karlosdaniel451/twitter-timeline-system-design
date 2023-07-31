package controller

import (
	"context"
	"fmt"
	"tweets/api/protobuf/tweets_service"
	"tweets/domain/models"
	"tweets/usecase"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TweetsController struct {
	tweets_service.UnimplementedTweetsServiceServer
	tweetsUseCase usecase.TweetsUseCase
}

func NewTweetsController(tweetsUseCase usecase.TweetsUseCase) *TweetsController {
	return &TweetsController{tweetsUseCase: tweetsUseCase}
}

func (controller *TweetsController) PostTweet(
	ctx context.Context,
	request *tweets_service.PostTweetRequest,
) (*tweets_service.PostTweetResponse, error) {

	userId, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, fmt.Errorf("error when parsing tweet: invalid user_id: %s", err)
	}

	var repliesTo uuid.UUID
	if request.GetRepliesTo() != "" {
		repliesTo, err = uuid.Parse(request.GetUserId())
		if err != nil {
			return nil, fmt.Errorf("error when parsing tweet: invalid replies_to: %s", err)
		}
	}

	var quoteTo uuid.UUID
	if request.GetRepliesTo() != "" {
		quoteTo, err = uuid.Parse(request.GetUserId())
		if err != nil {
			return nil, fmt.Errorf("error when parsing tweet: invalid quote_to: %s", err)
		}
	}

	tweet := models.Tweet{
		Text:      request.GetText(),
		UserId:    userId,
		RepliesTo: repliesTo,
		QuoteTo:   quoteTo,
	}

	createdTweet, err := controller.tweetsUseCase.CreateTweet(&tweet)
	if err != nil {
		return nil, err
	}

	response := tweets_service.PostTweetResponse{}
	response.Tweet = &tweets_service.Tweet{
		Id:        createdTweet.Id.String(),
		Text:      createdTweet.Text,
		UserId:    createdTweet.UserId.String(),
		RepliesTo: createdTweet.RepliesTo.String(),
		QuoteTo:   tweet.QuoteTo.String(),
		CreatedAt: timestamppb.New(createdTweet.CreatedAt),
	}

	return &response, nil
}

func (controller *TweetsController) DeleteTweetById(
	ctx context.Context,
	request *tweets_service.DeleteTweetByIdRequest,
) (*tweets_service.DeleteTweetByIdResponse, error) {

	return nil, nil
}
func (controller *TweetsController) GetTweetById(
	ctx context.Context,
	request *tweets_service.GetTweetByIdRequest,
) (*tweets_service.GetTweetByIdResponse, error) {

	return nil, nil
}

func (controller *TweetsController) GetAllTweets(
	ctx context.Context,
	request *tweets_service.GetAllTweetsRequest,
) (*tweets_service.GetAllTweetsResponse, error) {

	allTweets, err := controller.tweetsUseCase.GetAllTweets()
	if err != nil {
		return nil, err
	}

	response := tweets_service.GetAllTweetsResponse{}

	for _, tweet := range allTweets {
		response.Tweets = append(response.Tweets, &tweets_service.Tweet{
			Id:        tweet.Id.String(),
			Text:      tweet.Text,
			UserId:    tweet.UserId.String(),
			RepliesTo: tweet.RepliesTo.String(),
			QuoteTo:   tweet.QuoteTo.String(),
			CreatedAt: timestamppb.New(tweet.CreatedAt),
		})
	}

	return &response, nil
}

func (controller *TweetsController) GetTweetsOfUser(
	ctx context.Context,
	request *tweets_service.GetTweetByIdRequest,
) (*tweets_service.GetTweetByIdResponse, error) {

	return nil, nil
}
