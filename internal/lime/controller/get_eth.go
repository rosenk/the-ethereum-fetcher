package controller

import (
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi/v5"
	"github.com/ju-popov/the-ethereum-fetcher/internal/ethereum"
	"github.com/ju-popov/the-ethereum-fetcher/internal/lime/request"
	"github.com/ju-popov/the-ethereum-fetcher/internal/lime/response"
	"github.com/ju-popov/the-ethereum-fetcher/internal/lime/wrapper"
	"github.com/sumup-oss/go-pkgs/errors"
)

func (c *Controller) unmarshalGetETHRequest(httpRequest *http.Request) (*request.GetETH, error) {
	rlpHEX := chi.URLParam(httpRequest, "rlphex")

	transactionHashes, err := ethereum.DecodeRLPEncodedHashes(rlpHEX)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode RLP encoded hashes")
	}

	transactionHashes = append(transactionHashes, []common.Hash{
		common.HexToHash("0x9b2f6a3c2e1aed2cccf92ba666c22d053ad0d8a5da7aa1fd5477dcd6577b4524"),
		common.HexToHash("0x5a57e3051cb92e2d482515b07e7b3d1851722a74654657bd64a14c39ca3f9cf2"),
		common.HexToHash("0x71b9e2b44d40498c08a62988fac776d0eac0b5b9613c37f9f6f9a4b888a8b057"),
		common.HexToHash("0xc5f96bf1b54d3314425d2379bd77d7ed4e644f7c6e849a74832028b328d4d798"),
	}...)

	requestStruct := request.GetETH{
		TransactionHashes: transactionHashes,
	}

	if err := c.validate.Struct(requestStruct); err != nil {
		return nil, errors.Wrap(err, "failed to validate request")
	}

	return &requestStruct, nil
}

func (c *Controller) GetETH(httpRequest *http.Request) (*wrapper.HTTPResponse, error) {
	req, err := c.unmarshalGetETHRequest(httpRequest)
	if err != nil {
		return nil, wrapper.NewHTTPError(http.StatusBadRequest, err)
	}

	transactions, err := c.fetcherClient.GetTransactions(httpRequest.Context(), req.TransactionHashes)
	if err != nil {
		return nil, wrapper.NewHTTPError(
			http.StatusInternalServerError,
			errors.Wrap(err, "failed to get transactions"),
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
