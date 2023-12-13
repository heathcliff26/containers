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
	Body       io.ReadCloser
}

func (e *ErrHttpRequestFailed) Error() string {
	var body string
	b, err := io.ReadAll(e.Body)
	if err != nil {
		body = err.Error()
	} else {
		body = string(b)
	}
	return fmt.Sprintf("HTTP Request returned with Status Code %d, expected 200. Response body: %s", e.StatusCode, body)
}

// Outputs the response received from cloudflare
type ErrOperationFailed struct {
	Result io.ReadCloser
}

func (e *ErrOperationFailed) Error() string {
	var result string
	bytes, err := io.ReadAll(e.Result)
	if err != nil {
		result = err.Error()
	} else {
		result = string(bytes)
	}
	return "Remote api call returned without success, response: " + result
}
