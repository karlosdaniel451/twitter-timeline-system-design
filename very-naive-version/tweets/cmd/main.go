package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"tweets/api/grpc_api"
	"tweets/cmd/setup"
	"tweets/config"
)

func main() {
	appConfig := config.NewEmptyAppConfig()

	if err := setup.Setup(appConfig); err != nil {
		slog.Error("error when setting up application", "error", err)
		os.Exit(1)
	}

	fmt.Printf("App config params:\n%s\n", *appConfig)

	if err := grpc_api.StartApp(*appConfig); err != nil {
		slog.Error(
			"failed to start gRPC server at "+strconv.Itoa(appConfig.ListenerPort),
			"error", err,
		)
		os.Exit(1)
	}
}
