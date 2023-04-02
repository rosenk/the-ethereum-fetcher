package wrapper

import (
	"go.uber.org/zap"
)

const (
	logMessageControllerError = "LIME WRAPPER: CONTROLLER ERROR"
	logMessageRenderError     = "LIME WRAPPER: RENDER ERROR"
)

func emojiField(emoji string) zap.Field {
	return zap.String("emoji", emoji)
}
