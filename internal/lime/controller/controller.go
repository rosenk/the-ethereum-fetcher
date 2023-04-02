package controller

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ju-popov/the-ethereum-fetcher/internal/db/maindb"
	"github.com/ju-popov/the-ethereum-fetcher/internal/fetcher"
	"github.com/ju-popov/the-ethereum-fetcher/internal/lime/wrapper"
	"github.com/sumup-oss/go-pkgs/logger"
)

type Controller struct {
	logger        logger.StructuredLogger
	validate      *validator.Validate
	mainDBClient  *maindb.Client
	fetcherClient *fetcher.Client
}

func New(
	log logger.StructuredLogger,
	validate *validator.Validate,
	mainDBClient *maindb.Client,
	fetcherClient *fetcher.Client,
) *Controller {
	controller := &Controller{
		logger:        log,
		validate:      validate,
		mainDBClient:  mainDBClient,
		fetcherClient: fetcherClient,
	}

	return controller
}

func (c *Controller) Mount(router *chi.Mux) {
	router.Route("/lime", func(router chi.Router) {
		router.Get("/eth/{rlphex}", wrapper.New(c.logger).Wrap(c.GetETH))
		router.Get("/all", wrapper.New(c.logger).Wrap(c.GetAll))
	})
}
