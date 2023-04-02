package controller

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sumup-oss/go-pkgs/logger"
)

type Controller struct {
	logger     logger.StructuredLogger
	registerer prometheus.Registerer
	gatherer   prometheus.Gatherer
	path       string
}

func (c *Controller) Mount(router *chi.Mux) {
	router.Handle(
		c.path,
		promhttp.InstrumentMetricHandler(
			c.registerer,
			promhttp.HandlerFor(c.gatherer, promhttp.HandlerOpts{}),
		),
	)
}
