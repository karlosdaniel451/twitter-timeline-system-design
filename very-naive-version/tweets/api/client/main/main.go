package main

import (
	"context"
	"fmt"
	"log"
	"tweets/api/protobuf/tweets_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	serverHost string = "localhost"
	serverPort int = 8000
)

func main() {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", serverHost, serverPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("error: failed to connect to %s:%d: %s", serverHost, serverPort, err)
	}
	defer conn.Close()

	ctx := context.Background()

	tweetsService := tweets_service.NewTweetsServiceClient(conn)

	allTweets, err := tweetsService.GetAllTweets(
		ctx,
		&tweets_service.GetAllTweetsRequest{},
	)

	if err != nil {
		log.Fatalf("error when calling gRPC server: %s", err)
	}

	for _, tweet := range allTweets.Tweets {
		fmt.Println(tweet.String())
	}
}
