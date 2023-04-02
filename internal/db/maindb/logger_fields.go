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
	logMessageGetUserByUsername      = "MAIN DB: GET USER BY USERNAME"
	logMessageGetUserTransactions    = "MAIN DB: GET USER TRANSACTIONS"
	logMessageSaveUserTransactions   = "MAIN DB: SAVE USER TRANSACTIONS"
)

func emojiField(emoji string) zap.Field { //nolint:unparam
	return zap.String("emoji", emoji)
}

func dbNameField(dbNameField string) zap.Field {
	return zap.String("db_name", dbNameField)
}

func transactionHashField(transactionHash common.Hash) zap.Field {
	return zap.String("transaction_hash", transactionHash.String())
}

func userIDField(userID int64) zap.Field {
	return zap.Int64("user_id", userID)
}

func usernameField(username string) zap.Field {
	return zap.String("username", username)
}
