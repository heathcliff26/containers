package client

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBaseDomain(t *testing.T) {
	tMatrix := []struct {
		input, output string
	}{
		{"", ""},
		{"not a domain", ""},
		{"foo.example.org", "example.org"},
		{"bar.example.org", "example.org"},
		{"example.net", "example.net"},
	}

	assert := assert.New(t)

	for _, tCase := range tMatrix {
		r := getBaseDomain(tCase.input)
		assert.Equal(tCase.output, r, "Expected %s to result in %s", tCase.input, tCase.output)
	}
}

func TestValidIPv4(t *testing.T) {
	tMatrix := []struct {
		IP string
		Ok bool
	}{
		{"", false},
		{"not an ip", false},
		{"100.100.100.100", true},
		{"172.198.10.100", true},
		{"fd00::dead", false},
	}

	for _, tCase := range tMatrix {
		assert.Equal(t, tCase.Ok, validIPv4(tCase.IP))
	}
}

func TestValidIPv6(t *testing.T) {
	tMatrix := []struct {
		IP string
		Ok bool
	}{
		{"", false},
		{"not an ip", false},
		{"100.100.100.100", false},
		{"fd69::dead", true},
		{"fd00::dead", true},
	}

	for _, tCase := range tMatrix {
		assert.Equal(t, tCase.Ok, validIPv6(tCase.IP))
	}
}

func TestGetPublicIP(t *testing.T) {
	tMatrix := []struct {
		Name string
		f    func() (string, error)
	}{
		{"IPv4", getPublicIPv4},
		{"IPv6", getPublicIPv6},
	}

	for _, tCase := range tMatrix {
		t.Run(tCase.Name, func(t *testing.T) {
			ip, err := tCase.f()

			assert := assert.New(t)

			if err != nil {
				if tCase.Name == "IPv6" && strings.Contains(err.Error(), "Get \"https://ipv6.icanhazip.com\": dial tcp") {
					t.Skip("No IPv6 connectivity")
				}
				t.Fatalf("Received unexpedted error: %s", err.Error())
			}
			assert.NotEmpty(ip)
		})
	}
}
