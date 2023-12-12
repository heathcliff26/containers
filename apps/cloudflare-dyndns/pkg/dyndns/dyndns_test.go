package dyndns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelayGetFunctions(t *testing.T) {
	domains := []string{"foo.example.org", "bad.example.org"}
	d := &ClientData{
		proxy:   true,
		domains: domains,
		ipv4:    "100.100.100.100",
		ipv6:    "fd00::dead",
	}
	t.Run("Proxy", func(t *testing.T) {
		assert.Equal(t, true, d.Proxy())
	})
	t.Run("Domains", func(t *testing.T) {
		assert.Equal(t, domains, d.Domains())
	})
	t.Run("IPv4", func(t *testing.T) {
		assert.Equal(t, "100.100.100.100", d.IPv4())
	})
	t.Run("IPv6", func(t *testing.T) {
		assert.Equal(t, "fd00::dead", d.IPv6())
	})
}

func TestRelaySetFunctions(t *testing.T) {
	t.Run("SetProxy", func(t *testing.T) {
		d := &ClientData{}
		proxy := true
		d.SetProxy(proxy)
		assert.Equal(t, proxy, d.Proxy())
	})
	t.Run("SetDomains", func(t *testing.T) {
		d := &ClientData{}
		domains := []string{"foo.example.org", "bad.example.org"}
		d.SetDomains(domains)
		assert.Equal(t, domains, d.Domains())
	})
	t.Run("AddDomain", func(t *testing.T) {
		d := &ClientData{}
		d.AddDomain("foo.example.org")
		assert.Equal(t, []string{"foo.example.org"}, d.Domains())
		d.AddDomain("bad.example.org")
		assert.Equal(t, []string{"foo.example.org", "bad.example.org"}, d.Domains())
	})
	t.Run("SetIPv4", func(t *testing.T) {
		d := &ClientData{}
		err := d.SetIPv4("100.100.100.100")
		assert.Equal(t, "100.100.100.100", d.IPv4())
		assert.Nil(t, err)
	})
	t.Run("SetIPv6", func(t *testing.T) {
		d := &ClientData{}
		err := d.SetIPv6("fd00::dead")
		assert.Equal(t, "fd00::dead", d.IPv6())
		assert.Nil(t, err)
	})
}
