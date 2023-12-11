package client

import (
	"encoding/json"
	"fmt"
	"io"
)

type ErrMissingSecret struct{}

func (e ErrMissingSecret) Error() string {
	return "No secret provided for authenticating with the API."
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
type ErrCloudflareOperationFailed struct {
	result cloudflareResponse
}

func (e *ErrCloudflareOperationFailed) Error() string {
	var result string
	bytes, err := json.Marshal(e.result)
	if err != nil {
		result = err.Error()
	} else {
		result = string(bytes)
	}
	return "Cloudflare api call returned without success, response: " + result
}
