package ethereum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sumup-oss/go-pkgs/errors"
)

type TransactionReceipt struct {
	Status          uint64
	BlockHash       common.Hash
	BlockNumber     *big.Int
	ContractAddress common.Address
	LogsCount       int
}

func (c *Client) GetTransactionReceipt(ctx context.Context, txHash common.Hash) (*TransactionReceipt, error) {
	c.logger.Info(
		logMessageGetTransactionReceipt,
		transactionHashField(txHash),
		emojiField("💎"),
	)

	hash := common.HexToHash(txHash.String())

	receipt, err := c.client.TransactionReceipt(ctx, hash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transaction receipt %v: %s", txHash, err.Error())
	}

	result := &TransactionReceipt{
		Status:          receipt.Status,
		BlockHash:       receipt.BlockHash,
		BlockNumber:     receipt.BlockNumber,
		ContractAddress: receipt.ContractAddress,
		LogsCount:       len(receipt.Logs),
	}

	return result, nil
}
