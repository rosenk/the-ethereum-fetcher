package controller

import (
	"go.uber.org/zap"
)

func emojiField(emoji string) zap.Field { //nolint:unused
	return zap.String("emoji", emoji)
}
