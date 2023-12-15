package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/heathcliff26/containers/apps/cloudflare-dyndns/pkg/client"
	"github.com/heathcliff26/containers/apps/cloudflare-dyndns/pkg/config"
	"github.com/heathcliff26/containers/apps/cloudflare-dyndns/pkg/dyndns"
)

type Server struct {
	Addr         string
	Domains      map[string]bool
	SSL          config.SSLConfig
	createClient func(string, bool) (dyndns.Client, error)
}

type RequestParams struct {
	// To maintain compatibility with https://github.com/1rfsNet/Fritz-Box-Cloudflare-DynDNS, the cf_key and domain can be used as well for token and domains when doing GET
	Token   string   `json:"token"`           // Token needed for accessing cloudflare api.
	Domains []string `json:"domains"`         // The domain to update, parsed from comma (,) separated string, needs at least 1.
	IPv4    string   `json:"ipv4,omitempty"`  // IPv4 Address, optional, when IPv6 set
	IPv6    string   `json:"ipv6,omitempty"`  // IPv6 Address, optional, when IPv4 set
	Proxy   bool     `json:"proxy,omitempty"` // Indicate if domain should be proxied, defaults to true
}

const (
	MESSAGE_WRONG_METHOD           = "Wrong Method, expected GET or POST"
	MESSAGE_WRONG_CONTENT_TYPE     = "Wrong Content-Type, expected application/json"
	MESSAGE_REQUEST_PARSING_FAILED = "Failed to parse the request, received wrong or malformed parameters"
	MESSAGE_UNAUTHORIZED           = "Failed to authenticate to cloudflare"
	MESSAGE_DOMAINS_FORBIDDEN      = "At least one of the provided domains is not allowed to be handled by this server"
	MESSAGE_INVALID_IP             = "Either no IP or an invalid IP has been provided"
	MESSAGE_MISSING_DOMAIN         = "No domains have been provided"
	MESSAGE_FAILED_UPDATE          = "Failed to update the records"
	MESSAGE_SUCCESS                = "Updated dyndns records"
)

// Return a new Server, created from the provided config
func NewServer(c config.ServerConfig) *Server {
	return &Server{
		Addr:         ":" + strconv.Itoa(c.Port),
		Domains:      newDomainMap(c.Domains),
		SSL:          c.SSL,
		createClient: client.NewCloudflareClient,
	}
}

func newDomainMap(domains []string) map[string]bool {
	m := make(map[string]bool, len(domains))
	for _, d := range domains {
		m[d] = true
	}
	return m
}

// Ensures that all domains listed here are
func (s *Server) verifyAllowedDomains(domains []string) bool {
	if len(s.Domains) == 0 {
		return true
	}

	for _, domain := range domains {
		forbidden := true
		d := strings.Split(domain, ".")
		domain = ""
		for i := len(d) - 1; i >= 0; i-- {
			if domain == "" {
				domain = d[i]
			} else {
				domain = d[i] + "." + domain
			}
			if s.Domains[domain] {
				forbidden = false
				break
			}
		}
		if forbidden {
			return false
		}
	}
	return true
}

