package wrapper

type HTTPError struct {
	statusCode    int
	internalError error
}

func NewHTTPError(statusCode int, internalError error) HTTPError {
	return HTTPError{
		statusCode:    statusCode,
		internalError: internalError,
	}
}

func (e HTTPError) Error() string {
	return e.internalError.Error()
}
