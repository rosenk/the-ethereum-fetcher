package bootstrap

import (
	"github.com/ju-popov/the-ethereum-fetcher/internal/config"
	"github.com/ju-popov/the-ethereum-fetcher/internal/ethereum"
	"github.com/sumup-oss/go-pkgs/logger"
)

func EthereumClient(
	log logger.StructuredLogger,
	conf *config.Ethereum,
) *ethereum.Client {
	return ethereum.NewClient(log, *conf.Address)
}
