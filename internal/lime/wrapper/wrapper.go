package wrapper

import (
	"net/http"

	"github.com/sumup-oss/go-pkgs/logger"
)

type Wrapper struct {
	logger logger.StructuredLogger
}

func New(log logger.StructuredLogger) *Wrapper {
	return &Wrapper{
		logger: log,
	}
}

type ControllerFunc func(httpRequest *http.Request) (*HTTPResponse, error)

func (w *Wrapper) Wrap(controllerFunc ControllerFunc) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		responseData, err := controllerFunc(httpRequest)
		if err != nil {
			w.logger.Error(
				logMessageControllerError,
				emojiField("🏦❌"),
				logger.ErrorField(err),
			)

			if err := w.renderError(responseWriter, err); err != nil {
				w.logger.Error(
					logMessageRenderError,
					emojiField("🏦🛑"),
					logger.ErrorField(err),
				)
			}

			return
		}

		if err := w.renderResponse(responseWriter, responseData); err != nil {
			w.logger.Error(
				logMessageRenderError,
				emojiField("🏦🛑"),
				logger.ErrorField(err),
			)
		}
	}
}
