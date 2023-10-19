package setup

import (
	"fmt"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"strconv"
	"tweets/api/grpc_api/controller"
	"tweets/api/grpc_api/protobuf/tweets_service"
	"tweets/api/grpc_api/protobuf/users_service"
	"tweets/config"
	"tweets/db"
	"tweets/errs"
	"tweets/repository"
	"tweets/usecase"
)

var (
	// Repositories
	UsersRepository  repository.UserRepository
	TweetsRepository repository.TweetRepository

	// Usecases
	UsersUseCase  usecase.UsersUseCase
	TweetsUseCase usecase.TweetsUseCase

	// API Controllers
	UsersController  controller.UsersController
	TweetsController controller.TweetsController
)

func Setup(appConfig *config.AppConfig) error {
	assertInterfaces()
	setupLogger()

	if err := setEnvVariables(appConfig); err != nil {
		return fmt.Errorf("error when setting environment variables: %s", err)
	}

	// Try to connect to the database server.
	if err := db.Connect(*appConfig); err != nil {
		return fmt.Errorf("failed to connect to database: %s", err)
	}

	// Setup for Users.
	UsersRepository = repository.NewUserRepositoryDB(db.DB)
	UsersUseCase = usecase.NewUsersUseCaseImpl(UsersRepository)
	UsersController = controller.NewUsersController(UsersUseCase)

	// Setup for Tweets.
	TweetsRepository = repository.NewTweetRepositoryDB(db.DB)
	TweetsUseCase = usecase.NewTweetsUseCaseImpl(TweetsRepository)
	TweetsController = controller.NewTweetsController(TweetsUseCase)

	return nil
}

func setEnvVariables(appConfig *config.AppConfig) error {
	// Try to load .env file for environment variables.
	if err := godotenv.Load(".env"); err != nil {
		slog.Error("error when loading .env file", "error", err)
		os.Exit(1)
	}

	appPortNumber, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		slog.Error("invalid config params: invalid app port number", "error", err)
		os.Exit(1)
	}

	appEnvironmentType, err := config.ParseAppEnvironmentType(
		os.Getenv("APP_ENVIRONMENT_TYPE"),
	)
	if err != nil {
		return fmt.Errorf("invalid config params: invalid app environment")
	}

	dbPortNumber, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return fmt.Errorf("invalid config params: invalid database port number")
	}

	appConfig.ListenerPort = appPortNumber
	appConfig.AppEnvironmentType = appEnvironmentType
	appConfig.DatabaseHost = os.Getenv("DB_HOST")
	appConfig.DatabaseUser = os.Getenv("DB_USER")
	appConfig.DatabasePort = dbPortNumber
	appConfig.DatabaseName = os.Getenv("DB_NAME")
	appConfig.DatabasePassword = os.Getenv("DB_PASSWORD")

	return nil
}

func setupLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func assertInterfaces() {
	// Assertions for Users.
	var _ users_service.UsersServiceServer = &controller.UsersController{}
	var _ repository.UserRepository = &repository.UserRepositoryDB{}

	// Assertions for Tweets.
	var _ tweets_service.TweetsServiceServer = &controller.TweetsController{}
	var _ repository.TweetRepository = &repository.TweetRepositoryDB{}

	// Assertions for custom errors.
	var _ error = &errs.ErrorNotFound{}
}
