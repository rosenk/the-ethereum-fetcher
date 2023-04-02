package cmd

import (
	"go.uber.org/zap"
)

const (
	logMessageEthereumClientConnectError = "ETHEREUM CLIENT: CONNECT ERROR"
	logMessageMainDBPingError            = "MAIN DB: PING ERROR"
	logMessageMainDBMigrateError         = "MAIN DB: MIGRATE ERROR"
	logMessageCommandError               = "COMMAND ERROR"
	logMessageTaskGroupStart             = "TASK GROUP: STARTING"
	logMessageTaskGroupError             = "TASK GROUP: ERROR"
	logMessageTaskGroupShutdown          = "TASK GROUP: SHUTDOWN"
)

func emojiField(emoji string) zap.Field {
	return zap.String("emoji", emoji)
}
