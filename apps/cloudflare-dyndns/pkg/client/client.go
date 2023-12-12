package client

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/heathcliff26/containers/apps/cloudflare-dyndns/pkg/dyndns"
)

const CLOUDFLARE_API_ENDPOINT = "https://api.cloudflare.com/client/v4/"

// Implements dyndns.Client
type cloudflareClient struct {
	// API is only here as a variable to enable local testing without relying on the cloudflare API
	endpoint string
	token    string
	data     *dyndns.ClientData
}

// Create a new CloudflareClient and test if the token is valid
func NewCloudflareClient(token string, proxy bool) (dyndns.Client, error) {
	if token == "" {
		return nil, dyndns.ErrMissingSecret{}
	}

	c := &cloudflareClient{
		endpoint: CLOUDFLARE_API_ENDPOINT,
		token:    token,
		data:     dyndns.NewClientData(proxy),
	}
	_, err := c.cloudflare(http.MethodGet, "zones", nil)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Give Access to ClientData
func (c *cloudflareClient) Data() *dyndns.ClientData {
	return c.data
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
		return cloudflareResponse{}, &dyndns.ErrHttpRequestFailed{StatusCode: res.StatusCode, Body: res.Body}
	}

	var result cloudflareResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return cloudflareResponse{}, err
	}
	if !result.Success {
		return cloudflareResponse{}, &dyndns.ErrOperationFailed{Result: res.Body}
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
		return "", dyndns.ErrNoDomain{}
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
		ip = c.Data().IPv4()
	} else if recordType == "AAAA" {
		ip = c.Data().IPv6()
	}

	record := cloudflareRecord{
		Content: ip,
		Name:    domain,
		Proxied: c.Data().Proxy(),
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
	if c.Data().IPv4() == "" && c.Data().IPv6() == "" {
		return dyndns.ErrNoIP{}
	}
	if c.Data().Domains() == nil || len(c.Data().Domains()) == 0 {
		return dyndns.ErrNoDomain{}
	}
	for _, domain := range c.Data().Domains() {
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
				if record.Content == c.Data().IPv4() {
					continue
				}
			case "AAAA":
				v6 = true
				if record.Content == c.Data().IPv6() {
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
		if !v4 && c.Data().IPv4() != "" {
			err = c.updateRecord(zone, domain, "A", "")
			if err != nil {
				return err
			}
		}
		// Create AAAA record if necessary
		if !v6 && c.Data().IPv6() != "" {
			err = c.updateRecord(zone, domain, "AAAA", "")
			if err != nil {
				return err
			}
		}
	}
	return nil
}
