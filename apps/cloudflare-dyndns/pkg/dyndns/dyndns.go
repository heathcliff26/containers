package dyndns

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Client interface {
	Data() *ClientData
	Update() error
}

// ClientData is a struct that contains data needed for all Clients
type ClientData struct {
	proxy   bool
	domains []string
	ipv4    string
	ipv6    string
}

func NewClientData(proxy bool) *ClientData {
	return &ClientData{proxy: proxy}
}

func (d *ClientData) Proxy() bool {
	return d.proxy
}

func (d *ClientData) Domains() []string {
	return d.domains
}

func (d *ClientData) IPv4() string {
	return d.ipv4
}

func (d *ClientData) IPv6() string {
	return d.ipv6
}

// Set proxy value
func (d *ClientData) SetProxy(proxy bool) {
	d.proxy = proxy
}

// Set domains to update, overrides previous domains
func (d *ClientData) SetDomains(domains []string) {
	d.domains = domains
}

// Add a domain to the list of domains
func (d *ClientData) AddDomain(domain string) {
	if d.domains == nil {
		d.domains = []string{domain}
	} else {
		d.domains = append(d.domains, domain)
	}
}

// Set IPv4 Address to use for creating entries
// Returns an error if string is neither empty nor a valid IP Address
func (d *ClientData) SetIPv4(val string) error {
	if val != "" && !ValidIPv4(val) {
		return &ErrInvalidIP{IP: val}
	}
	d.ipv4 = val
	return nil
}

// Set IPv6 Address to use for creating entries
// Returns an error if string is neither empty nor a valid IP Address
func (d *ClientData) SetIPv6(val string) error {
	if val != "" && !ValidIPv6(val) {
		return &ErrInvalidIP{IP: val}
	}
	d.ipv6 = val
	return nil
}

// Checks if data contains at least one IP and one domain
// Returns error otherwise
func (d *ClientData) CheckData() error {
	if d.IPv4() == "" && d.IPv6() == "" {
		return ErrNoIP{}
	}
	if d.Domains() == nil || len(d.Domains()) == 0 {
		return ErrNoDomain{}
	}
	return nil
}

// Fetch public IPs and run Update() if changed or last update has not succeeded
func runUpdate(c Client, updated *bool) {
	ipv4, err := GetPublicIPv4()
	if err != nil {
		slog.Error("Failed to get public IPv4", "err", err)
	}
	ipv6, err := GetPublicIPv6()
	if err != nil {
		slog.Error("Failed to get public IPv6", "err", err)
	}

	changed := ipv4 != c.Data().IPv4() || ipv6 != c.Data().IPv6()
	if changed && ipv4 != c.Data().IPv4() {
		err = c.Data().SetIPv4(ipv4)
		if err != nil {
			slog.Error("Failed to get public IPv4", "err", err)
			changed = false
		}
	}
	if changed && ipv6 != c.Data().IPv6() {
		err = c.Data().SetIPv6(ipv6)
		if err != nil {
			slog.Error("Failed to get public IPv6", "err", err)
			changed = false
		}
	}

	if changed || !*updated {
		*updated = false
		slog.Info("Deteced changed IP",
			slog.String("ipv4", c.Data().IPv4()),
			slog.String("ipv6", c.Data().IPv6()),
		)
		err = c.Update()
		if err != nil {
			slog.Error("Failed to Update() records", "err", err)
		} else {
			slog.Info("Updated records", slog.String("IPv4", c.Data().IPv4()), slog.String("IPv6", c.Data().IPv6()))
			*updated = true
		}
	} else {
		slog.Debug("No changed detected")
	}
}

// Fetch the public IP(s) and run Update() periodically.
// Is executed as blocking and for forever.
// Will not continue run of the loop if an error occurs.
// Exits gracefully on SIGTERM
func Run(c Client, interval time.Duration) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	var updated bool
	for {
		runUpdate(c, &updated)

		var elapsedTime time.Duration = 0
		for elapsedTime < interval {
			timer := time.NewTimer(1 * time.Second)
			select {
			case <-timer.C:
				elapsedTime += time.Duration(1 * time.Second)
			case <-quit:
				timer.Stop()
				slog.Info("Received SIGTERM, shutting down client")
				return
			}
		}
	}
}
