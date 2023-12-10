package client

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const CLOUDFLARE_API_ENDPOINT = "https://api.cloudflare.com/client/v4/"

type DyndnsClient interface {
	Proxy() bool
	Domains() []string
	IPv4() string
	IPv6() string
	SetDomains([]string)
	AddDomain(string)
	SetIPv4(string) error
	SetIPv6(string) error
	Update() error
	Run(time.Duration)
}

// Implements DyndnsClient
type cloudflareClient struct {
	// API is only here as a variable to enable local testing without relying on the cloudflare API
	endpoint string
	token    string
	proxy    bool
	domains  []string
	ipv4     string
	ipv6     string
}

// Create a new CloudflareClient and test if the token is valid
func NewCloudflareClient(token string, proxy bool) (DyndnsClient, error) {
	if token == "" {
		return nil, MissingSecretError{}
	}
	c := &cloudflareClient{
		endpoint: CLOUDFLARE_API_ENDPOINT,
		token:    token,
		proxy:    proxy,
	}
	_, err := c.cloudflare(http.MethodGet, "zones", nil)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Check if DNS entries will be set to proxied
func (c *cloudflareClient) Proxy() bool {
	return c.proxy
}

func (c *cloudflareClient) Domains() []string {
	return c.domains
}

func (c *cloudflareClient) IPv4() string {
	return c.ipv4
}

func (c *cloudflareClient) IPv6() string {
	return c.ipv6
}

// Set domains to update, overrides previous domains
func (c *cloudflareClient) SetDomains(domains []string) {
	c.domains = domains
}

// Add a domain to the list of domains
func (c *cloudflareClient) AddDomain(domain string) {
	if c.domains == nil {
		c.domains = []string{domain}
	} else {
		c.domains = append(c.domains, domain)
	}
}

// Set IPv4 Address to use for creating entries
// Returns an error if string is neither empty nor a valid IP Address
func (c *cloudflareClient) SetIPv4(val string) error {
	if val != "" && !validIPv4(val) {
		return NewInvalidIPError(val)
	}
	c.ipv4 = val
	return nil
}

// Set IPv6 Address to use for creating entries
// Returns an error if string is neither empty nor a valid IP Address
func (c *cloudflareClient) SetIPv6(val string) error {
	if val != "" && !validIPv6(val) {
		return NewInvalidIPError(val)
	}
	c.ipv6 = val
	return nil
}

// Sent request to cloudflare api
func (c *cloudflareClient) cloudflare(method string, url string, body io.Reader) (cloudflareResponse, error) {
	url = c.endpoint + url
	slog.Debug("New request to cloudflare api", slog.String("url", url), slog.String("method", method))

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return cloudflareResponse{}, err
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set("Content-Type", "application/json")
	}

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return cloudflareResponse{}, err
	}
	if res.StatusCode != 200 {
		return cloudflareResponse{}, &HttpRequestFailedError{
			StatusCode: res.StatusCode,
			Body:       res.Body,
		}
	}

	var result cloudflareResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return cloudflareResponse{}, err
	}
	if !result.Successs {
		return cloudflareResponse{}, &CloudflareOperationFailedError{result: result}
	}
	return result, nil
}

