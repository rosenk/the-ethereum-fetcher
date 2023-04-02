package controller

import (
	"go.uber.org/zap"
)

const (
	logMessageFailedToWriteResponseBody = "HEALTH CHECK CONTROLLER: FAILED TO WRITE RESPONSE BODY"
	logMessageReadinessCheck            = "HEALTH CHECK CONTROLLER: READINESS CHECK"
	logMessageLivenessCheck             = "HEALTH CHECK CONTROLLER: LIVENESS CHECK"
)

func emojiField(emoji string) zap.Field {
	return zap.String("emoji", emoji)
}

func httpStatusField(status int) zap.Field {
	return zap.Int("http_status", status)
}

func userAgentField(userAgent string) zap.Field {
	return zap.String("user_agent", userAgent)
}
