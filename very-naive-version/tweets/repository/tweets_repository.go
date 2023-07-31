package repository

import (
	"fmt"
	"tweets/domain/models"

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
	if result.Error == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("there is no tweet with id %s", id.String())
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &tweet, nil
}

func (repository TweetRepositoryDB) DeleteTweetById(id uuid.UUID) error {
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
	return nil, nil
}
