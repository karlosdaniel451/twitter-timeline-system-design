package db

import (
	"fmt"
	"tweets/config"
	"tweets/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(appConfig config.AppConfig) error {
	var err error

	connectionConfig := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		appConfig.DatabaseHost, appConfig.DatabaseUser, appConfig.DatabasePassword,
		appConfig.DatabaseName, appConfig.DatabasePort,
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
