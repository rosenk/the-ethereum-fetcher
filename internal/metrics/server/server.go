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

	router     *chi.Mux
	httpServer *http.Server
}

func (s *Server) Run(ctx context.Context) error {
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)

	doneCh := make(chan struct{})

	go func() {
		defer waitGroup.Done()

		select {
		case <-ctx.Done():
			s.logger.Info(
				logMessageShutdownSignal,
				emojiField("📈🛑"),
				addressField(s.httpServer.Addr),
			)

			shutdownContext, cancel := context.WithTimeout(context.Background(), s.gracefulShutdownTimeout)
			defer cancel()

			err := s.httpServer.Shutdown(shutdownContext) //nolint:contextcheck
			if err != nil {
				s.logger.Warn(
					logMessageShutdownError,
					emojiField("📈❌"),
					logger.ErrorField(err),
				)
			}
		case <-doneCh:
		}
	}()

	s.logger.Info(
		logMessageStart,
		emojiField("📈🚀"),
		addressField(s.httpServer.Addr),
	)

	err := s.httpServer.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		s.logger.Info(
			logMessageShutdown,
			emojiField("📈🛑"),
			addressField(s.httpServer.Addr),
			logger.ErrorField(err),
		)
	} else {
		s.logger.Error(
			logMessageStartError,
			emojiField("📈❌"),
			logger.ErrorField(err),
		)
	}

	close(doneCh)
	waitGroup.Wait()

	return errors.Wrap(err, "metrics server stopped")
}
