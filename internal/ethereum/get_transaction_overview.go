package ethereum

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type TransactionOverview struct {
	Hash            common.Hash     `json:"hash"`
	Status          uint64          `json:"status"`
	BlockHash       common.Hash     `json:"blockHash"`
	BlockNumber     *big.Int        `json:"blockNumber"`
	Sender          common.Address  `json:"sender"`
	To              *common.Address `json:"to"`
	ContractAddress common.Address  `json:"contractAddress"`
	LogsCount       int             `json:"logsCount"`
	Data            []byte          `json:"data"`
	Value           *big.Int        `json:"value"`
}

func (c *Client) GetTransactionOverview(ctx context.Context, txHash common.Hash) (*TransactionOverview, error) {
	c.logger.Info(
		logMessageGetTransactionOverview,
		transactionHashField(txHash),
		emojiField("💎"),
	)

	commonData, err := c.GetTransaction(ctx, txHash)
	if err != nil {
		return nil, err
	}

	if commonData == nil {
		return nil, nil
	}

	receiptData, err := c.GetTransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}

	return &TransactionOverview{
		Hash:            commonData.Hash,
		Status:          receiptData.Status,
		BlockHash:       receiptData.BlockHash,
		BlockNumber:     receiptData.BlockNumber,
		Sender:          commonData.Sender,
		To:              commonData.To,
		ContractAddress: receiptData.ContractAddress,
		LogsCount:       receiptData.LogsCount,
		Data:            commonData.Data,
		Value:           commonData.Value,
	}, nil
}
