package bootstrap

import (
	"github.com/ju-popov/the-ethereum-fetcher/internal/config"
	"github.com/ju-popov/the-ethereum-fetcher/internal/jwt"
	"github.com/sumup-oss/go-pkgs/logger"
)

func JWT(
	log logger.StructuredLogger,
	conf *config.JWT,
) *jwt.Client {
	return jwt.NewClient(log, *conf.Duration, *conf.SigningKey)
}
