package wrapper

type HTTPResponse struct {
	statusCode int
	body       interface{}
}

func NewHTTPResponse(statusCode int, body interface{}) *HTTPResponse {
	return &HTTPResponse{
		statusCode: statusCode,
		body:       body,
	}
}
