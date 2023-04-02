package controller

import (
	"io"
	"net/http"

	"github.com/sumup-oss/go-pkgs/logger"
)

const (
	readyContentTypeHeader    = "Content-Type"
	readyContentTypeTextPlain = "text/plain; charset=utf-8"
	readyResponse             = "🏥 Ready"
)

func (c *Controller) Ready(response http.ResponseWriter, request *http.Request) {
	_, _ = io.Copy(io.Discard, request.Body)

	httpStatus := http.StatusOK

	response.Header().Set(readyContentTypeHeader, readyContentTypeTextPlain)
	response.WriteHeader(httpStatus)

	_, err := response.Write([]byte(readyResponse))
	if err != nil {
		c.logger.Error(
			logMessageFailedToWriteResponseBody,
			emojiField("🏥❌"),
			logger.ErrorField(err),
		)

		return
	}

	c.logger.Debug(
		logMessageReadinessCheck,
		emojiField("🏥"),
		userAgentField(request.UserAgent()),
		httpStatusField(httpStatus),
	)
}
