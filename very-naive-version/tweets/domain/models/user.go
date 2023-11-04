package models

import (
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

type User struct {
	gorm.Model
	Id                *uuid.UUID     `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid();"`
	Name              *string        `json:"name" gorm:"varchar(100);not null"`
	Username          *string        `json:"user_name" gorm:"varchar(15);uniqueIndex;not null"`
	Email             *string        `json:"email" gorm:"varchar(320);uniqueIndex;not null"`
	Description       *string        `json:"description" gorm:"varchar(200)"`
	Location          *string        `json:"location" gorm:"varchar(100)"`
	PinnedTweet       *Tweet         `json:"pinned_tweet"`
	PinnedTweetId     *uuid.UUID     `json:"pinned_tweet_id" gorm:"type:uuid"`
	ProfileImageUrl   *string        `json:"profile_image_url"`
	Protected         *bool          `json:"protected" gorm:"not null;default:false"`
	PublicMetrics     *PublicMetrics `json:"public_metrics" gorm:"embedded"`
	Url               *string        `json:"url"`
	Verified          *bool          `json:"verified" gorm:"not null;default:false"`
	MostRecentTweet   *Tweet         `json:"most_recent_tweet"`
	MostRecentTweetId *uuid.UUID     `json:"most_recent_tweet_id" gorm:"type:uuid"`
	Followers         []*User        `json:"followers" gorm:"many2many:follows"`
	FollowerIds       []uuid.UUID    `json:"follower_ids" gorm:"-"` // Field will be ignored by GORM.
	Followees         []*User        `json:"followees" gorm:"many2many:follows"`
	FolloweeIds       []uuid.UUID    `json:"followee_ids" gorm:"-"` // Field will be ignored by GORM.
	Tweets            []*Tweet       `json:"tweets"`
	CreatedAt         time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type PublicMetrics struct {
	FollowersCount uint64 `json:"followers_count" gorm:"not null;default:0"`
	FollowingCount uint64 `json:"following_count" gorm:"not null;default:0"`
	TweetCount     uint64 `json:"tweet_count" gorm:"not null;default:0"`
	ListedCount    uint64 `json:"listed_count" gorm:"not null;default:0"`
}

type Follow struct {
	FollowerId uuid.UUID      `json:"follower_id" gorm:"primaryKey;type:uuid"`
	FolloweeId uuid.UUID      `json:"followee_id" gorm:"primaryKey;type:uuid"`
	Follower   User           `json:"follower" gorm:"foreignKey:FollowerId;references:Id"`
	Followee   User           `json:"followee" gorm:"foreignKey:FolloweeId;references:Id"`
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type UserUniqueFieldsChecking struct {
	Id       bool
	Username bool
	Email    bool
}
