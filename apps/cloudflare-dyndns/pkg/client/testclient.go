package client

import "time"

// This is a stub implementation of DyndnsClient, it is only meant to be used for testing
type testClient struct {
	ipv4, ipv6 string
	domains    []string
	proxy      bool
}

// Create a new testClient, fails if the token is empty
func NewTestClient(token string, proxy bool) (DyndnsClient, error) {
	if token == "" {
		return nil, ErrMissingSecret{}
	}
	return &testClient{proxy: proxy}, nil
}

// Check if DNS entries will be set to proxied
func (c *testClient) Proxy() bool {
	return c.proxy
}

func (c *testClient) Domains() []string {
	return c.domains
}

func (c *testClient) IPv4() string {
	return c.ipv4
}

func (c *testClient) IPv6() string {
	return c.ipv6
}

// Set domains to update, overrides previous domains
func (c *testClient) SetDomains(domains []string) {
	c.domains = domains
}

// Add a domain to the list of domains
func (c *testClient) AddDomain(domain string) {
	if c.domains == nil {
		c.domains = []string{domain}
	} else {
		c.domains = append(c.domains, domain)
	}
}

// Set IPv4 Address to use for creating entries
// Returns an error if string is neither empty nor a valid IP Address
func (c *testClient) SetIPv4(val string) error {
	if val != "" && !validIPv4(val) {
		return &ErrInvalidIP{val}
	}
	c.ipv4 = val
	return nil
}

// Set IPv6 Address to use for creating entries
// Returns an error if string is neither empty nor a valid IP Address
func (c *testClient) SetIPv6(val string) error {
	if val != "" && !validIPv6(val) {
		return &ErrInvalidIP{val}
	}
	c.ipv6 = val
	return nil
}

// Stub implementation, does initial check regarding IP and domains
func (c *testClient) Update() error {
	if c.ipv4 == "" && c.ipv6 == "" {
		return ErrNoIP{}
	}
	if c.domains == nil || len(c.domains) == 0 {
		return ErrNoDomain{}
	}
	return nil
}

// Stub implementation, not useful for testing
func (c *testClient) Run(_ time.Duration) {}
