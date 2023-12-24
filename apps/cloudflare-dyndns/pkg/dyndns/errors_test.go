package dyndns

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewErrHttpRequestFailed(t *testing.T) {
	result := &ErrHttpRequestFailed{StatusCode: 400, Body: "testresult"}
	r := io.NopCloser(strings.NewReader(result.Body))
	defer r.Close()
	err := NewErrHttpRequestFailed(400, r)
	assert.Equal(t, result, err)
}

func TestNewErrOperationFailed(t *testing.T) {
	result := &ErrOperationFailed{Result: "testresult"}
	r := io.NopCloser(strings.NewReader(result.Result))
	defer r.Close()
	err := NewErrOperationFailed(r)
	assert.Equal(t, result, err)
}
