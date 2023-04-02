package fetcher

import (
	"github.com/ju-popov/the-ethereum-fetcher/internal/db/maindb"
	"github.com/ju-popov/the-ethereum-fetcher/internal/ethereum"
	"github.com/sumup-oss/go-pkgs/logger"
)

type Client struct {
	logger         logger.StructuredLogger
	mainDBClient   *maindb.Client
	ethereumClient *ethereum.Client
}

func NewClient(log logger.StructuredLogger, mainDBClient *maindb.Client, ethereumClient *ethereum.Client) *Client {
	return &Client{
		logger:         log,
		mainDBClient:   mainDBClient,
		ethereumClient: ethereumClient,
	}
}
