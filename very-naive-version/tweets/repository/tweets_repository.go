package repository

import (
	"fmt"
	"tweets/domain/models"
	repositoryerrors "tweets/repository/repository_errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TweetRepository interface {
	CreateTweet(tweet *models.Tweet) (*models.Tweet, error)
	GetTweetById(id uuid.UUID) (*models.Tweet, error)
	DeleteTweetById(id uuid.UUID) error
	GetAllTweets() ([]*models.Tweet, error)
	GetTweetsOfUser(user_id uuid.UUID) ([]*models.Tweet, error)
}

type TweetRepositoryDB struct {
	db *gorm.DB
}

func NewTweetRepositoryDB(db *gorm.DB) *TweetRepositoryDB {
	return &TweetRepositoryDB{db: db}
}

func (repository TweetRepositoryDB) CreateTweet(tweet *models.Tweet) (*models.Tweet, error) {
	// query := `
	// INSERT INTO tweeets (id, text, user_id, replies_to, quote_to)
	// VALUES (?, ?, ?, ?, ?)
	// `
	// result, err := repository.db.Exec(
	// 	query, tweet.Id, tweet.Text, tweet.UserId, tweet.RepliesTo, tweet.QuoteTo,
	// )
	// if err != nil {
	// 	return nil, err
	// }

	// rowsAffected, err := result.RowsAffected()
	// if err != nil {
	// 	return nil, err
	// }

	// if rowsAffected != 1 {
	// 	return nil, fmt.Errorf("it was not possible to insert tweet: %s", err)
	// }

	// return

	result := repository.db.Create(tweet)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("it was not possible to insert tweet: %s", result.Error)
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return tweet, nil
}

func (repository TweetRepositoryDB) GetTweetById(id uuid.UUID) (*models.Tweet, error) {
	// query := `
	// SELECT *
	// FROM tweets
	// WHERE tweets.id == ?
	// `
	// result := repository.db.QueryRow(query, id)
	// if result.Err() != nil {
	// 	return nil, result.Err()
	// }

	// var tweet models.Tweet
	// err := result.Scan(&tweet)
	// if err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		return nil, fmt.Errorf("no tweet found with id %s", id.String())
	// 	}
	// 	return nil, err
	// }

	// return &tweet, nil

	var tweet models.Tweet

	result := repository.db.First(&tweet, "id = ?", id)
	if result.Error != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, &repositoryerrors.ErrorNotFound{
				Msg: fmt.Sprintf("there is no tweet with id %s", id.String()),
			}
		}
		return nil, result.Error
	}

	return &tweet, nil
}

func (repository TweetRepositoryDB) DeleteTweetById(id uuid.UUID) error {
	var tweet models.Tweet

	result := repository.db.First(&tweet, id)
	if result.Error != nil {
		if result.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return &repositoryerrors.ErrorNotFound{
				Msg: fmt.Sprintf("there is no tweet with id %s", id.String()),
			}
		}
		return result.Error
	}
	result = result.Delete(&tweet)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository TweetRepositoryDB) GetAllTweets() ([]*models.Tweet, error) {
	allTweets := make([]*models.Tweet, 0)

	result := repository.db.Find(&allTweets)
	if result.Error != nil {
		return nil, result.Error
	}

	return allTweets, nil
}

func (repository TweetRepositoryDB) GetTweetsOfUser(user_id uuid.UUID) ([]*models.Tweet, error) {
	userTweets := make([]*models.Tweet, 0)

	result := repository.db.Find(&userTweets, "user_id = ?", user_id)
	if result.Error != nil {
		return nil, result.Error
	}

	return userTweets, nil
}
