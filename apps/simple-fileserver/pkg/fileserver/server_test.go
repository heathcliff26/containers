package fileserver

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggingWrapper(t *testing.T) {
	fs := NewFileserver("./testdata", true)

	req := httptest.NewRequest(http.MethodGet, "/test.html", nil)
	rr := httptest.NewRecorder()

	var buf bytes.Buffer
	log.SetOutput(&buf)
	t.Cleanup(func() {
		log.SetOutput(os.Stderr)
	})

	assert := assert.New(t)

	fs.loggingWrapper(rr, req)

	assert.Empty(buf.String())

	fs.Log = true

	fs.loggingWrapper(rr, req)
	output := buf.String()

	assert.Contains(output, "status=200")
	assert.Contains(output, "path=\"/test.html\"")
}
