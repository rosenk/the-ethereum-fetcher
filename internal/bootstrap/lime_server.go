package bootstrap

import (
	"github.com/go-playground/validator/v10"
	"github.com/ju-popov/the-ethereum-fetcher/internal/config"
	"github.com/ju-popov/the-ethereum-fetcher/internal/db/maindb"
	"github.com/ju-popov/the-ethereum-fetcher/internal/fetcher"
	ethcontroller "github.com/ju-popov/the-ethereum-fetcher/internal/lime/eth/controller"
	limeserver "github.com/ju-popov/the-ethereum-fetcher/internal/lime/server"
	"github.com/sumup-oss/go-pkgs/logger"
)

func LimeServer(
	log logger.StructuredLogger,
	validate *validator.Validate,
	mainDBClient *maindb.Client,
	fetcherClient *fetcher.Client,
	conf *config.Lime,
) *limeserver.Server {
	server, router := limeserver.NewBuilder().
		WithListenAddress(*conf.Server.ListenAddress).
		WithReadHeaderTimeout(*conf.Server.ReadHeaderTimeout).
		WithReadTimeout(*conf.Server.ReadTimeout).
		WithWriteTimeout(*conf.Server.WriteTimeout).
		WithGracefulShutdownTimeout(*conf.Server.GracefulShutdownTimeout).
		Build(log)

	controller := ethcontroller.New(log, validate, mainDBClient, fetcherClient)

	controller.Mount(router)

	return server
}
