package postgresql

import (
	"regexp"
	"strings"
	"time"

	"go.uber.org/zap"
)

func nameField(name string) zap.Field {
	return zap.String("name", name)
}

func emojiField(emoji string) zap.Field {
	return zap.String("emoji", emoji)
}

func durationField(delay interface{}) zap.Field {
	nanoSeconds, ok := (delay).(float64)
	if ok {
		return zap.Int64("duration", int64(nanoSeconds))
	}

	return zap.Skip()
}

func durationHumanField(delay interface{}) zap.Field {
	nanoSeconds, ok := (delay).(float64)
	if ok {
		return zap.String("duration_human", (time.Duration(nanoSeconds) * time.Nanosecond).String())
	}

	return zap.Skip()
}

func timeField(name string, value interface{}) zap.Field {
	stringValue, ok := value.(string)
	if ok {
		t, err := time.Parse(time.RFC3339Nano, stringValue)
		if err == nil {
			return zap.String(name, t.Format("2006-01-02T15:04:05.000Z0700"))
		}
	}

	return zap.Skip()
}

func startField(start interface{}) zap.Field {
	return timeField("start", start)
}

func timeSQLField(timeSQL interface{}) zap.Field {
	return timeField("time_sql", timeSQL)
}

func queryField(query interface{}) zap.Field {
	stringQuery, ok := query.(string)
	if ok {
		return zap.String("query", strings.Trim(regexp.MustCompile(`\s+`).ReplaceAllString(stringQuery, " "), " "))
	}

	return zap.Skip()
}
