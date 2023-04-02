package config

import "time"

type Shutdown struct {
	ForcefulShutdownTimeout *time.Duration `mapstructure:"forceful_shutdown_timeout" json:"forcefulShutdownTimeout" validate:"required,notblank"` //nolint:lll
}
