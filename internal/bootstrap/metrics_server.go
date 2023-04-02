package bootstrap

import (
	"github.com/ju-popov/the-ethereum-fetcher/internal/config"
	metricscontroller "github.com/ju-popov/the-ethereum-fetcher/internal/metrics/controller"
	metricsserver "github.com/ju-popov/the-ethereum-fetcher/internal/metrics/server"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sumup-oss/go-pkgs/logger"
)

func MetricsServer(
	log logger.StructuredLogger,
	conf *config.Metrics,
	metricsRegisterer prometheus.Registerer,
	metricsGatherer prometheus.Gatherer,
) *metricsserver.Server {
	server, router := metricsserver.NewBuilder().
		WithListenAddress(*conf.Server.ListenAddress).
		WithReadHeaderTimeout(*conf.Server.ReadHeaderTimeout).
		WithReadTimeout(*conf.Server.ReadTimeout).
		WithWriteTimeout(*conf.Server.WriteTimeout).
		WithGracefulShutdownTimeout(*conf.Server.GracefulShutdownTimeout).
		Build(log)

	controller := metricscontroller.NewBuilder().
		WithPath(*conf.Controller.Path).
		WithRegisterer(metricsRegisterer).
		WithGatherer(metricsGatherer).
		Build(log)

	controller.Mount(router)

	return server
}
