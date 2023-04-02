package config

import "time"

type MetricsServer struct {
	ListenAddress           *string        `mapstructure:"listen_address"            json:"listenAddress"           validate:"required,notblank"` //nolint:lll
	ReadHeaderTimeout       *time.Duration `mapstructure:"read_header_timeout"       json:"readHeaderTimeout"       validate:"required,notblank"` //nolint:lll
	ReadTimeout             *time.Duration `mapstructure:"read_timeout"              json:"readTimeout"             validate:"required,notblank"` //nolint:lll
	WriteTimeout            *time.Duration `mapstructure:"write_timeout"             json:"writeTimeout"            validate:"required,notblank"` //nolint:lll
	GracefulShutdownTimeout *time.Duration `mapstructure:"graceful_shutdown_timeout" json:"gracefulShutdownTimeout" validate:"required,notblank"` //nolint:lll
}

type MetricsController struct {
	Path *string `mapstructure:"path" json:"path" validate:"required,notblank"`
}

type Metrics struct {
	Server     *MetricsServer     `mapstructure:"server"     json:"server"    validate:"required"`
	Controller *MetricsController `mapstructure:"controller" json:"controller" validate:"required"`
}
