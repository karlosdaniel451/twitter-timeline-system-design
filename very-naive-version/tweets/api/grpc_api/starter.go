package grpc_api

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
	"strconv"
	"tweets/api/grpc_api/middleware"
	"tweets/api/grpc_api/protobuf/tweets_service"
	"tweets/cmd/setup"
	"tweets/config"
)

func StartApp(config config.AppConfig) error {
	// Create the gRPC server with no service and thus incapable of accept RPCs.
	grpcServer := grpc.NewServer(middleware.GetUnaryMiddlewares())

	// Register the TwitterService implementation to the gRPC server.
	tweets_service.RegisterTweetsServiceServer(grpcServer, &setup.TweetsController)

	// Enable server reflection to be used by tools like gRPCurl.
	// Source: https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md
	reflection.Register(grpcServer)

	// Start TCP connection listener.
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(config.ListenerPort))
	if err != nil {
		return fmt.Errorf(
			"error when starting TCP connection listener at %s: %s",
			listener.Addr().String(),
			err,
		)
	}

	// Start gRCP server on the TCP listener and block execution in the call to Serve().
	slog.Error("starting gRPC server at " + listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		return fmt.Errorf(
			"error when starting gRPC server at %s: %s",
			listener.Addr().String(),
			err,
		)
	}

	return nil
}
