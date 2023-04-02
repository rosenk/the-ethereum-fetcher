package controller

import (
	"github.com/sumup-oss/go-pkgs/logger"
)

type Builder struct {
	readyPath string
	livePath  string
}

func NewBuilder() *Builder {
	return &Builder{
		readyPath: defaultReadyPath,
		livePath:  defaultLivePath,
	}
}

func (b *Builder) WithReadyPath(readyPath string) *Builder {
	b.readyPath = readyPath

	return b
}

func (b *Builder) WithLivePath(livePath string) *Builder {
	b.livePath = livePath

	return b
}

func (b *Builder) Build(log logger.StructuredLogger) *Controller {
	return &Controller{
		logger:    log,
		readyPath: b.readyPath,
		livePath:  b.livePath,
	}
}
