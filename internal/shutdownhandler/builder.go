package shutdownhandler

import (
	"time"

	"github.com/sumup-oss/go-pkgs/logger"
)

type Builder struct {
	shutdownTimeout time.Duration
}

func NewBuilder() *Builder {
	return &Builder{
		shutdownTimeout: defaultShutdownTimeout,
	}
}

func (b *Builder) WithShutdownTimeout(shutdownTimeout time.Duration) *Builder {
	b.shutdownTimeout = shutdownTimeout

	return b
}

func (b *Builder) Build(log logger.StructuredLogger) *Handler {
	return &Handler{
		logger:          log,
		shutdownTimeout: b.shutdownTimeout,
	}
}
