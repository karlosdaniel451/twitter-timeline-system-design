package setup

import (
	"log/slog"
	"os"
	"tweets/api/grpc/controller"
	"tweets/api/grpc/controller/protobuf/tweets_service"
	"tweets/db"
	"tweets/repository"
	repositoryerrors "tweets/repository/repository_errors"
	"tweets/usecase"
)

var (
	// Repositories
	TweetsRepository repository.TweetRepository

	// Usecases
	TweetsUseCase usecase.TweetsUseCase

	// API Controllers
	TweetsController controller.TweetsController
)

func Setup() {
	assertInterfaces()

	setupLogger()

	// Try to connect to the database server.
	if err := db.Connect(); err != nil {
		slog.Error("failed to connect to database", "error", err)
	}

	// Setup for Tweets.
	TweetsRepository = repository.NewTweetRepositoryDB(db.DB)
	TweetsUseCase = usecase.NewTweetsUseCaseImpl(TweetsRepository)
	TweetsController = controller.NewTweetsController(TweetsUseCase)
}

func setupLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func assertInterfaces() {
	var _ tweets_service.TweetsServiceServer = &controller.TweetsController{}
	var _ repository.TweetRepository = &repository.TweetRepositoryDB{}
	var _ error = &repositoryerrors.ErrorNotFound{}
}
