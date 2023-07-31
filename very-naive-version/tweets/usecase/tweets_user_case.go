package usecase

import (
	"tweets/domain/models"
	"tweets/repository"

	"github.com/google/uuid"
)

type TweetsUseCase interface {
	CreateTweet(tweet *models.Tweet) (*models.Tweet, error)
	GetTweetById(id uuid.UUID) (*models.Tweet, error)
	DeleteTweetById(id uuid.UUID) error
	GetAllTweets() ([]*models.Tweet, error)
	GetTweetOfUser(user_id uuid.UUID) ([]*models.Tweet, error)
}

type TweetsUseCaseImpl struct {
	tweetsRepository repository.TweetRepository
}

func NewTweetsUseCaseImpl(
	tweetsRepository repository.TweetRepository) *TweetsUseCaseImpl {

	return &TweetsUseCaseImpl{tweetsRepository: tweetsRepository}
}

func (useCase *TweetsUseCaseImpl) CreateTweet(
	tweet *models.Tweet) (*models.Tweet, error) {
	
	createdTweet, err := useCase.tweetsRepository.CreateTweet(tweet)
	if err != nil {
		return nil, err
	}
	
	return createdTweet, err
}

func (useCase *TweetsUseCaseImpl) GetTweetById(id uuid.UUID) (*models.Tweet, error) {
	tweet, err := useCase.tweetsRepository.GetTweetById(id)
	if err != nil {
		return nil, err
	}

	return tweet, err
}

func (useCase *TweetsUseCaseImpl) DeleteTweetById(id uuid.UUID) error {
	return useCase.tweetsRepository.DeleteTweetById(id)
}

func (useCase *TweetsUseCaseImpl) GetAllTweets() ([]*models.Tweet, error) {
	return useCase.tweetsRepository.GetAllTweets()
}

func (useCase *TweetsUseCaseImpl) GetTweetOfUser(user_id uuid.UUID) ([]*models.Tweet, error) {
	tweets, err := useCase.tweetsRepository.GetTweetsOfUser(user_id)
	if err != nil {
		return nil, err
	}

	return tweets, nil
}
