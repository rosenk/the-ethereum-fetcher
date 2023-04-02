package config

import "time"

type LimeServer struct {
	ListenAddress           *string        `mapstructure:"listen_address"            json:"listenAddress"           validate:"required,notblank"` //nolint:lll
	ReadHeaderTimeout       *time.Duration `mapstructure:"read_header_timeout"       json:"readHeaderTimeout"       validate:"required,notblank"` //nolint:lll
	ReadTimeout             *time.Duration `mapstructure:"read_timeout"              json:"readTimeout"             validate:"required,notblank"` //nolint:lll
	WriteTimeout            *time.Duration `mapstructure:"write_timeout"             json:"writeTimeout"            validate:"required,notblank"` //nolint:lll
	GracefulShutdownTimeout *time.Duration `mapstructure:"graceful_shutdown_timeout" json:"gracefulShutdownTimeout" validate:"required,notblank"` //nolint:lll
}

type Lime struct {
	Server *LimeServer `mapstructure:"server" json:"server" validate:"required"`
}
