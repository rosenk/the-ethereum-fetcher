package maindb

import (
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

const (
	logMessageRunningMigrations      = "MAIN DB: RUNNING MIGRATIONS"
	logMessagePing                   = "MAIN DB: PING"
	logMessageClose                  = "MAIN DB: CLOSE"
	logMessageBeginTX                = "MAIN DB: BEGIN TX"
	logMessageCommitTX               = "MAIN DB: COMMIT TX"
	logMessageRollbackTX             = "MAIN DB: ROLLBACK TX"
	logMessageLockTransactionCache   = "MAIN DB: LOCK TRANSACTION CACHE"
	logMessageSaveTransactionCache   = "MAIN DB: SAVE TRANSACTION CACHE"
	logMessageGetTransactionCache    = "MAIN DB: GET TRANSACTION CACHE"
	logMessageGetAllTransactionCache = "MAIN DB: GET ALL TRANSACTION CACHE"
	logMessageCloseRowsError         = "MAIN DB: CLOSE ROWS ERROR"
)

func emojiField(emoji string) zap.Field {
	return zap.String("emoji", emoji)
}

func dbNameField(dbNameField string) zap.Field {
	return zap.String("db_name", dbNameField)
}

func transactionHashField(transactionHash common.Hash) zap.Field {
	return zap.String("transaction_hash", transactionHash.String())
}