// Main function of the server, used to server requests from clients
func (s *Server) requestHandler(rw http.ResponseWriter, req *http.Request) {
	if strings.Split(req.URL.String(), "?")[0] != "/" {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	// Verify right method
	if req.Method != http.MethodGet && req.Method != http.MethodPost {
		slog.Debug("Received request with wrong method type", "method", req.Method)
		rw.WriteHeader(http.StatusMethodNotAllowed)
		sendResponse(rw, MESSAGE_WRONG_METHOD, false)
		return
	}
	// Verify right Content-Type
	if req.Method == http.MethodPost && req.Header.Get("Content-Type") != "application/json" {
		slog.Debug("Received request with wrong Content-Type", "Content-Type", req.Header.Get("Content-Type"))
		rw.WriteHeader(http.StatusUnsupportedMediaType)
		sendResponse(rw, MESSAGE_WRONG_CONTENT_TYPE, false)
		return
	}

	// Parse parameters
	params := RequestParams{Proxy: true}
	if req.Method == http.MethodPost {
		err := json.NewDecoder(req.Body).Decode(&params)
		if err != nil {
			slog.Info("Failed to parse request from json", "err", err)
			rw.WriteHeader(http.StatusBadRequest)
			sendResponse(rw, MESSAGE_REQUEST_PARSING_FAILED, false)
			return
		}
	} else {
		err := parseURLParams(req.URL, &params)
		if err != nil {
			slog.Info("Failed to parse request from url parameters", "err", err)
			rw.WriteHeader(http.StatusBadRequest)
			sendResponse(rw, MESSAGE_REQUEST_PARSING_FAILED, false)
			return
		}
	}

	if !s.verifyAllowedDomains(params.Domains) {
		slog.Info("Request contained domain(s) not on the whitelist")
		rw.WriteHeader(http.StatusForbidden)
		sendResponse(rw, MESSAGE_DOMAINS_FORBIDDEN, false)
		return
	}

	// Create client
	c, err := s.createClient(params.Token, params.Proxy)
	if err != nil {
		slog.Info("Failed to create client", "err", err)
		rw.WriteHeader(http.StatusUnauthorized)
		sendResponse(rw, MESSAGE_UNAUTHORIZED, false)
		return
	}
	err = c.Data().SetIPv4(params.IPv4)
	if err != nil {
		slog.Info("Failed to set IPv4 for client", "err", err)
		rw.WriteHeader(http.StatusBadRequest)
		sendResponse(rw, MESSAGE_INVALID_IP, false)
		return
	}
	err = c.Data().SetIPv6(params.IPv6)
	if err != nil {
		slog.Info("Failed to set IPv6 for client", "err", err)
		rw.WriteHeader(http.StatusBadRequest)
		sendResponse(rw, MESSAGE_INVALID_IP, false)
		return
	}
	c.Data().SetDomains(params.Domains)

	// Update records
	err = c.Update()
	if err != nil {
		switch err.(type) {
		case dyndns.ErrNoIP:
			slog.Info("Received request with no IP")
			rw.WriteHeader(http.StatusBadRequest)
			sendResponse(rw, MESSAGE_INVALID_IP, false)
			return
		case dyndns.ErrNoDomain:
			slog.Info("Received request with no domains")
			rw.WriteHeader(http.StatusBadRequest)
			sendResponse(rw, MESSAGE_MISSING_DOMAIN, false)
			return
		default:
			slog.Info("Failed to update records", "err", err, slog.Group("req",
				slog.String("domains", fmt.Sprintf("%v", params.Domains)),
				slog.String("ipv4", params.IPv4),
				slog.String("ipv6", params.IPv6),
				slog.Bool("proxy", params.Proxy),
			))
			rw.WriteHeader(http.StatusOK)
			sendResponse(rw, MESSAGE_FAILED_UPDATE, false)
			return
		}
	}
	slog.Info("Successfully update record with the following request", slog.Group("req",
		slog.String("domains", fmt.Sprintf("%v", params.Domains)),
		slog.String("ipv4", params.IPv4),
		slog.String("ipv6", params.IPv6),
		slog.Bool("proxy", params.Proxy),
	))
	sendResponse(rw, MESSAGE_SUCCESS, true)
}

// Starts the server and exits with error if that fails
func (s *Server) Run() error {
	http.HandleFunc("/", s.requestHandler)

	var err error
	if s.SSL.Enabled {
		err = http.ListenAndServeTLS(s.Addr, s.SSL.Cert, s.SSL.Key, nil)
	} else {
		err = http.ListenAndServe(s.Addr, nil)
	}
	// This just means the server was closed after running
	if errors.Is(err, http.ErrServerClosed) {
		slog.Info("Server closed, exiting")
		return nil
	}
	return err
}
