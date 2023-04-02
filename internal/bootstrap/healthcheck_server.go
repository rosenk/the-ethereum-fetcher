package bootstrap

import (
	"github.com/ju-popov/the-ethereum-fetcher/internal/config"
	healthcheckcontroller "github.com/ju-popov/the-ethereum-fetcher/internal/healthcheck/controller"
	healthcheckserver "github.com/ju-popov/the-ethereum-fetcher/internal/healthcheck/server"
	"github.com/sumup-oss/go-pkgs/logger"
)

func HealthcheckServer(
	log logger.StructuredLogger,
	conf *config.HealthCheck,
) *healthcheckserver.Server {
	server, router := healthcheckserver.NewBuilder().
		WithListenAddress(*conf.Server.ListenAddress).
		WithReadHeaderTimeout(*conf.Server.ReadHeaderTimeout).
		WithReadTimeout(*conf.Server.ReadTimeout).
		WithWriteTimeout(*conf.Server.WriteTimeout).
		WithGracefulShutdownTimeout(*conf.Server.GracefulShutdownTimeout).
		Build(log)

	controller := healthcheckcontroller.NewBuilder().
		WithReadyPath(*conf.Controller.ReadyPath).
		WithLivePath(*conf.Controller.LivePath).
		Build(log)

	controller.Mount(router)

	return server
}
