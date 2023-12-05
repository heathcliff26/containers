package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogLevelVarInitialized(t *testing.T) {
	if logLevel == nil {
		t.Fatal("Variable logLevel is nil")
	}
}

func TestServerRootHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	rr := httptest.NewRecorder()

	ServerRootHandler(rr, req)

	assert := assert.New(t)

	assert.Equal(http.StatusOK, rr.Code)
	body := rr.Body.String()
	assert.Contains(body, "<html>")
	assert.Contains(body, "</html>")
	assert.Contains(body, "<a href='/metrics'>")
}

func TestCreateSpeedtest(t *testing.T) {
	t.Run("SpeedtestCLI", func(t *testing.T) {
		s, err := createSpeedtest("./main_test.go")
		if err != nil {
			t.Fatalf("Failed to create speedtest: %v", err)
		}
		assert.Equal(t, "*speedtest.SpeedtestCLI", reflect.TypeOf(s).String())
	})
	t.Run("Speedtest", func(t *testing.T) {
		s, err := createSpeedtest("")
		if err != nil {
			t.Fatalf("Failed to create speedtest: %v", err)
		}
		assert.Equal(t, "*speedtest.Speedtest", reflect.TypeOf(s).String())
	})
}