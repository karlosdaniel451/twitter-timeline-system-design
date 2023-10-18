package models

import (
	"time"

	"github.com/google/uuid"
)

type Tweet struct {
	Id        *uuid.UUID     `gorm:"primaryKey;type:uuid;default:gen_random_uuid();"`
	Text      *string        `gorm:"varchar(280);not null"`
	UserId    *uuid.UUID     `gorm:"type:uuid;not null"`
	Replies   *Tweet         `gorm:"foreignKey:RepliesTo"`
	RepliesTo *uuid.UUID     `gorm:"type:uuid"`
	Quote     *Tweet         `gorm:"foreignKey:QuoteTo"`
	QuoteTo   *uuid.UUID     `gorm:"type:uuid"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
