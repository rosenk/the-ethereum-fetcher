package wrapper

import (
	"go.uber.org/zap"
)

const (
	logMessageRenderError = "ACQUIRER WRAPPER: RENDER ERROR"
)

func emojiField(emoji string) zap.Field {
	return zap.String("emoji", emoji)
}
