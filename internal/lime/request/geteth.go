package request

import "github.com/ethereum/go-ethereum/common"

type GetETH struct {
	TransactionHashes []common.Hash `json:"-" validate:"max=8"`
}
