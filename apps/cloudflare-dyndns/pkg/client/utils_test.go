package client

import (
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
