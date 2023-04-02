package wrapper

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sumup-oss/go-pkgs/errors"
)

type errorResponseDetails struct {
	Description string `json:"description,omitempty"`
	Field       string `json:"field,omitempty"`
	Reason      string `json:"reason,omitempty"`
}

type errorResponse struct {
	Description string                 `json:"description,omitempty"`
	Details     []errorResponseDetails `json:"details,omitempty"`
	Error       string                 `json:"error"`
}

func parameterize(value string, separator string, preserveCase bool) string {
	value = regexp.MustCompile(`[^a-zA-Z0-9-_]+`).ReplaceAllString(value, separator)
	value = regexp.MustCompile(regexp.QuoteMeta(separator)+`{2,}`).ReplaceAllString(value, separator)
	value = strings.Trim(value, separator)

	if !preserveCase {
		value = strings.ToLower(value)
	}

	return value
}

func (w *Wrapper) renderError(writer http.ResponseWriter, err error) error {
	var httpError HTTPError
	if errors.As(err, &httpError) {
		return w.renderHTTPError(writer, httpError)
	}

	return w.renderHTTPError(
		writer,
		NewHTTPError(http.StatusInternalServerError, err),
	)
}

func (w *Wrapper) renderHTTPError(writer http.ResponseWriter, err HTTPError) error {
	var (
		details          []errorResponseDetails
		validationErrors validator.ValidationErrors
	)

	if errors.As(err.internalError, &validationErrors) {
		for _, fieldError := range validationErrors {
			field := fieldError.Namespace()
			if index := strings.Index(field, "."); index > -1 {
				field = field[index+1:]
			}

			details = append(details, errorResponseDetails{
				Description: fieldError.Error(),
				Field:       field,
				Reason:      fieldError.ActualTag(),
			})
		}
	}

	return w.renderResponse(
		writer,
		&HTTPResponse{
			statusCode: err.statusCode,
			body: &errorResponse{
				Description: err.Error(),
				Details:     details,
				Error:       parameterize(strings.ToUpper(http.StatusText(err.statusCode)), "_", true),
			},
		},
	)
}
