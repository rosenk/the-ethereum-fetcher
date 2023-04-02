package controller

import (
	"net/http"

	"github.com/ju-popov/the-ethereum-fetcher/internal/lime/wrapper"
	"github.com/sumup-oss/go-pkgs/errors"
)

func (c *Controller) GetMy(httpRequest *http.Request) (*wrapper.HTTPResponse, error) {
	userID, err := c.jwt.ValidateJWTToken(httpRequest.Header.Get("AUTH_TOKEN"))
	if err != nil {
		return nil, wrapper.NewHTTPError(http.StatusUnauthorized, err)
	}

	transactions, err := c.mainDBClient.GetUserTransactions(httpRequest.Context(), *userID)
	if err != nil {
		return nil, wrapper.NewHTTPError(
			http.StatusInternalServerError,
			errors.Wrap(err, "failed to get user transactions from the main database: %s", err.Error()),
		)
	}

	return c.renderTransactions(transactions)
}
