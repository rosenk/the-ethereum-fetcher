package controller

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sumup-oss/go-pkgs/logger"
)

type Builder struct {
	path       string
	registerer prometheus.Registerer
	gatherer   prometheus.Gatherer
}

func NewBuilder() *Builder {
	return &Builder{
		path:       defaultPath,
		registerer: prometheus.DefaultRegisterer,
		gatherer:   prometheus.DefaultGatherer,
	}
}

func (b *Builder) WithPath(path string) *Builder {
	b.path = path

	return b
}

func (b *Builder) WithRegisterer(registerer prometheus.Registerer) *Builder {
	b.registerer = registerer

	return b
}

func (b *Builder) WithGatherer(gatherer prometheus.Gatherer) *Builder {
	b.gatherer = gatherer

	return b
}

func (b *Builder) Build(log logger.StructuredLogger) *Controller {
	return &Controller{
		logger:     log,
		registerer: b.registerer,
		gatherer:   b.gatherer,
		path:       b.path,
	}
}
