package server

import (
	"testing"

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
