package models

import (
	"time"

	"github.com/google/uuid"
)

type Tweet struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();"`
	Text      string    `gorm:"varchar(280);not null"`
	UserId    uuid.UUID `gorm:"type:uuid;not null"`
	RepliesTo uuid.UUID
	QuoteTo   uuid.UUID
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
