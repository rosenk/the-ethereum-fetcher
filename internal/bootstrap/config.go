package bootstrap

import (
	"github.com/go-playground/validator/v10"
	"github.com/ju-popov/the-ethereum-fetcher/internal/config"
	"github.com/sumup-oss/go-pkgs/errors"
)

func Config(validate *validator.Validate, configFilename string) (*config.Config, error) {
	conf, err := config.ReadConfig(configFilename)
	if err != nil {
		return nil, errors.Wrap(err, "read config: %s", err.Error())
	}

	err = validate.Struct(conf)
	if err != nil {
		return nil, errors.Wrap(err, "validate config: %s", err.Error())
	}

	return conf, nil
}
