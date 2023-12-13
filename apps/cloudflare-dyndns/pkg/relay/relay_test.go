package relay

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/heathcliff26/containers/apps/cloudflare-dyndns/pkg/dyndns"
	"github.com/heathcliff26/containers/apps/cloudflare-dyndns/pkg/server"
	"github.com/stretchr/testify/assert"
)

func TestNewRelayClient(t *testing.T) {
	c, err := NewRelay("", true, "")

	assert := assert.New(t)

	assert.Equal(dyndns.ErrMissingSecret{}, err)
	assert.Nil(c)

	c, err = NewRelay("testtoken", true, "")

	assert.Equal(dyndns.ErrMissingEndpoint{}, err)
	assert.Nil(c)
}

func TestRelayGetEndpoint(t *testing.T) {
	r := &relay{
		endpoint: "dyndns.example.org",
	}
	assert.Equal(t, "dyndns.example.org", r.Endpoint())
}

func TestRelayUpdate(t *testing.T) {
	assert := assert.New(t)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(http.MethodPost, req.Method)
		assert.Equal("application/json", req.Header.Get("Content-Type"))

		expectedParams := server.RequestParams{
			Token:   "testtoken",
			Domains: []string{"foo.example.org"},
			IPv4:    "100.100.100.100",
			IPv6:    "fd00::dead",
			Proxy:   true,
		}

		var params server.RequestParams
		err := json.NewDecoder(req.Body).Decode(&params)
		assert.Nil(err)
		assert.Equal(expectedParams, params)

		res := server.Response{Message: "test", Success: true}
		b, err := json.Marshal(res)
		if err != nil {
			t.Fatalf("Could not convert server.Response to json body, err: %v", err)
		}

		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write(b)
	}))
	defer server.Close()

	d := dyndns.NewClientData(true)
	d.SetDomains([]string{"foo.example.org"})
	err1 := d.SetIPv4("100.100.100.100")
	err2 := d.SetIPv6("fd00::dead")
	if err1 != nil && err2 != nil {
		t.Fatalf("Unexpected error setting IPs: err1=\"%v\", err2=\"%v\"", err1, err2)
	}
	r := &relay{
		token:    "testtoken",
		endpoint: server.URL + "/",
		data:     d,
	}
	err := r.Update()
	assert.Nil(err)
}
