package logbuilder

import (
	"github.com/sumup-oss/go-pkgs/errors"
	"github.com/sumup-oss/go-pkgs/logger"
	"go.uber.org/zap/zapcore"
)

type Builder struct {
	level          string
	encoding       string
	stdoutEnabled  bool
	syslogEnabled  bool
	syslogFacility string
	syslogTag      string
	fields         []zapcore.Field
}

func NewBuilder() *Builder {
	return &Builder{
		level:          "INFO",
		encoding:       "json",
		stdoutEnabled:  true,
		syslogEnabled:  false,
		syslogFacility: "LOCAL0",
		syslogTag:      "the-ethereum-fetcher",
	}
}

func (b *Builder) WithLevel(level string) *Builder {
	b.level = level

	return b
}

func (b *Builder) WithEncoding(encoding string) *Builder {
	b.encoding = encoding

	return b
}

func (b *Builder) WithStdoutEnabled(stdoutEnabled bool) *Builder {
	b.stdoutEnabled = stdoutEnabled

	return b
}

func (b *Builder) WithSyslogEnabled(syslogEnabled bool) *Builder {
	b.syslogEnabled = syslogEnabled

	return b
}

func (b *Builder) WithSyslogFacility(syslogFacility string) *Builder {
	b.syslogFacility = syslogFacility

	return b
}

func (b *Builder) WithSyslogTag(syslogTag string) *Builder {
	b.syslogTag = syslogTag

	return b
}

func (b *Builder) WithField(field zapcore.Field) *Builder {
	b.fields = append(b.fields, field)

	return b
}

func (b *Builder) Build() (*logger.ZapLogger, error) {
	log, err := logger.NewZapLogger(logger.Configuration{
		Level:          b.level,
		Encoding:       b.encoding,
		StdoutEnabled:  b.stdoutEnabled,
		SyslogEnabled:  b.syslogEnabled,
		SyslogFacility: b.syslogFacility,
		SyslogTag:      b.syslogTag,
		Fields:         b.fields,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to build logger: %s", err.Error())
	}

	return log, nil
}

func (b *Builder) BuildStructuredNopLogger(level string) *logger.StructuredNopLogger {
	return logger.NewStructuredNopLogger(level)
}
