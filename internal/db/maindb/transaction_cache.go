package maindb

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ju-popov/the-ethereum-fetcher/internal/ethereum"
	"github.com/sumup-oss/go-pkgs/errors"
	"github.com/sumup-oss/go-pkgs/logger"
)

func (c *Client) LockTransactionCache(ctx context.Context, dbTx *sql.Tx, hash common.Hash) error {
	c.logger.Info(
		logMessageLockTransactionCache,
		emojiField("💽"),
		dbNameField(c.name),
		transactionHashField(hash),
	)

	_, err := dbTx.ExecContext(ctx, "SELECT pg_advisory_xact_lock($1)", hash.Big().Int64())
	if err != nil {
		return errors.Wrap(err, "failed to lock transaction: %s", err.Error())
	}

	return nil
}

func (c *Client) GetTransactionCache(
	ctx context.Context,
	dbTx *sql.Tx,
	hash common.Hash,
) (*ethereum.TransactionFull, error) {
	c.logger.Info(
		logMessageGetTransactionCache,
		emojiField("💽"),
		dbNameField(c.name),
		transactionHashField(hash),
	)

	var data json.RawMessage

	err := dbTx.QueryRowContext(ctx, `
		SELECT data
		FROM transactions_cache
		WHERE hash = $1
	`, hash.String()).Scan(&data)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil //nolint:nilnil
		}

		return nil, errors.Wrap(err, "failed to get transaction: %s", err.Error())
	}

	var result ethereum.TransactionFull
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal transaction: %s", err.Error())
	}

	return &result, nil
}

func (c *Client) SaveTransactionCache(ctx context.Context, dbTx *sql.Tx, transaction ethereum.TransactionFull) error {
	c.logger.Info(
		logMessageSaveTransactionCache,
		emojiField("💽"),
		dbNameField(c.name),
		transactionHashField(transaction.Hash),
	)

	data, err := json.Marshal(transaction)
	if err != nil {
		return errors.Wrap(err, "failed to marshal transaction: %s", err.Error())
	}

	_, err = dbTx.ExecContext(ctx, `
		INSERT INTO transactions_cache (hash, data, created_at, updated_at)
		VALUES ($1, $2, NOW(), NOW())
		ON CONFLICT (hash) DO UPDATE SET data = $2, updated_at = NOW()
	`, transaction.Hash.String(), data)
	if err != nil {
		return errors.Wrap(err, "failed to save transaction: %s", err.Error())
	}

	return nil
}

func (c *Client) GetAllTransactionCache(ctx context.Context) ([]ethereum.TransactionFull, error) {
	c.logger.Info(
		logMessageGetAllTransactionCache,
		emojiField("💽"),
		dbNameField(c.name),
	)

	rows, err := c.db.QueryContext(ctx, `
		SELECT data
		FROM transactions_cache
	`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all transactions: %s", err.Error())
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			c.logger.Error(
				logMessageCloseRowsError,
				emojiField("💽"),
				dbNameField(c.name),
				logger.ErrorField(err),
			)
		}
	}(rows)

	results := make([]ethereum.TransactionFull, 0)

	for rows.Next() {
		var data json.RawMessage

		err := rows.Scan(&data)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan transaction: %s", err.Error())
		}

		var result ethereum.TransactionFull
		if err := json.Unmarshal(data, &result); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal transaction: %s", err.Error())
		}

		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "failed during row iteration: %s", err.Error())
	}

	return results, nil
}
