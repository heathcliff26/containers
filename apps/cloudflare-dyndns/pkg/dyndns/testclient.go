package dyndns

import (
	"time"
)

// This is a stub implementation of Client, it is only meant to be used for testing
type testClient struct {
	data *ClientData
}

// Create a new testClient, fails if the token is empty
func NewTestClient(token string, proxy bool) (Client, error) {
	if token == "" {
		return nil, ErrMissingSecret{}
	}
	return &testClient{
		data: NewClientData(proxy),
	}, nil
}

// Give Access to ClientData
func (c *testClient) Data() *ClientData {
	return c.data
}

// Stub implementation, does initial check regarding IP and domains
func (c *testClient) Update() error {
	if c.Data().IPv4() == "" && c.Data().IPv6() == "" {
		return ErrNoIP{}
	}
	if c.Data().Domains() == nil || len(c.Data().Domains()) == 0 {
		return ErrNoDomain{}
	}
	return nil
}

// Stub implementation, not useful for testing
func (c *testClient) Run(_ time.Duration) {}
