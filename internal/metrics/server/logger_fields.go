package server

import (
	"go.uber.org/zap"
)

const (
	logMessageStart          = "METRICS SERVER: STARTING"
	logMessageStartError     = "METRICS SERVER: START ERROR"
	logMessageShutdownSignal = "METRICS SERVER: SHUTDOWN RECEIVED"
	logMessageShutdown       = "METRICS SERVER: SHUTDOWN"
	logMessageShutdownError  = "METRICS SERVER: SHUTDOWN ERROR"
)

func emojiField(emoji string) zap.Field {
	return zap.String("emoji", emoji)
}

func addressField(address string) zap.Field {
	return zap.String("address", address)
}
