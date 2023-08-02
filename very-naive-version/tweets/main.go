package main

import (
	"fmt"
	"log"
	"net"
	"tweets/api/controller"
	"tweets/api/protobuf/tweets_service"
	"tweets/config"
	"tweets/db"
	"tweets/repository"
	repositoryerrors "tweets/repository/repository_errors"
	"tweets/usecase"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	portNumber int = 8000
)

func main() {
	var _ tweets_service.TweetsServiceServer = &controller.TweetsController{}
	var _ repository.TweetRepository = &repository.TweetRepositoryDB{}
	var _ error = &repositoryerrors.ErrorNotFound{}

	if err := db.Connect(); err != nil {
		log.Fatalf("error: failed to connect to database: %s", err)
	}

	dbConn, _ := db.DB.DB()
	defer dbConn.Close()

	tweetsRepository := repository.NewTweetRepositoryDB(db.DB)
	tweetsUserCase := usecase.NewTweetsUseCaseImpl(tweetsRepository)
	tweetsServer := controller.NewTweetsController(tweetsUserCase)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", portNumber))
	if err != nil {
		log.Fatalf("error: failed to listen to %d: %s", portNumber, err)
	}

	defer listener.Close()

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		logging.UnaryServerInterceptor(config.InterceptorLogger(config.GetLogConfig())),
	))

	tweets_service.RegisterTweetsServiceServer(grpcServer, tweetsServer)
	reflection.Register(grpcServer)

	log.Printf("starting gRPC server at %s\n", listener.Addr())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("error: failed to start gRPC server at %d: %s", portNumber, err)
	}

}
