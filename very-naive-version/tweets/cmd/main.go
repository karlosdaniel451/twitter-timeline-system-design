package main

import (
	"log/slog"
	"os"
	"strconv"
	"tweets/api/grpc_api"
	"tweets/cmd/setup"
	"tweets/config"
)

var (
	portNumber = os.Getenv("APP_PORT")
)

func main() {
	setup.Setup()

	// Config params.
	var (
		portNumber     int
		appEnvironment config.AppEnvironment
	)

	portNumber, err := strconv.Atoi(portNumberUnparsed)
	if err != nil {
		slog.Error("invalid config params: invalid port number", "error", err)
		os.Exit(1)
	}

	appEnvironment, err = config.ParseAppEnvironment(appEnvironmentUnparsed)
	if err != nil {
		slog.Error("invalid config params: invalid app environment", "error", err)
		os.Exit(1)
	}

	appConfig := config.NewAppConfig(portNumber, appEnvironment)

	if err = grpc_api.StartApp(*appConfig); err != nil {
		slog.Error(
			"failed to start gRPC server at "+portNumberUnparsed,
			"error", err,
		)
	}
}
