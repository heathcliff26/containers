package client

import (
	"encoding/json"
	"fmt"
	"io"
)

type MissingSecretError struct{}

func (e MissingSecretError) Error() string {
	return "No secret provided for authenticating with the API."
}

type InvalidIPError struct {
	ip string
}

func NewInvalidIPError(ip string) error {
	return &InvalidIPError{
		ip: ip,
	}
}

func (e *InvalidIPError) Error() string {
	return fmt.Sprintf("\"%s\" is not a valid ip address", e.ip)
}

type NoIPError struct{}

func (e NoIPError) Error() string {
	return "Can't update dyndns entry, no IPs provided"
}

type NoDomainError struct{}

func (e NoDomainError) Error() string {
	return "Can't update dyndns entry, no valid domain provided"
}

// Shows the actual status code, as well as the response body.
// Shows the error instead if it can't read the response body.
type HttpRequestFailedError struct {
	StatusCode int
	Body       io.ReadCloser
}

func (e *HttpRequestFailedError) Error() string {
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
type CloudflareOperationFailedError struct {
	result cloudflareResponse
}

func (e *CloudflareOperationFailedError) Error() string {
	var result string
	bytes, err := json.Marshal(e.result)
	if err != nil {
		result = err.Error()
	} else {
		result = string(bytes)
	}
	return "Cloudflare api call returned without success, response: " + result
}
