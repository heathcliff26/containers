package relay

import (
	"testing"

	"github.com/heathcliff26/containers/apps/cloudflare-dyndns/pkg/dyndns"
	"github.com/stretchr/testify/assert"
)

func TestNewRelayClient(t *testing.T) {
	c, err := NewRelay("", true, "")

	assert := assert.New(t)

	assert.Equal(dyndns.ErrMissingSecret{}, err)
	assert.Nil(c)

	c, err = NewRelay("testtoken", true, "")

	assert.Equal(dyndns.ErrMissingEndpoint{}, err)
	assert.Nil(c)
}

func TestRelayGetEndpoint(t *testing.T) {
	r := &relay{
		endpoint: "dyndns.example.org",
	}
	assert.Equal(t, "dyndns.example.org", r.Endpoint())
}
