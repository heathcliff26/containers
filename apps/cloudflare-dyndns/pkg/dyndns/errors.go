package dyndns

import (
	"fmt"
	"io"
)

type ErrMissingToken struct{}

func (e ErrMissingToken) Error() string {
	return "No token provided for authenticating with the API."
}

type ErrMissingEndpoint struct{}

func (e ErrMissingEndpoint) Error() string {
	return "No endpoint provided"
}

type ErrInvalidIP struct {
	IP string
}

func (e *ErrInvalidIP) Error() string {
	return fmt.Sprintf("\"%s\" is not a valid ip address", e.IP)
}

type ErrNoIP struct{}

func (e ErrNoIP) Error() string {
	return "Can't update dyndns entry, no IPs provided"
}

type ErrNoDomain struct{}

func (e ErrNoDomain) Error() string {
	return "Can't update dyndns entry, no valid domain provided"
}

// Shows the actual status code, as well as the response body.
// Shows the error instead if it can't read the response body.
type ErrHttpRequestFailed struct {
	StatusCode int
	Body       string
}

func NewErrHttpRequestFailed(status int, resBody io.ReadCloser) *ErrHttpRequestFailed {
	var body string
	b, err := io.ReadAll(resBody)
	if err != nil {
		body = err.Error()
	} else {
		body = string(b)
	}
	return &ErrHttpRequestFailed{
		StatusCode: status,
		Body:       body,
	}
}

func (e *ErrHttpRequestFailed) Error() string {
	return fmt.Sprintf("HTTP Request returned with Status Code %d, expected 200. Response body: %s", e.StatusCode, e.Body)
}

// Outputs the response received from cloudflare
type ErrOperationFailed struct {
	Result string
}

func NewErrOperationFailed(res io.ReadCloser) *ErrOperationFailed {
	var result string
	b, err := io.ReadAll(res)
	if err != nil {
		result = err.Error()
	} else {
		result = string(b)
	}
	return &ErrOperationFailed{
		Result: result,
	}
}

func (e *ErrOperationFailed) Error() string {
	return "Remote api call returned without success, response: " + e.Result
}
