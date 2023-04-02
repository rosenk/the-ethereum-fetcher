package fetcher

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ju-popov/the-ethereum-fetcher/internal/ethereum"
	"github.com/sumup-oss/go-pkgs/errors"
	"github.com/sumup-oss/go-pkgs/logger"
)

func (c *Client) GetTransaction(ctx context.Context, txHash common.Hash) (*ethereum.TransactionOverview, error) {
	c.logger.Info(
		logMessageGetTransaction,
		transactionHashField(txHash),
		emojiField("📥"),
	)

	tx, err := c.mainDBClient.BeginTx(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to start tx")
	}

	defer func() {
		if tx == nil {
			return
		}

		err = c.mainDBClient.RollbackTx(tx)
		if err != nil {
			c.logger.Error(
				logMessageGetTransaction,
				transactionHashField(txHash),
				emojiField("📥"),
				logger.ErrorField(errors.Wrap(err, "failed to rollback tx")),
			)
		}
	}()

	transaction, err := c.mainDBClient.GetTransactionCache(ctx, tx, txHash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transaction %v", txHash)
	}

	if transaction != nil {
		return transaction, nil
	}

	if err := c.mainDBClient.LockTransactionCache(ctx, tx, txHash); err != nil {
		return nil, errors.Wrap(err, "failed to lock transaction %v", txHash)
	}

	transaction, err = c.mainDBClient.GetTransactionCache(ctx, tx, txHash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transaction %v", txHash)
	}

	if transaction != nil {
		return transaction, nil
	}

	transaction, err = c.ethereumClient.GetTransactionOverview(ctx, txHash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transaction %v", txHash)
	}

	if transaction == nil {
		return nil, nil
	}

	err = c.mainDBClient.SaveTransactionCache(ctx, tx, *transaction)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save transaction %v", txHash)
	}

	err = c.mainDBClient.CommitTx(tx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to commit tx")
	}

	tx = nil

	return transaction, nil
}
