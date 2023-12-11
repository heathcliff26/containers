package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Response struct {
	Message string `json:"msg"`
	Success bool   `json:"success"`
}

// Send a response to the writer and handle impossible parse errors
func sendResponse(rw http.ResponseWriter, msg string, success bool) {
	res := Response{msg, success}

	b, err := json.Marshal(res)
	if err != nil {
		if success {
			rw.WriteHeader(http.StatusInternalServerError)
		}
		slog.Error("Failed to create Response", "err", err)
		return
	}

	_, err = rw.Write(b)
	if err != nil {
		slog.Error("Failed to send response to client", "err", err)
	}
}

// Parse the url parameters into the struct RequestParams
func parseURLParams(url *url.URL, result *RequestParams) error {
	params := url.Query()
	if params.Has("token") {
		result.Token = params.Get("token")
	} else {
		// Compatibility for https://github.com/1rfsNet/Fritz-Box-Cloudflare-DynDNS
		result.Token = params.Get("cf_key")
	}
	var s string
	if params.Has("domains") {
		s = params.Get("domains")
	} else {
		// Compatibility for https://github.com/1rfsNet/Fritz-Box-Cloudflare-DynDNS
		s = params.Get("domain")
	}
	result.Domains = strings.Split(s, ",")
	if params.Has("ipv4") {
		result.IPv4 = params.Get("ipv4")
	}
	if params.Has("ipv6") {
		result.IPv6 = params.Get("ipv6")
	}
	if params.Has("proxy") {
		s = params.Get("proxy")
		var err error
		result.Proxy, err = strconv.ParseBool(s)
		if err != nil {
			return err
		}
	}
	return nil
}
