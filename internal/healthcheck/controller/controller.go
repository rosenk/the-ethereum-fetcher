package controller

import (
	"github.com/go-chi/chi/v5"
	"github.com/sumup-oss/go-pkgs/logger"
)

type Controller struct {
	logger    logger.StructuredLogger
	readyPath string
	livePath  string
}

func (c *Controller) Mount(router *chi.Mux) {
	router.Get(c.readyPath, c.Ready)
	router.Get(c.livePath, c.Live)
}
