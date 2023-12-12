package client

import "encoding/json"

type cloudflareResponse struct {
	Errors   []cloudflareMessages `json:"errors"`
	Messages []cloudflareMessages `json:"messages"`
	Success  bool                 `json:"success"`
	Result   json.RawMessage      `json:"result"`
}

type cloudflareMessages struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type cloudflareZone struct {
	Id string `json:"id"`
}

// Also used in requests, so "omitempty" needs to be set, otherwise requests fail
type cloudflareRecord struct {
	Content    string `json:"content,omitempty"`
	Name       string `json:"name,omitempty"`
	Proxied    bool   `json:"proxied,omitempty"`
	Type       string `json:"type,omitempty"`
	Comment    string `json:"comment,omitempty"`
	TTL        int    `json:"ttl,omitempty"`
	Id         string `json:"id,omitempty"`
	ModifiedOn string `json:"modified_on,omitempty"`
}
