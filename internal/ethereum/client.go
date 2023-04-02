package ethereum

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sumup-oss/go-pkgs/errors"
	"github.com/sumup-oss/go-pkgs/logger"
)

type Client struct {
	client  *ethclient.Client
	logger  logger.StructuredLogger
	address string
}

func NewClient(logger logger.StructuredLogger, address string) *Client {
	return &Client{
		logger:  logger,
		address: address,
	}
}

func (c *Client) Connect(ctx context.Context) error {
	c.logger.Info(
		logMessageConnect,
		emojiField("💎"),
	)

	client, err := ethclient.DialContext(ctx, c.address)
	if err != nil {
		return errors.Wrap(err, "failed to connect to ethereum node: %s", err.Error())
	}

	c.client = client

	return nil
}

func (c *Client) Close() {
	c.logger.Info(
		logMessageClose,
		emojiField("💎"),
	)

	c.client.Close()
}
