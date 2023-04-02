package controller

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

func (hc *Server) Run(ctx context.Context) error {
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	defer waitGroup.Wait()

	doneCh := make(chan struct{})
	defer close(doneCh)

	go func() {
		defer waitGroup.Done()

		select {
		case <-ctx.Done():
			hc.logger.Info(
				logMessageShutdownSignal,
				emojiField("🏥🛑"),
				addressField(hc.httpServer.Addr),
			)

			shutdownContext, cancel := context.WithTimeout(context.Background(), hc.gracefulShutdownTimeout)
			defer cancel()

			if err := hc.httpServer.Shutdown(shutdownContext); err != nil { //nolint:contextcheck
				hc.logger.Warn(
					logMessageShutdownError,
					emojiField("🏥❌"),
					logger.ErrorField(err),
				)
			}
		case <-doneCh:
		}
	}()

	hc.logger.Info(
		logMessageStart,
		emojiField("🏥🚀"),
		addressField(hc.httpServer.Addr),
	)

	err := hc.httpServer.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		hc.logger.Info(
			logMessageShutdown,
			emojiField("🏥🛑"),
			addressField(hc.httpServer.Addr),
			logger.ErrorField(err),
		)
	} else {
		hc.logger.Error(
			logMessageStartError,
			emojiField("🏥❌"),
			logger.ErrorField(err),
		)
	}

	return errors.Wrap(err, "health check stopped: %s", err.Error())
}
