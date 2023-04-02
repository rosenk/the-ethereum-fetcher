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
		return nil, errors.Wrap(err, "failed to start tx: %s", err.Error())
	}

	return tx, nil
}

func (c *Client) CommitTx(dbTx *sql.Tx) error {
	c.logger.Info(
		logMessageCommitTX,
		emojiField("💽"),
		dbNameField(c.name),
	)

	if err := dbTx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit tx: %s", err.Error())
	}

	return nil
}

func (c *Client) RollbackTx(dbTx *sql.Tx) error {
	c.logger.Info(
		logMessageRollbackTX,
		emojiField("💽"),
		dbNameField(c.name),
	)

	if err := dbTx.Rollback(); err != nil {
		return errors.Wrap(err, "failed to rollback tx: %s", err.Error())
	}

	return nil
}
