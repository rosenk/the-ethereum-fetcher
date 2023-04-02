package bootstrap

import (
	"github.com/ju-popov/the-ethereum-fetcher/internal/config"
	"github.com/ju-popov/the-ethereum-fetcher/internal/shutdownhandler"
	"github.com/sumup-oss/go-pkgs/logger"
)

func ShutdownHandler(
	log logger.StructuredLogger,
	conf *config.Shutdown,
) *shutdownhandler.Handler {
	return shutdownhandler.NewBuilder().
		WithShutdownTimeout(*conf.ForcefulShutdownTimeout).
		Build(log)
}
