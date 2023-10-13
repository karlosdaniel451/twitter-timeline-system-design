package main

import (
	"log/slog"
	"net"
	"os"
	"tweets/api/grpc/controller/protobuf/tweets_service"
	"tweets/cmd/setup"
	"tweets/config"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	portNumber = os.Getenv("APP_PORT")
)

func main() {
	setup.Setup()

	// Start TCP connection listener.
	listener, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		slog.Error("failed to listen to "+portNumber, "error", err)
	}

	// Create the gRPC server. But does not accept any RPCs and it has no gRPC service yet.
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		logging.UnaryServerInterceptor(
			config.InterceptorLogger(slog.Default()), config.GetLogOptions()...,
		),
	))

	// Register the TwitterService implementation to the gRPC server.
	tweets_service.RegisterTweetsServiceServer(grpcServer, &setup.TweetsController)

	// Enable server reflection to be used by tools like gRPCurl.
	// Source: https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md
	reflection.Register(grpcServer)

	// Start gRCP server on the TCP listener and block execution in the call to Serve().
	slog.Error("starting gRPC server at " + listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		slog.Error(
			"failed to start gRPC server at "+listener.Addr().String(),
			"error", err,
		)
	}
}
