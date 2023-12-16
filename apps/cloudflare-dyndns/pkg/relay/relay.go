package relay

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/heathcliff26/containers/apps/cloudflare-dyndns/pkg/dyndns"
	"github.com/heathcliff26/containers/apps/cloudflare-dyndns/pkg/server"
)

type relay struct {
	token    string
	endpoint string
	data     *dyndns.ClientData
}

// Create a new Relay
func NewRelay(token string, proxy bool, endpoint string) (dyndns.Client, error) {
	if token == "" {
		return nil, dyndns.ErrMissingToken{}
	}
	if endpoint == "" {
		return nil, dyndns.ErrMissingEndpoint{}
	}
	return &relay{
		token:    token,
		endpoint: endpoint,
		data:     dyndns.NewClientData(proxy),
	}, nil
}

func (r *relay) Endpoint() string {
	return r.endpoint
}

// Give Access to ClientData
func (r *relay) Data() *dyndns.ClientData {
	return r.data
}

func (r *relay) Update() error {
	err := r.Data().CheckData()
	if err != nil {
		return err
	}
	params := server.RequestParams{
		Token:   r.token,
		Domains: r.Data().Domains(),
		IPv4:    r.Data().IPv4(),
		IPv6:    r.Data().IPv6(),
		Proxy:   r.Data().Proxy(),
	}
	body, err := json.Marshal(params)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, r.endpoint, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	c := http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return dyndns.NewErrHttpRequestFailed(res.StatusCode, res.Body)
	}

	var result server.Response
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return err
	}
	if !result.Success {
		return &dyndns.ErrOperationFailed{Result: res.Body}
	}
	return nil
}
