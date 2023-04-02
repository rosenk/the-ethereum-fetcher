package shutdownhandler

import (
	"os"
	"time"

	"go.uber.org/zap"
)

const (
	logMessageShutdownSignal = "SHUTDOWN HANDLER: SHUTDOWN SIGNAL RECEIVED"
	logMessageShutdownForce  = "SHUTDOWN HANDLER: FORCING TERMINATION (DEADLINE EXCEEDED)"
)

func emojiField(emoji string) zap.Field {
	return zap.String("emoji", emoji)
}

func signalField(sig os.Signal) zap.Field {
	return zap.String("signal", sig.String())
}

func deadlineHumanField(deadline time.Duration) zap.Field {
	return zap.String("deadline_human", deadline.String())
}
