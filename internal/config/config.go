package config

import (
	"strings"

	"github.com/spf13/viper"
	"github.com/sumup-oss/go-pkgs/errors"
)

const (
	envPrefix = "slowpay"
	configDir = "slowpay"
)

// force_shutdown_timeout

type Config struct {
	Environment *string      `mapstructure:"environment" json:"environment" validate:"required,notblank,max=128"`
	Logger      *Logger      `mapstructure:"logger"      json:"logger"      validate:"required"`
	HealthCheck *HealthCheck `mapstructure:"healthcheck" json:"healthcheck" validate:"required"`
	Metrics     *Metrics     `mapstructure:"metrics"     json:"metrics"     validate:"required"`
	Shutdown    *Shutdown    `mapstructure:"shutdown"    json:"shutdown"    validate:"required"`
	DB          *DB          `mapstructure:"db"          json:"db"          validate:"required"`
	Ethereum    *Ethereum    `mapstructure:"ethereum"    json:"ethereum"    validate:"required"`
	Lime        *Lime        `mapstructure:"lime"        json:"lime"        validate:"required"`
}

func ReadConfig(filename string) (*Config, error) {
	setOptions(filename)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "read config")
	}

	config := Config{} //nolint:exhaustruct
	if err := viper.Unmarshal(&config); err != nil {
		return nil, errors.Wrap(err, "unmarshal config")
	}

	return &config, nil
}

func setOptions(filename string) {
	if filename != "" {
		viper.SetConfigFile(filename)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("/etc/" + configDir)
		viper.AddConfigPath("$HOME/." + configDir)
		viper.AddConfigPath(".")
	}

	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.AutomaticEnv()
}
