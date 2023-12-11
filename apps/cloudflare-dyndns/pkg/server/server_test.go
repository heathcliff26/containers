package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/heathcliff26/containers/apps/cloudflare-dyndns/pkg/client"
	"github.com/stretchr/testify/assert"
)

func TestVerifyAllowedDomains(t *testing.T) {
	tMatrix := []struct {
		Name               string
		Domains, Whitelist []string
		Ok                 bool
	}{
		{
			Name:      "OnlyAllowedDomains",
			Domains:   []string{"foo.example.org", "bar.example.net", "example.net"},
			Whitelist: []string{"example.org", "example.net"},
			Ok:        true,
		},
		{
			Name:      "WhitelistNil",
			Domains:   []string{"foo.example.org", "bar.example.net", "example.net"},
			Whitelist: nil,
			Ok:        true,
		},
		{
			Name:      "WhitelistEmpty",
			Domains:   []string{"foo.example.org", "bar.example.net", "example.net"},
			Whitelist: []string{},
			Ok:        true,
		},
		{
			Name:      "ForbiddenDomains",
			Domains:   []string{"foo.example.org", "bar.example.net", "example.net"},
			Whitelist: []string{"example.org"},
			Ok:        false,
		},
	}

	for _, tCase := range tMatrix {
		t.Run(tCase.Name, func(t *testing.T) {
			s := Server{Domains: tCase.Whitelist}
			assert.Equal(t, tCase.Ok, s.verifyAllowedDomains(tCase.Domains))
		})
	}
}

func TestRequestHandler(t *testing.T) {
	s := Server{
		Domains:      []string{"example.org"},
		createClient: client.NewTestClient,
	}

	tMatrix := []struct {
		Name, Method, Target string
		Request              RequestParams
		Status               int
		Response             Response
	}{
		{
			Name:   "WrongPath",
			Method: http.MethodGet,
			Target: "/index.html",
			Status: http.StatusNotFound,
		},
		{
			Name:   "WrongMethod",
			Method: http.MethodPut,
			Target: "/",
			Status: http.StatusMethodNotAllowed,
			Response: Response{
				Message: MESSAGE_WRONG_METHOD,
				Success: false,
			},
		},
		{
			Name:   "ParseErrorGet",
			Method: http.MethodGet,
			Target: "/?token=testtoken&domains=foo.example.net&ipv4=100.100.100.100&ipv6=fd00::dead&proxy=notabool",
			Status: http.StatusBadRequest,
			Response: Response{
				Message: MESSAGE_REQUEST_PARSING_FAILED,
				Success: false,
			},
		},
		{
			Name:   "ForbiddenDomain",
			Method: http.MethodPost,
			Target: "/",
			Request: RequestParams{
				Domains: []string{"foo.forbidden.org"},
			},
			Status: http.StatusForbidden,
			Response: Response{
				Message: MESSAGE_DOMAINS_FORBIDDEN,
				Success: false,
			},
		},
		{
			Name:   "Unauthorized",
			Method: http.MethodPost,
			Target: "/",
			Request: RequestParams{
				Token: "",
			},
			Status: http.StatusUnauthorized,
			Response: Response{
				Message: MESSAGE_UNAUTHORIZED,
				Success: false,
			},
		},
		{
			Name:   "InvalidIPv4",
			Method: http.MethodPost,
			Target: "/",
			Request: RequestParams{
				Token: "testtoken",
				IPv4:  "Not an IPv4",
			},
			Status: http.StatusBadRequest,
			Response: Response{
				Message: MESSAGE_INVALID_IP,
				Success: false,
			},
		},
		{
			Name:   "InvalidIPv6",
			Method: http.MethodPost,
			Target: "/",
			Request: RequestParams{
				Token: "testtoken",
				IPv6:  "Not an IPv6",
			},
			Status: http.StatusBadRequest,
			Response: Response{
				Message: MESSAGE_INVALID_IP,
				Success: false,
			},
		},
		{
			Name:   "MissingIP",
			Method: http.MethodPost,
			Target: "/",
			Request: RequestParams{
				Token: "testtoken",
			},
			Status: http.StatusBadRequest,
			Response: Response{
				Message: MESSAGE_INVALID_IP,
				Success: false,
			},
		},
		{
			Name:   "MissingDomains",
			Method: http.MethodPost,
			Target: "/",
			Request: RequestParams{
				Token: "testtoken",
				IPv4:  "100.100.100.100",
				IPv6:  "fd00::dead",
			},
			Status: http.StatusBadRequest,
			Response: Response{
				Message: MESSAGE_MISSING_DOMAIN,
				Success: false,
			},
		},
		{
			Name:   "Success",
			Method: http.MethodPost,
			Target: "/",
			Request: RequestParams{
				Token:   "testtoken",
				IPv4:    "100.100.100.100",
				IPv6:    "fd00::dead",
				Domains: []string{"foo.example.org"},
			},
			Status: http.StatusOK,
			Response: Response{
				Message: MESSAGE_SUCCESS,
				Success: true,
			},
		},
	}

	for _, tCase := range tMatrix {
		t.Run(tCase.Name, func(t *testing.T) {
			var buf bytes.Buffer
			if tCase.Method == http.MethodPost {
				err := json.NewEncoder(&buf).Encode(tCase.Request)
				if err != nil {
					t.Fatalf("Unexpected error creating request body: %v", err)
				}
			}

			req := httptest.NewRequest(tCase.Method, tCase.Target, &buf)
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			s.requestHandler(rr, req)

			assert := assert.New(t)

			res := rr.Result()
			assert.Equal(tCase.Status, res.StatusCode)
			if res.StatusCode == http.StatusNotFound {
				// There is no body in this case, end test here
				return
			}

			var response Response
			err := json.NewDecoder(res.Body).Decode(&response)
			if !assert.Nil(err) {
				t.Fatalf("Failed to decode response: %v", err)
			}
			assert.Equal(tCase.Response, response)
		})
	}

	t.Run("WrongContentType", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		rr := httptest.NewRecorder()

		s.requestHandler(rr, req)

		assert := assert.New(t)

		res := rr.Result()
		assert.Equal(http.StatusUnsupportedMediaType, res.StatusCode)

		var response Response
		err := json.NewDecoder(res.Body).Decode(&response)
		if !assert.Nil(err) {
			t.Fatalf("Failed to decode response: %v", err)
		}
		assert.Equal(Response{MESSAGE_WRONG_CONTENT_TYPE, false}, response)
	})

	t.Run("ParseErrorPost", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		s.requestHandler(rr, req)

		assert := assert.New(t)

		res := rr.Result()
		assert.Equal(http.StatusBadRequest, res.StatusCode)

		var response Response
		err := json.NewDecoder(res.Body).Decode(&response)
		if !assert.Nil(err) {
			t.Fatalf("Failed to decode response: %v", err)
		}
		assert.Equal(Response{MESSAGE_REQUEST_PARSING_FAILED, false}, response)
	})
}
