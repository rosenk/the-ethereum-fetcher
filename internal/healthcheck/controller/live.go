package controller

import (
	"io"
	"net/http"

	"github.com/sumup-oss/go-pkgs/logger"
)

const (
	liveContentTypeHeader    = "Content-Type"
	liveContentTypeTextPlain = "text/plain; charset=utf-8"
	liveResponse             = "🏥 Live"
)

func (c *Controller) Live(response http.ResponseWriter, request *http.Request) {
	_, _ = io.Copy(io.Discard, request.Body)

	httpStatus := http.StatusOK

	response.Header().Set(liveContentTypeHeader, liveContentTypeTextPlain)
	response.WriteHeader(httpStatus)

	_, err := response.Write([]byte(liveResponse))
	if err != nil {
		c.logger.Error(
			logMessageFailedToWriteResponseBody,
			emojiField("🏥❌"),
			logger.ErrorField(err),
		)

		return
	}

	c.logger.Debug(
		logMessageLivenessCheck,
		emojiField("🏥"),
		userAgentField(request.UserAgent()),
		httpStatusField(httpStatus),
	)
}
