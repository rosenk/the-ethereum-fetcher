package bootstrap

import (
	"github.com/ju-popov/the-ethereum-fetcher/internal/db/maindb"
	"github.com/ju-popov/the-ethereum-fetcher/internal/ethereum"
	"github.com/ju-popov/the-ethereum-fetcher/internal/fetcher"
	"github.com/sumup-oss/go-pkgs/logger"
)

func Fetcher(
	log logger.StructuredLogger,
	mainDBClient *maindb.Client,
	ethereumClient *ethereum.Client,
) *fetcher.Client {
	return fetcher.NewClient(log, mainDBClient, ethereumClient)
}
