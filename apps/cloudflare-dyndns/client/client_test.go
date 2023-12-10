package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Only test missing secret is checked, login needs to be tested separately since
func TestNewCloudflareClient(t *testing.T) {
	c, err := NewCloudflareClient("", true)

	assert := assert.New(t)

	assert.Equal(MissingSecretError{}, err)
	assert.Nil(c)
}

func TestClouflareAuthentication(t *testing.T) {
	assert := assert.New(t)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(http.MethodGet, req.Method)
		assert.Equal("/zones", req.URL.String())

		auth := req.Header.Get("Authorization")
		res := cloudflareResponse{}

		if res.Successs = assert.Equal("Bearer testtoken", auth); res.Successs {
			rw.WriteHeader(http.StatusOK)
		} else {
			rw.WriteHeader(http.StatusUnauthorized)
		}
		b, err := json.Marshal(res)
		if err != nil {
			t.Fatalf("Could not convert cloudflareResponse to json body, err: %v", err)
		}

		_, _ = rw.Write(b)
	}))
	defer server.Close()

	c := &cloudflareClient{
		endpoint: server.URL + "/",
		token:    "testtoken",
	}

	_, err := c.cloudflare(http.MethodGet, "zones", nil)
	assert.Nil(err)
}

func TestCloudflareClientGetFunctions(t *testing.T) {
	domains := []string{"foo.example.org", "bar.example.org"}
	c := &cloudflareClient{
		proxy:   true,
		domains: domains,
		ipv4:    "100.100.100.100",
		ipv6:    "fd00::dead",
	}
	t.Run("Proxy", func(t *testing.T) {
		assert.Equal(t, true, c.Proxy())
	})
	t.Run("Domains", func(t *testing.T) {
		assert.Equal(t, domains, c.Domains())
	})
	t.Run("IPv4", func(t *testing.T) {
		assert.Equal(t, "100.100.100.100", c.IPv4())
	})
	t.Run("IPv6", func(t *testing.T) {
		assert.Equal(t, "fd00::dead", c.IPv6())
	})
}

func TestCloudflareClientSetFunctions(t *testing.T) {
	t.Run("SetDomains", func(t *testing.T) {
		c := &cloudflareClient{}
		domains := []string{"foo.example.org", "bar.example.org"}
		c.SetDomains(domains)
		assert.Equal(t, domains, c.Domains())
	})
	t.Run("AddDomain", func(t *testing.T) {
		c := &cloudflareClient{}
		c.AddDomain("foo.example.org")
		assert.Equal(t, []string{"foo.example.org"}, c.Domains())
		c.AddDomain("bar.example.org")
		assert.Equal(t, []string{"foo.example.org", "bar.example.org"}, c.Domains())
	})
	t.Run("SetIPv4", func(t *testing.T) {
		c := &cloudflareClient{}
		err := c.SetIPv4("100.100.100.100")
		assert.Equal(t, "100.100.100.100", c.IPv4())
		assert.Nil(t, err)
	})
	t.Run("SetIPv6", func(t *testing.T) {
		c := &cloudflareClient{}
		err := c.SetIPv6("fd00::dead")
		assert.Equal(t, "fd00::dead", c.IPv6())
		assert.Nil(t, err)
	})
}

func TestGetZoneId(t *testing.T) {
	assert := assert.New(t)

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(http.MethodGet, req.Method)
		assert.Equal("/zones?name=example.org&status=active", req.URL.String())
		assert.Equal("Bearer testtoken", req.Header.Get("Authorization"))

		res := cloudflareResponse{Successs: true}

		result := []cloudflareZone{{Id: "44a6dc905d4ff61b"}}
		b, err := json.Marshal(result)
		if err != nil {
			t.Fatalf("Could not convert []cloudflareZone to json, err: %v", err)
		}
		res.Result = b

		b, err = json.Marshal(res)
		if err != nil {
			t.Fatalf("Could not convert cloudflareResponse to json body, err: %v", err)
		}

		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write(b)
	}))
	defer server.Close()

	c := &cloudflareClient{
		endpoint: server.URL + "/",
		token:    "testtoken",
	}

	res, err := c.getZoneId("example.org")
	if !assert.Nil(err) {
		t.Fatalf("Failed to get zone id: %v", err)
	}
	assert.Equal("44a6dc905d4ff61b", res)
}

