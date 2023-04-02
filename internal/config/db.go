package config

import "time"

type PostgreSQL struct {
	Host                  *string        `mapstructure:"host"                     json:"host"                  validate:"required,notblank"`                                    //nolint:lll
	Port                  *int           `mapstructure:"port"                     json:"port"                  validate:"required,min=1,max=65535"`                             //nolint:lll
	SSLMode               *string        `mapstructure:"ssl_mode"                 json:"sslMode"               validate:"required,oneof=require verify-full verify-ca disable"` //nolint:lll
	Username              *string        `mapstructure:"username"                 json:"username"              validate:"required,notblank"`                                    //nolint:lll
	Password              *string        `mapstructure:"password"                 json:"password"              validate:"required,notblank"`                                    //nolint:lll
	Database              *string        `mapstructure:"database"                 json:"database"              validate:"required,notblank"`                                    //nolint:lll
	Schema                *string        `mapstructure:"schema"                   json:"schema"                validate:"required,notblank"`                                    //nolint:lll
	Timezone              *string        `mapstructure:"timezone"                 json:"timezone"              validate:"required,notblank"`                                    //nolint:lll
	ConnectTimeoutSeconds *int           `mapstructure:"connect_timeout_seconds"  json:"connectTimeoutSeconds" validate:"required,min=0"`                                       //nolint:lll
	MaxIdleConnections    *int           `mapstructure:"max_idle_connections"     json:"maxIdleConnections"    validate:"required,min=0"`                                       //nolint:lll
	MaxOpenConnections    *int           `mapstructure:"max_open_connections"     json:"maxOpenConnections"    validate:"required,min=0"`                                       //nolint:lll
	ConnectionMaxLifetime *time.Duration `mapstructure:"connection_max_lifetime"  json:"connectionMaxLifetime" validate:"required,min=0"`                                       //nolint:lll
	ConnectionMaxIdleTime *time.Duration `mapstructure:"connection_max_idle_time" json:"connectionMaxIdleTime" validate:"required,min=0"`                                       //nolint:lll
}

type DB struct {
	Name       *string     `mapstructure:"name"        json:"name"       validate:"required,notblank"`
	PostgreSQL *PostgreSQL `mapstructure:"postgresql"  json:"postgresql" validate:"required"`
}
