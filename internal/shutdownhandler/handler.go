package shutdownhandler

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sumup-oss/go-pkgs/errors"
	"github.com/sumup-oss/go-pkgs/logger"
)

var ErrShutdown = errors.New("shutdown")

type Handler struct {
	shutdownTimeout time.Duration
	logger          logger.StructuredLogger
}

func (s *Handler) Run(ctx context.Context) error {
	osSignalChan := make(chan os.Signal, 1)

	signal.Notify(osSignalChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-osSignalChan:
		signal.Stop(osSignalChan)

		s.logger.Info(
			logMessageShutdownSignal,
			emojiField("🛑"),
			signalField(sig),
			deadlineHumanField(s.shutdownTimeout),
		)

		go func() {
			time.Sleep(s.shutdownTimeout)

			s.logger.Error(
				logMessageShutdownForce,
				emojiField("🛑❌"),
				deadlineHumanField(s.shutdownTimeout),
			)

			os.Exit(1)
		}()
	case <-ctx.Done():
	}

	return ErrShutdown
}
