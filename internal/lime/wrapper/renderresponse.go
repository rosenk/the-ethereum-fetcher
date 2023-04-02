package wrapper

import (
	"encoding/json"
	"net/http"

	"github.com/sumup-oss/go-pkgs/errors"
)

const (
	contentTypeHeader    = "Content-Type"
	contentTypeTextPlain = "application/json; charset=utf-8"
)

func (w *Wrapper) renderResponse(writer http.ResponseWriter, response *HTTPResponse) error {
	writer.Header().Set(contentTypeHeader, contentTypeTextPlain)
	writer.WriteHeader(response.statusCode)

	if response.body == nil {
		return nil
	}

	data, err := json.Marshal(response.body)
	if err != nil {
		return errors.Wrap(err, "failed to marshal response body: %s", err.Error())
	}

	_, err = writer.Write(data)
	if err != nil {
		return errors.Wrap(err, "failed to write response body: %s", err.Error())
	}

	return nil
}
