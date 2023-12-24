package dyndns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTestClient(t *testing.T) {
	assert := assert.New(t)

	c, err := NewTestClient("", false)
	assert.Equal(ErrMissingToken{}, err)
	assert.Nil(c)

	c, err = NewTestClient("test", true)
	assert.Nil(err)
	assert.NotEmpty(c)
}
