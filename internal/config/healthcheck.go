package config

import "time"

type HealthCheckServer struct {
	ListenAddress           *string        `mapstructure:"listen_address"            json:"listenAddress"           validate:"required,notblank"` //nolint:lll
	ReadHeaderTimeout       *time.Duration `mapstructure:"read_header_timeout"       json:"readHeaderTimeout"       validate:"required,notblank"` //nolint:lll
	ReadTimeout             *time.Duration `mapstructure:"read_timeout"              json:"readTimeout"             validate:"required,notblank"` //nolint:lll
	WriteTimeout            *time.Duration `mapstructure:"write_timeout"             json:"writeTimeout"            validate:"required,notblank"` //nolint:lll
	GracefulShutdownTimeout *time.Duration `mapstructure:"graceful_shutdown_timeout" json:"gracefulShutdownTimeout" validate:"required,notblank"` //nolint:lll
}

type HealthCheckController struct {
	ReadyPath *string `mapstructure:"ready_path" json:"readyPath" validate:"required,notblank"`
	LivePath  *string `mapstructure:"live_path"  json:"livePath"  validate:"required,notblank"`
}

type HealthCheck struct {
	Server     *HealthCheckServer     `mapstructure:"server"     json:"server"     validate:"required"`
	Controller *HealthCheckController `mapstructure:"controller" json:"controller" validate:"required"`
}
