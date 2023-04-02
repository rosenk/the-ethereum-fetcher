package server

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sumup-oss/go-pkgs/errors"
	"github.com/sumup-oss/go-pkgs/logger"
)

type Server struct {
	logger                  logger.StructuredLogger
	gracefulShutdownTimeout time.Duration
	listenAddress           string

	router     *chi.Mux
	httpServer *http.Server
}

func (s *Server) Run(ctx context.Context) error {
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	defer waitGroup.Wait()

	doneCh := make(chan struct{})
	defer close(doneCh)

	go func() {
		defer waitGroup.Done()

		select {
		case <-ctx.Done():
			s.logger.Info(
				logMessageShutdownSignal,
				emojiField("🏦🛑"),
				addressField(s.listenAddress),
			)

			shutdownContext, cancel := context.WithTimeout(context.Background(), s.gracefulShutdownTimeout)
			defer cancel()

			if err := s.httpServer.Shutdown(shutdownContext); err != nil { //nolint:contextcheck
				s.logger.Warn(
					logMessageShutdownError,
					emojiField("🏦❌"),
					logger.ErrorField(err),
				)
			}
		case <-doneCh:
		}
	}()

	s.logger.Info(
		logMessageStart,
		emojiField("🏦🚀"),
		addressField(s.listenAddress),
	)

	err := s.httpServer.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		s.logger.Info(
			logMessageShutdown,
			emojiField("🏦🛑"),
			addressField(s.listenAddress),
			logger.ErrorField(err),
		)
	} else {
		s.logger.Error(
			logMessageStartError,
			emojiField("🏦❌"),
			logger.ErrorField(err),
		)
	}

	return errors.Wrap(err, "lime stopped")
}
