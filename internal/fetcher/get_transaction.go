package fetcher

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ju-popov/the-ethereum-fetcher/internal/ethereum"
	"github.com/sumup-oss/go-pkgs/errors"
	"github.com/sumup-oss/go-pkgs/logger"
)

//nolint:funlen,cyclop
func (c *Client) GetTransaction(ctx context.Context, txHash common.Hash) (*ethereum.TransactionFull, error) {
	c.logger.Info(
		logMessageGetTransaction,
		transactionHashField(txHash),
		emojiField("📥"),
	)

	dbTx, err := c.mainDBClient.BeginTx(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to start tx: %s", err.Error())
	}

	defer func() {
		if dbTx == nil {
			return
		}

		err = c.mainDBClient.RollbackTx(dbTx)
		if err != nil {
			c.logger.Error(
				logMessageGetTransaction,
				transactionHashField(txHash),
				emojiField("📥"),
				logger.ErrorField(errors.Wrap(err, "failed to rollback tx: %s", err.Error())),
			)
		}
	}()

	transaction, err := c.mainDBClient.GetTransactionCache(ctx, dbTx, txHash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transaction %v: %s", txHash, err.Error())
	}

	if transaction != nil {
		return transaction, nil
	}

	if err := c.mainDBClient.LockTransactionCache(ctx, dbTx, txHash); err != nil {
		return nil, errors.Wrap(err, "failed to lock transaction %v: %s", txHash, err.Error())
	}

	transaction, err = c.mainDBClient.GetTransactionCache(ctx, dbTx, txHash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transaction %v: %s", txHash, err.Error())
	}

	if transaction != nil {
		return transaction, nil
	}

	transaction, err = c.ethereumClient.GetTransactionFull(ctx, txHash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transaction %v: %s", txHash, err.Error())
	}

	if transaction == nil {
		return nil, nil //nolint:nilnil
	}

	err = c.mainDBClient.SaveTransactionCache(ctx, dbTx, *transaction)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save transaction %v: %s", txHash, err.Error())
	}

	err = c.mainDBClient.CommitTx(dbTx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to commit tx: %s", err.Error())
	}

	dbTx = nil

	return transaction, nil
}
