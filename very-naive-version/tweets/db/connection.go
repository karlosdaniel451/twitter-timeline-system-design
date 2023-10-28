package db

import (
	"fmt"
	"gorm.io/gorm/logger"
	"tweets/config"
	"tweets/domain/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(appConfig config.AppConfig) error {
	var err error

	connectionConfig := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		appConfig.DatabaseHost, appConfig.DatabaseUser, appConfig.DatabasePassword,
		appConfig.DatabaseName, appConfig.DatabasePort,
	)

	gormLogger := logger.Default

	switch appConfig.AppEnvironmentType {
	case config.Development:
		gormLogger = logger.Default.LogMode(logger.Warn)
	case config.Debugging:
		gormLogger = logger.Default.LogMode(logger.Info)
	case config.Testing:
		gormLogger = logger.Default.LogMode(logger.Warn)
	case config.Production:
		gormLogger = logger.Default.LogMode(logger.Error)
	}

	DB, err = gorm.Open(postgres.Open(connectionConfig), &gorm.Config{
		Logger: gormLogger,
	})

	if err != nil {
		return err
	}

	err = DB.AutoMigrate(&models.User{}, &models.Tweet{}, &models.Follow{})
	if err != nil {
		return err
	}

	return nil
}
