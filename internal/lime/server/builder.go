package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sumup-oss/go-pkgs/logger"
)

type Builder struct {
	listenAddress           string
	readHeaderTimeout       time.Duration
	readTimeout             time.Duration
	writeTimeout            time.Duration
	gracefulShutdownTimeout time.Duration
}

func NewBuilder() *Builder {
	return &Builder{
		listenAddress:           defaultListenAddress,
		readHeaderTimeout:       defaultReadHeaderTimeout,
		readTimeout:             defaultReadTimeout,
		writeTimeout:            defaultWriteTimeout,
		gracefulShutdownTimeout: defaultGracefulShutdownTimeout,
	}
}

func (b *Builder) WithListenAddress(listenAddress string) *Builder {
	b.listenAddress = listenAddress

	return b
}

func (b *Builder) WithReadHeaderTimeout(readHeaderTimeout time.Duration) *Builder {
	b.readHeaderTimeout = readHeaderTimeout

	return b
}

func (b *Builder) WithReadTimeout(readTimeout time.Duration) *Builder {
	b.readTimeout = readTimeout

	return b
}

func (b *Builder) WithWriteTimeout(writeTimeout time.Duration) *Builder {
	b.writeTimeout = writeTimeout

	return b
}

func (b *Builder) WithGracefulShutdownTimeout(gracefulShutdownTimeout time.Duration) *Builder {
	b.gracefulShutdownTimeout = gracefulShutdownTimeout

	return b
}

func (b *Builder) Build(log logger.StructuredLogger) (*Server, *chi.Mux) {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)

	return &Server{
		logger:                  log,
		gracefulShutdownTimeout: b.gracefulShutdownTimeout,
		listenAddress:           b.listenAddress,
		router:                  router,
		httpServer: &http.Server{
			Handler:           router,
			Addr:              b.listenAddress,
			ReadHeaderTimeout: b.readHeaderTimeout,
			ReadTimeout:       b.readTimeout,
			WriteTimeout:      b.writeTimeout,
		},
	}, router
}
