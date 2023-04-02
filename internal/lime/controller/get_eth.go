package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ju-popov/the-ethereum-fetcher/internal/ethereum"
	"github.com/ju-popov/the-ethereum-fetcher/internal/lime/request"
	"github.com/ju-popov/the-ethereum-fetcher/internal/lime/wrapper"
	"github.com/sumup-oss/go-pkgs/errors"
)

func (c *Controller) unmarshalGetETHRequest(httpRequest *http.Request) (*request.GetETH, error) {
	rlpHEX := chi.URLParam(httpRequest, "rlphex")

	transactionHashes, err := ethereum.DecodeRLPEncodedHashes(rlpHEX)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode RLP encoded hashes: %s", err.Error())
	}

	// transactionHashes = append(transactionHashes, []common.Hash{
	//	common.HexToHash("0x9b2f6a3c2e1aed2cccf92ba666c22d053ad0d8a5da7aa1fd5477dcd6577b4524"),
	//	common.HexToHash("0x5a57e3051cb92e2d482515b07e7b3d1851722a74654657bd64a14c39ca3f9cf2"),
	//	common.HexToHash("0x71b9e2b44d40498c08a62988fac776d0eac0b5b9613c37f9f6f9a4b888a8b057"),
	//	common.HexToHash("0xc5f96bf1b54d3314425d2379bd77d7ed4e644f7c6e849a74832028b328d4d798"),
	// }...)

	requestStruct := request.GetETH{
		TransactionHashes: transactionHashes,
	}

	if err := c.validate.Struct(requestStruct); err != nil {
		return nil, errors.Wrap(err, "failed to validate request: %s", err.Error())
	}

	return &requestStruct, nil
}

func (c *Controller) GetETH(httpRequest *http.Request) (*wrapper.HTTPResponse, error) {
	var userID *int64

	if authToken := httpRequest.Header.Get("AUTH_TOKEN"); authToken != "" {
		var err error

		userID, err = c.jwt.ValidateJWTToken(authToken)
		if err != nil {
			return nil, wrapper.NewHTTPError(http.StatusUnauthorized, err)
		}
	}

	req, err := c.unmarshalGetETHRequest(httpRequest)
	if err != nil {
		return nil, wrapper.NewHTTPError(http.StatusBadRequest, err)
	}

	transactions, err := c.fetcherClient.GetTransactionsSimultaneously(httpRequest.Context(), req.TransactionHashes)
	if err != nil {
		return nil, wrapper.NewHTTPError(
			http.StatusInternalServerError,
			errors.Wrap(err, "failed to get transactions: %s", err.Error()),
		)
	}

	if userID != nil {
		if err := c.mainDBClient.SaveUserTransactions(httpRequest.Context(), *userID, transactions); err != nil {
			return nil, wrapper.NewHTTPError(
				http.StatusInternalServerError,
				errors.Wrap(err, "failed to save user transactions to the main database: %s", err.Error()),
			)
		}
	}

	return c.renderTransactions(transactions)
}
