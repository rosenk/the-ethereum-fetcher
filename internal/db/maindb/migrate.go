package maindb

import (
	"github.com/sumup-oss/go-pkgs/errors"
)

func (c *Client) Migrate() error {
	c.logger.Info(
		logMessageRunningMigrations,
		emojiField("💽"),
		dbNameField(c.name),
	)

	_, err := c.db.Exec(`
		CREATE TABLE IF NOT EXISTS transactions_cache (
			hash TEXT PRIMARY KEY,
			data JSONB NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);
	`)
	if err != nil {
		return errors.Wrap(err, "failed to create transactions table")
	}

	return nil
}
