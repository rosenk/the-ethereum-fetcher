package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ju-popov/the-ethereum-fetcher/internal/lime/request"
	"github.com/ju-popov/the-ethereum-fetcher/internal/lime/response"
	"github.com/ju-popov/the-ethereum-fetcher/internal/lime/wrapper"
	"github.com/sumup-oss/go-pkgs/errors"
	"golang.org/x/crypto/bcrypt"
)

func (c *Controller) unmarshalAuthenticateRequest(httpRequest *http.Request) (*request.Authenticate, error) {
	var requestStruct request.Authenticate

	if httpRequest.Body != nil {
		body, err := io.ReadAll(httpRequest.Body)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read request body: %s", err.Error())
		}

		if err := json.Unmarshal(body, &requestStruct); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal request body: %s", err.Error())
		}
	}

	if err := c.validate.Struct(requestStruct); err != nil {
		return nil, errors.Wrap(err, "failed to validate request: %s", err.Error())
	}

	return &requestStruct, nil
}

func (c *Controller) Authenticate(httpRequest *http.Request) (*wrapper.HTTPResponse, error) {
	req, err := c.unmarshalAuthenticateRequest(httpRequest)
	if err != nil {
		return nil, wrapper.NewHTTPError(http.StatusBadRequest, err)
	}

	userID, hashedPassword, err := c.mainDBClient.GetUserByUsername(httpRequest.Context(), req.Username)
	if err != nil {
		return nil, wrapper.NewHTTPError(http.StatusInternalServerError, err)
	}

	if (userID == nil) || (hashedPassword == nil) {
		return nil, wrapper.NewHTTPError(http.StatusUnauthorized, errors.New("invalid username or password"))
	}

	if err = bcrypt.CompareHashAndPassword([]byte(*hashedPassword), []byte(req.Password)); err != nil {
		return nil, wrapper.NewHTTPError(http.StatusUnauthorized, errors.New("invalid username or password"))
	}

	token, err := c.jwt.GenerateJWTToken(*userID)
	if err != nil {
		return nil, wrapper.NewHTTPError(http.StatusInternalServerError, err)
	}

	return wrapper.NewHTTPResponse(http.StatusOK, &response.Authenticate{
		Token: *token,
	}), nil
}
