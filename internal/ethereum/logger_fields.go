package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

const (
	logMessageConnect                = "ETHEREUM: CONNECT"
	logMessageClose                  = "ETHEREUM: CLOSE"
	logMessageGetTransactionOverview = "ETHEREUM: GET TRANSACTION OVERVIEW"
	logMessageGetTransactionReceipt  = "ETHEREUM: GET TRANSACTION RECEIPT"
)

func emojiField(emoji string) zap.Field {
	return zap.String("emoji", emoji)
}

func transactionHashField(transactionHash common.Hash) zap.Field {
	return zap.String("transaction_hash", transactionHash.String())
}
