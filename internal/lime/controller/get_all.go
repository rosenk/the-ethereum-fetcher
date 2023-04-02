package controller

import (
	"net/http"

	"github.com/ju-popov/the-ethereum-fetcher/internal/lime/wrapper"
	"github.com/sumup-oss/go-pkgs/errors"
)

func (c *Controller) GetAll(httpRequest *http.Request) (*wrapper.HTTPResponse, error) {
	transactions, err := c.mainDBClient.GetAllTransactionCache(httpRequest.Context())
	if err != nil {
		return nil, wrapper.NewHTTPError(
			http.StatusInternalServerError,
			errors.Wrap(err, "failed to get all transactions from the main database: %s", err.Error()),
		)
	}

	return c.renderTransactions(transactions)
}
