package maindb

import (
	"database/sql"

	"github.com/sumup-oss/go-pkgs/errors"
	"github.com/sumup-oss/go-pkgs/logger"
)

type Client struct {
	logger logger.StructuredLogger
	name   string
	db     *sql.DB
}

func NewClient(log logger.StructuredLogger, name string, db *sql.DB) *Client {
	return &Client{
		logger: log,
		name:   name,
		db:     db,
	}
}

func (c *Client) Ping() error {
	c.logger.Info(
		logMessagePing,
		emojiField("💽"),
		dbNameField(c.name),
	)

	err := c.db.Ping()
	if err != nil {
		return errors.Wrap(err, "failed to ping postgresql server: %s", err.Error())
	}

	return nil
}

func (c *Client) Close() error {
	c.logger.Info(
		logMessageClose,
		emojiField("💽"),
		dbNameField(c.name),
	)

	err := c.db.Close()
	if err != nil {
		return errors.Wrap(err, "failed to close postgresql server: %s", err.Error())
	}

	return nil
}