func TestGetRecords(t *testing.T) {
	assert := assert.New(t)

	records := []cloudflareRecord{
		{
			Content: "100.100.100.100",
			Id:      "21d167bb587e1d3e",
			Type:    "A",
		},
		{
			Content: "fd00::dead",
			Id:      "ff0012854eddab59",
			Type:    "AAAA",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(http.MethodGet, req.Method)
		assert.Equal("/zones/6384bd8687814061/dns_records?name=foo.example.org", req.URL.String())
		assert.Equal("Bearer testtoken", req.Header.Get("Authorization"))

		res := cloudflareResponse{Successs: true}

		b, err := json.Marshal(records)
		if err != nil {
			t.Fatalf("Could not convert []cloudflareRecords to json, err: %v", err)
		}
		res.Result = b

		b, err = json.Marshal(res)
		if err != nil {
			t.Fatalf("Could not convert cloudflareResponse to json body, err: %v", err)
		}

		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write(b)
	}))
	defer server.Close()

	c := &cloudflareClient{
		endpoint: server.URL + "/",
		token:    "testtoken",
	}

	res, err := c.getRecords("6384bd8687814061", "foo.example.org")
	if !assert.Nil(err) {
		t.Fatalf("Failed to get records: %v", err)
	}
	assert.Equal(records, res)
}

func TestUpdateRecord(t *testing.T) {
	zone, domain := "78fc43dc6a8c5e7c", "bar.example.org"
	tMatrix := []struct {
		Name   string
		Proxy  bool
		Record cloudflareRecord
	}{
		{
			Name: "UpdateA",
			Record: cloudflareRecord{
				Content: "100.100.100.100",
				Id:      "e1cfccf8b4f40a27",
				Type:    "A",
			},
		},
		{
			Name: "UpdateAAAA",
			Record: cloudflareRecord{
				Content: "fd00::dead",
				Id:      "d39c32e77ba9c477",
				Type:    "AAAA",
			},
		},
		{
			Name: "CreateA",
			Record: cloudflareRecord{
				Content: "10.8.100.100",
				Id:      "6dbf0f498e60487f",
				Type:    "A",
			},
		},
		{
			Name: "CreateAAAA",
			Record: cloudflareRecord{
				Content: "fd69::dead",
				Id:      "ce8a2c45433edf26",
				Type:    "AAAA",
			},
		},
	}

	for _, tCase := range tMatrix {
		t.Run(tCase.Name, func(t *testing.T) {
			assert := assert.New(t)

			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				assert.Equal("Bearer testtoken", req.Header.Get("Authorization"))
				assert.Equal("application/json", req.Header.Get("Content-Type"))

				var record cloudflareRecord
				err := json.NewDecoder(req.Body).Decode(&record)
				if !assert.Nil(err) {
					t.Fatalf("Could not convert request to cloudflareRecord: %v", err)
				}
				assert.Equal(tCase.Record.Content, record.Content)
				assert.Equal(domain, record.Name)
				assert.Equal(tCase.Proxy, record.Proxied)
				assert.Equal(tCase.Record.Type, record.Type)
				assert.Equal(1, record.TTL)
				if tCase.Record.Id != "" {
					assert.Equal(http.MethodPut, req.Method)
					assert.Equal("/zones/"+zone+"/dns_records/"+tCase.Record.Id, req.URL.String())
				} else {
					assert.Equal(http.MethodPost, req.Method)
					assert.Equal("/zones/"+zone+"/dns_records", req.URL.String())
				}

				res := cloudflareResponse{Successs: true}

				b, err := json.Marshal(res)
				if err != nil {
					t.Fatalf("Could not convert cloudflareResponse to json body, err: %v", err)
				}

				rw.WriteHeader(http.StatusOK)
				_, _ = rw.Write(b)
			}))
			defer server.Close()

			c := &cloudflareClient{
				endpoint: server.URL + "/",
				token:    "testtoken",
				proxy:    tCase.Proxy,
			}
			if tCase.Record.Type == "A" {
				err := c.SetIPv4(tCase.Record.Content)
				assert.Nil(err)
			} else {
				err := c.SetIPv6(tCase.Record.Content)
				assert.Nil(err)
			}
			err := c.updateRecord(zone, domain, tCase.Record.Type, tCase.Record.Id)
			if !assert.Nil(err) {
				t.Fatalf("Failed to update record: %v", err)
			}
		})
	}
}
