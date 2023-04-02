package postgresql

import (
	"context"
	"strings"

	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/sumup-oss/go-pkgs/logger"
	"go.uber.org/zap"
)

type Adapter struct {
	logger logger.StructuredLogger
}

func NewLoggerAdapter(log logger.StructuredLogger) *Adapter {
	return &Adapter{logger: log}
}

//nolint:cyclop
func (a Adapter) Log(_ context.Context, level sqldblogger.Level, msg string, data map[string]interface{}) {
	fields := make([]zap.Field, 0)

	msg = strings.ToUpper(msg)

	fields = append(fields, emojiField("💽"))

	for key, value := range data {
		switch key {
		case "duration":
			fields = append(fields, durationField(value))
			fields = append(fields, durationHumanField(value))
		case "start":
			fields = append(fields, startField(value))
		case "time_sql":
			fields = append(fields, timeSQLField(value))
		case "query":
			fields = append(fields, queryField(value))
		default:
			fields = append(fields, zap.Any(key, value))
		}
	}

	// upcase msg

	switch level {
	case sqldblogger.LevelError:
		a.logger.Error(msg, fields...)
	case sqldblogger.LevelInfo:
		a.logger.Info(msg, fields...)
	case sqldblogger.LevelDebug:
		a.logger.Debug(msg, fields...)
	case sqldblogger.LevelTrace:
		a.logger.Debug(msg, fields...)
	default:
		// trace will use zap debug
		a.logger.Debug(msg, fields...)
	}
}
