package controller

import (
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ju-popov/the-ethereum-fetcher/internal/lime/response"
	"github.com/ju-popov/the-ethereum-fetcher/internal/lime/wrapper"
	"github.com/sumup-oss/go-pkgs/errors"
)

func (c *Controller) GetAll(httpRequest *http.Request) (*wrapper.HTTPResponse, error) {
	transactions, err := c.mainDBClient.GetAllTransactionCache(httpRequest.Context())
	if err != nil {
		return nil, wrapper.NewHTTPError(
			http.StatusInternalServerError,
			errors.Wrap(err, "failed to get all transactions from the main database"),
		)
	}

	result := make([]response.GetETH, len(transactions))
	for index, transaction := range transactions {
		result[index] = response.GetETH{
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
