package fetcher

import (
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

const (
	logMessageGetTransaction = "FETCHER: GET TRANSACTION"
)

func emojiField(emoji string) zap.Field {
	return zap.String("emoji", emoji)
}

func transactionHashField(transactionHash common.Hash) zap.Field {
	return zap.String("transaction_hash", transactionHash.String())
}
