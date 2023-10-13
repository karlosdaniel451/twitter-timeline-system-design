package middleware

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"log/slog"
	logging2 "tweets/config/logging"
)

func GetUnaryMiddlewares() grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		logging.UnaryServerInterceptor(
			logging2.InterceptorLogger(slog.Default()), logging2.GetLogOptions()...,
		),
	)
}
