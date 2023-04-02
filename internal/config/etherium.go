package config

type Ethereum struct {
	Address *string `mapstructure:"address" json:"address" validate:"required,notblank"`
}
