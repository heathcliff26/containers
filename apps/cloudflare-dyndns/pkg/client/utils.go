package client

import "strings"

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
