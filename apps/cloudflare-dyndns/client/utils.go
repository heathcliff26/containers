package client

import (
	"io"
	"net"
	"net/http"
	"strings"
)

// Example: foo.example.org -> example.org
func getBaseDomain(domain string) string {
	s := strings.Split(domain, ".")
	l := len(s)
	// Should be at least 2 entries
	if l < 2 {
		return ""
	}
	return strings.Join(s[l-2:], ".")
}

// Validate if a string is a valid ip
func validIP(ip string) bool {
	return !(net.ParseIP(ip) == nil)
}

// Validate if a string is an IPv4
func validIPv4(ip string) bool {
	return validIP(ip) && strings.Contains(ip, ".")
}

// Validate if a string is an IPv6
func validIPv6(ip string) bool {
	return validIP(ip) && strings.Contains(ip, ":")
}

// Use icanhazip.com to get the public ip
func getPublicIP(version string) (string, error) {
	res, err := http.Get("https://" + version + ".icanhazip.com")
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		return "", &HttpRequestFailedError{
			StatusCode: res.StatusCode,
			Body:       res.Body,
		}
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	ip := strings.ReplaceAll(string(b), "\n", "")
	if !validIP(ip) {
		return "", NewInvalidIPError(ip)
	}

	return ip, nil
}

// Equal to getPublicIP("ipv4")
func getPublicIPv4() (string, error) {
	return getPublicIP("ipv4")
}

// Equal to getPublicIP("ipv6")
func getPublicIPv6() (string, error) {
	return getPublicIP("ipv6")
}
