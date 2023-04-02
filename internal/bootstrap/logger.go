package bootstrap

import (
	"github.com/ju-popov/the-ethereum-fetcher/internal/config"
	"github.com/ju-popov/the-ethereum-fetcher/internal/logbuilder"
	"github.com/sumup-oss/go-pkgs/errors"
	"github.com/sumup-oss/go-pkgs/logger"
)

func Logger(conf *config.Logger) (*logger.ZapLogger, error) {
	log, err := logbuilder.NewBuilder().
		WithLevel(*conf.Level).
		WithEncoding(*conf.Encoding).
		WithStdoutEnabled(*conf.StdoutEnabled).
		WithSyslogEnabled(*conf.SyslogEnabled).
		WithSyslogFacility(*conf.SyslogFacility).
		WithSyslogTag(*conf.SyslogTag).
		Build()
	if err != nil {
		return nil, errors.Wrap(err, "create logger")
	}

	return log, nil
}
