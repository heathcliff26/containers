package dyndns

import (
	"io"
	"net"
	"net/http"
	"strings"
)

// Validate if a string is a valid ip
func validIP(ip string) bool {
	return !(net.ParseIP(ip) == nil)
}

// Validate if a string is an IPv4
func ValidIPv4(ip string) bool {
	return validIP(ip) && strings.Contains(ip, ".")
}

// Validate if a string is an IPv6
func ValidIPv6(ip string) bool {
	return validIP(ip) && strings.Contains(ip, ":")
}

// Check if host has an IPv6 interface
func HasIPv6Support() bool {
	conn, err := net.Dial("udp", "[2606:4700:4700::1111]:80")
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// Use icanhazip.com to get the public ip
func getPublicIP(version string) (string, error) {
	res, err := http.Get("https://" + version + ".icanhazip.com")
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		return "", &ErrHttpRequestFailed{res.StatusCode, res.Body}
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	ip := strings.ReplaceAll(string(b), "\n", "")
	if !validIP(ip) {
		return "", &ErrInvalidIP{ip}
	}

	return ip, nil
}

// Equal to getPublicIP("ipv4")
func GetPublicIPv4() (string, error) {
	return getPublicIP("ipv4")
}

// Equal to getPublicIP("ipv6")
func GetPublicIPv6() (string, error) {
	if !HasIPv6Support() {
		return "", nil
	}
	return getPublicIP("ipv6")
}
