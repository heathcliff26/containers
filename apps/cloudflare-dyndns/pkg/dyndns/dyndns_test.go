package dyndns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientDataGetFunctions(t *testing.T) {
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

func TestClientDataSetFunctions(t *testing.T) {
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

func TestClientDataCheckUpdate(t *testing.T) {
	d := NewClientData(false)

	assert := assert.New(t)

	// Testing with no IPs
	err := d.CheckData()
	assert.ErrorIs(err, ErrNoIP{})

	// Testing with no Domains and IPv4 only
	err = d.SetIPv4("100.100.100.100")
	assert.Nil(err)
	err = d.CheckData()
	assert.ErrorIs(err, ErrNoDomain{})

	// Testing with no Domains and dual stack
	err = d.SetIPv6("fd00::dead")
	assert.Nil(err)
	err = d.CheckData()
	assert.ErrorIs(err, ErrNoDomain{})

	// Testing with no Domains and IPv6 only
	err = d.SetIPv4("")
	assert.Nil(err)
	err = d.CheckData()
	assert.ErrorIs(err, ErrNoDomain{})
}

func TestRunUpdate(t *testing.T) {
	c := &testClient{
		data:       NewClientData(true),
		FailUpdate: false,
	}
	c.Data().SetDomains([]string{"foo.example.org"})
	var updated = false

	assert := assert.New(t)

	assert.Equal(0, c.UpdateCount, "No updates have been run yet")
	runUpdate(c, &updated)
	assert.Equal(1, c.UpdateCount, "Counter should have increased to 1")
	assert.Equal(true, updated, "updated should now be true")

	runUpdate(c, &updated)
	assert.Equal(1, c.UpdateCount, "Counter should stay at 1")
	assert.Equal(true, updated, "updated should still be true")

	// Reset client for second test
	err1 := c.Data().SetIPv4("")
	err2 := c.Data().SetIPv6("")
	c.FailUpdate = true
	if err1 != nil || err2 != nil {
		t.Fatalf("Unexpected error updating IPs, err1=%v, err2=%v", err1, err2)
	}

	runUpdate(c, &updated)
	assert.Equal(1, c.UpdateCount, "Counter should stay at 1")
	assert.Equal(false, updated, "updated should now be false")

	// Third run should run again
	c.FailUpdate = false
	runUpdate(c, &updated)
	assert.Equal(2, c.UpdateCount, "Counter should have increased to 2")
	assert.Equal(true, updated, "updated should now be true again")
}
