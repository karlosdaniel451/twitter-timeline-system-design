package db

import (
	"fmt"
	"tweets/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	port     = 5432
	database = "twitter_naive"
	password = "123456789"
)

var DB *gorm.DB

func Connect() error {
	var err error

	connectionConfig := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		host, user, password, database, port,
	)

	DB, err = gorm.Open(postgres.Open(connectionConfig), &gorm.Config{})

	if err != nil {
		return err
	}

	err = DB.AutoMigrate(&models.Tweet{})
	if err != nil {
		return err
	}

	return nil
}
