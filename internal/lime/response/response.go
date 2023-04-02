package response

import "math/big"

type GetETH struct {
	TransactionHash   string   `json:"transactionHash"`
	TransactionStatus uint64   `json:"transactionStatus"`
	BlockHash         string   `json:"blockHash"`
	BlockNumber       *big.Int `json:"blockNumber"`
	From              string   `json:"from"`
	To                *string  `json:"to"`
	ContractAddress   *string  `json:"contractAddress"`
	LogsCount         int      `json:"logsCount"`
	Input             string   `json:"input"`
	Value             *string  `json:"value"`
}
