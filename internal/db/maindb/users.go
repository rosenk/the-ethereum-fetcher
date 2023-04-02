package maindb

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/ju-popov/the-ethereum-fetcher/internal/ethereum"
	"github.com/sumup-oss/go-pkgs/errors"
	"github.com/sumup-oss/go-pkgs/logger"
)

func (c *Client) GetUserByUsername(ctx context.Context, username string) (*int64, *string, error) {
	c.logger.Info(
		logMessageGetUserByUsername,
		emojiField("💽"),
		dbNameField(c.name),
		usernameField(username),
	)

	var (
		userID         int64
		hashedPassword string
	)

	err := c.db.QueryRowContext(ctx, `
		SELECT id, hashed_password
		FROM users
		WHERE username = $1
	`, username).Scan(&userID, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil, nil
		}

		return nil, nil, errors.Wrap(err, "failed to get user: %s", err.Error())
	}

	return &userID, &hashedPassword, nil
}

func (c *Client) GetUserTransactions(ctx context.Context, userID int64) ([]ethereum.TransactionFull, error) {
	c.logger.Info(
		logMessageGetUserTransactions,
		emojiField("💽"),
		dbNameField(c.name),
		userIDField(userID),
	)

	rows, err := c.db.QueryContext(ctx, `
		SELECT data
		FROM user_transactions
		INNER JOIN transactions_cache ON user_transactions.transaction_hash = transactions_cache.hash
		WHERE user_transactions.user_id = $1
        ORDER BY user_transactions.id DESC
	`, userID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user transactions: %s", err.Error())
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

func (c *Client) SaveUserTransactions(
	ctx context.Context,
	userID int64,
	transactions []ethereum.TransactionFull,
) error {
	c.logger.Info(
		logMessageSaveUserTransactions,
		emojiField("💽"),
		dbNameField(c.name),
		userIDField(userID),
	)

	for _, transaction := range transactions {
		_, err := c.db.ExecContext(ctx, `
		INSERT INTO user_transactions (user_id, transaction_hash)
		VALUES ($1, $2)
		ON CONFLICT (user_id, transaction_hash) DO NOTHING
	`, userID, transaction.Hash.String())
		if err != nil {
			return errors.Wrap(err, "failed to save user transaction: %s", err.Error())
		}
	}

	return nil
}