// Retrieve the zone id of a given domain from cloudflare.
// Expects to receive at least 1 zone.
func (c *cloudflareClient) getZoneId(domain string) (string, error) {
	zone := getBaseDomain(domain)
	url := "zones?name=" + zone + "&status=active"

	slog.Info("Fetching zone id", slog.String("zone", zone))
	res, err := c.cloudflare(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	var zones []cloudflareZone
	err = json.Unmarshal(res.Result, &zones)
	if err != nil {
		return "", err
	} else if len(zones) < 1 {
		return "", NoDomainError{}
	}
	return zones[0].Id, nil
}

// Retrieve all records for a given zone.
func (c *cloudflareClient) getRecords(zone string, domain string) ([]cloudflareRecord, error) {
	url := "zones/" + zone + "/dns_records?name=" + domain

	slog.Info("Fetching records",
		slog.String("zone", zone),
		slog.String("domain", domain),
	)
	res, err := c.cloudflare(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var records []cloudflareRecord
	err = json.Unmarshal(res.Result, &records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

// Update the record in cloudflare.
// Sets the TTL to 1, which means automatic.
// Acceptable inputs for recordType are "A" and "AAAA".
func (c *cloudflareClient) updateRecord(zone string, domain string, recordType string, recordId string) error {
	url := "zones/" + zone + "/dns_records"
	method := http.MethodPost
	if recordId != "" {
		method = http.MethodPut
		url = url + "/" + recordId
	}
	var ip string
	if recordType == "A" {
		ip = c.IPv4()
	} else if recordType == "AAAA" {
		ip = c.IPv6()
	}

	record := cloudflareRecord{
		Content: ip,
		Name:    domain,
		Proxied: c.Proxy(),
		Type:    recordType,
		TTL:     1,
	}

	body, err := json.Marshal(record)
	if err != nil {
		return err
	}

	slog.Info("Updating record",
		slog.String("zone", zone),
		slog.String("domain", domain),
		slog.String("type", recordType),
		slog.String("recordId", recordId),
		slog.Bool("proxied", record.Proxied),
		slog.String("content", record.Content),
	)
	_, err = c.cloudflare(method, url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	return nil
}

// Update all domains, needs at least one of IPv4/IPv6 set.
// Will return with the first error
func (c *cloudflareClient) Update() error {
	if c.ipv4 == "" && c.ipv6 == "" {
		return NoIPError{}
	}
	if c.domains == nil || len(c.domains) == 0 {
		return NoDomainError{}
	}
	for _, domain := range c.domains {
		zone, err := c.getZoneId(domain)
		if err != nil {
			return err
		}

		records, err := c.getRecords(zone, domain)
		if err != nil {
			return err
		}

		// Iterate over all records and update the A and AAAA record if necessary
		var v4, v6 bool = false, false
		for _, record := range records {
			slog.Debug("Received record",
				slog.String("domain", domain),
				slog.String("type", record.Type),
				slog.String("content", record.Content),
				slog.String("modified_on", record.ModifiedOn),
			)
			switch record.Type {
			case "A":
				v4 = true
				if record.Content == c.IPv4() {
					continue
				}
			case "AAAA":
				v6 = true
				if record.Content == c.IPv6() {
					continue
				}
			default:
				continue
			}
			err = c.updateRecord(zone, domain, record.Type, record.Id)
			if err != nil {
				return err
			}
		}

		// Create A record if necessary
		if !v4 && c.IPv4() != "" {
			err = c.updateRecord(zone, domain, "A", "")
			if err != nil {
				return err
			}
		}
		// Create AAAA record if necessary
		if !v6 && c.IPv6() != "" {
			err = c.updateRecord(zone, domain, "AAAA", "")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Fetch the public IP(s) and run Update() periodically.
// Is executed as blocking and for forever.
// Will not continue run of the loop if an error occurs.
// Exits gracefully on SIGTERM
func (c *cloudflareClient) Run(interval time.Duration) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	for {
		ipv4, err := getPublicIPv4()
		if err != nil {
			slog.Error("Failed to get IPv4", "err", err)
		}
		ipv6, err := getPublicIPv6()
		if err != nil {
			slog.Error("Failed to get IPv6", "err", err)
		}

		changed := ipv4 != c.IPv4() || ipv6 != c.IPv6()
		if changed && ipv4 != c.IPv4() {
			err = c.SetIPv4(ipv4)
			if err != nil {
				slog.Error("Failed to get public IPv4", "err", err)
				changed = false
			}
		}
		if changed && ipv6 != c.IPv6() {
			err = c.SetIPv6(ipv6)
			if err != nil {
				slog.Error("Failed to get public IPv6", "err", err)
				changed = false
			}
		}
		if changed {
			slog.Info("Deteced changed IP",
				slog.String("ipv4", c.IPv4()),
				slog.String("ipv6", c.IPv6()),
			)
			err = c.Update()
			if err != nil {
				slog.Error("Failed to Update records", "err", err)
			}
		} else {
			slog.Debug("No changed detected")
		}
		var elapsedTime time.Duration = 0
		for elapsedTime < interval {
			timer := time.NewTimer(1 * time.Second)
			select {
			case <-timer.C:
				elapsedTime += time.Duration(1 * time.Second)
			case <-quit:
				timer.Stop()
				slog.Info("Received SIGTERM, shutting down client")
				return
			}
		}
	}
}
