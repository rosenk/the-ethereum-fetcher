package controller

import (
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ju-popov/the-ethereum-fetcher/internal/ethereum"
	"github.com/ju-popov/the-ethereum-fetcher/internal/lime/response"
	"github.com/ju-popov/the-ethereum-fetcher/internal/lime/wrapper"
)

func (c *Controller) renderTransactions(transactions []ethereum.TransactionFull) (*wrapper.HTTPResponse, error) {
	result := make([]response.TransactionsList, len(transactions))

	for index, transaction := range transactions {
		result[index] = response.TransactionsList{
			TransactionHash:   strings.ToLower(transaction.Hash.Hex()),
			TransactionStatus: transaction.Status,
			BlockHash:         strings.ToLower(transaction.BlockHash.Hex()),
			BlockNumber:       transaction.BlockNumber,
			From:              strings.ToLower(transaction.Sender.Hex()),
			LogsCount:         transaction.LogsCount,
			Input:             "0x" + hex.EncodeToString(transaction.Data),
		}

		if transaction.To != nil {
			value := strings.ToLower(transaction.To.Hex())
			result[index].To = &value
		}

		if transaction.ContractAddress != (common.Address{}) {
			value := strings.ToLower(transaction.ContractAddress.Hex())
			result[index].ContractAddress = &value
		}

		if transaction.Value != nil {
			value := transaction.Value.String()
			result[index].Value = &value
		}
	}

	return wrapper.NewHTTPResponse(
		http.StatusOK,
		&result,
	), nil
}
