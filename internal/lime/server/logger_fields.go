package server

import (
	"go.uber.org/zap"
)

const (
	logMessageStart          = "LIME SERVER: STARTING"
	logMessageStartError     = "LIME SERVER: START ERROR"
	logMessageShutdownSignal = "LIME SERVER: SHUTDOWN RECEIVED"
	logMessageShutdown       = "LIME SERVER: SHUTDOWN"
	logMessageShutdownError  = "LIME SERVER: SHUTDOWN ERROR"
)

func emojiField(emoji string) zap.Field {
	return zap.String("emoji", emoji)
}

func addressField(address string) zap.Field {
	return zap.String("address", address)
}
