package controller

import (
	"go.uber.org/zap"
)

const (
	logMessageStart          = "HEALTH CHECK SERVER: STARTING"
	logMessageStartError     = "HEALTH CHECK SERVER: START ERROR"
	logMessageShutdownSignal = "HEALTH CHECK SERVER: SHUTDOWN RECEIVED"
	logMessageShutdown       = "HEALTH CHECK SERVER: SHUTDOWN"
	logMessageShutdownError  = "HEALTH CHECK SERVER: SHUTDOWN ERROR"
)

func emojiField(emoji string) zap.Field {
	return zap.String("emoji", emoji)
}

func addressField(address string) zap.Field {
	return zap.String("address", address)
}
