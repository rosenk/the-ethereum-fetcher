package config

import "time"

type JWT struct {
	Duration   *time.Duration `mapstructure:"duration"    json:"duration"   validate:"required,notblank"`
	SigningKey *string        `mapstructure:"signing_key" json:"signingKey" validate:"required,notblank"`
}
