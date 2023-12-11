package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendResponse(t *testing.T) {
	tMatrix := []struct {
		Msg     string
		Success bool
	}{
		{"Successfull test", true},
		{"Failed test", false},
	}

	for _, tCase := range tMatrix {
		rr := httptest.NewRecorder()
		sendResponse(rr, tCase.Msg, tCase.Success)

		assert := assert.New(t)

		assert.Equal(http.StatusOK, rr.Code)
		var res Response
		err := json.NewDecoder(rr.Body).Decode(&res)
		if !assert.Nil(err) {
			t.Fatalf("Failed to decode body: %v", err)
		}
		assert.Equal(tCase.Success, res.Success)
		assert.Equal(tCase.Msg, res.Message)
	}
}

func TestParseURLParams(t *testing.T) {
	tMatrix := []struct {
		Name   string
		URL    string
		Result RequestParams
	}{
		{
			Name: "SingleDomain",
			URL:  "http://example.org/?token=testtoken&domains=foo.example.net&ipv4=100.100.100.100&ipv6=fd00::dead&proxy=true",
			Result: RequestParams{
				Token:   "testtoken",
				Domains: []string{"foo.example.net"},
				IPv4:    "100.100.100.100",
				IPv6:    "fd00::dead",
				Proxy:   true,
			},
		},
		{
			Name: "MultipleDomains",
			URL:  "http://example.org/?token=testtoken&domains=foo.example.net,bar.example.org,example.net&ipv4=100.100.100.100&ipv6=fd00::dead&proxy=true",
			Result: RequestParams{
				Token:   "testtoken",
				Domains: []string{"foo.example.net", "bar.example.org", "example.net"},
				IPv4:    "100.100.100.100",
				IPv6:    "fd00::dead",
				Proxy:   true,
			},
		},
		{
			Name: "WithoutProxyParam",
			URL:  "http://example.org/?token=testtoken&domains=foo.example.net&ipv4=100.100.100.100&ipv6=fd00::dead",
			Result: RequestParams{
				Token:   "testtoken",
				Domains: []string{"foo.example.net"},
				IPv4:    "100.100.100.100",
				IPv6:    "fd00::dead",
				Proxy:   true,
			},
		},
		{
			Name: "NoProxy",
			URL:  "http://example.org/?token=testtoken&domains=foo.example.net&ipv4=100.100.100.100&ipv6=fd00::dead&proxy=false",
			Result: RequestParams{
				Token:   "testtoken",
				Domains: []string{"foo.example.net"},
				IPv4:    "100.100.100.100",
				IPv6:    "fd00::dead",
				Proxy:   false,
			},
		},
		{
			Name: "IPv4Only",
			URL:  "http://example.org/?token=testtoken&domains=foo.example.net&ipv4=100.100.100.100&proxy=true",
			Result: RequestParams{
				Token:   "testtoken",
				Domains: []string{"foo.example.net"},
				IPv4:    "100.100.100.100",
				IPv6:    "",
				Proxy:   true,
			},
		},
		{
			Name: "IPv6Only",
			URL:  "http://example.org/?token=testtoken&domains=foo.example.net&ipv6=fd00::dead&proxy=true",
			Result: RequestParams{
				Token:   "testtoken",
				Domains: []string{"foo.example.net"},
				IPv4:    "",
				IPv6:    "fd00::dead",
				Proxy:   true,
			},
		},
		{
			Name: "LegacyParams",
			URL:  "http://example.org/?cf_key=testtoken&domain=foo.example.net&ipv4=100.100.100.100&ipv6=fd00::dead&proxy=true",
			Result: RequestParams{
				Token:   "testtoken",
				Domains: []string{"foo.example.net"},
				IPv4:    "100.100.100.100",
				IPv6:    "fd00::dead",
				Proxy:   true,
			},
		},
	}

	for _, tCase := range tMatrix {
		t.Run(tCase.Name, func(t *testing.T) {
			res := RequestParams{Proxy: true}
			url, err := url.Parse(tCase.URL)
			if err != nil {
				t.Fatalf("Something went wrong when creating the url: %v", err)
			}
			err = parseURLParams(url, &res)
			assert.Nil(t, err)
			assert.Equal(t, tCase.Result, res)
		})
	}
}
