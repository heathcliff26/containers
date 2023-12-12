package dyndns

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		assert.Equal(t, tCase.Ok, ValidIPv4(tCase.IP))
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
		assert.Equal(t, tCase.Ok, ValidIPv6(tCase.IP))
	}
}

func TestGetPublicIP(t *testing.T) {
	tMatrix := []struct {
		Name string
		f    func() (string, error)
	}{
		{"IPv4", GetPublicIPv4},
		{"IPv6", GetPublicIPv6},
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
