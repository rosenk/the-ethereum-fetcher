package maindb

import (
	"context"
	"database/sql"

	"github.com/sumup-oss/go-pkgs/errors"
)

func (c *Client) BeginTx(ctx context.Context) (*sql.Tx, error) {
	c.logger.Info(
		logMessageBeginTX,
		emojiField("💽"),
		dbNameField(c.name),
	)

	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to start tx")
	}

	return tx, nil
}

func (c *Client) CommitTx(tx *sql.Tx) error {
	c.logger.Info(
		logMessageCommitTX,
		emojiField("💽"),
		dbNameField(c.name),
	)

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit tx")
	}

	return nil
}

func (c *Client) RollbackTx(tx *sql.Tx) error {
	c.logger.Info(
		logMessageRollbackTX,
		emojiField("💽"),
		dbNameField(c.name),
	)

	if err := tx.Rollback(); err != nil {
		return errors.Wrap(err, "failed to rollback tx")
	}

	return nil
}
