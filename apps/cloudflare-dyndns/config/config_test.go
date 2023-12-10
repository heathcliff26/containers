package config

import (
	"log/slog"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidConfigs(t *testing.T) {
	c1 := Config{
		LogLevel: "info",
		Server: ServerConfig{
			Port:    8080,
			Domains: []string{"example.org", "example.net"},
		},
		Client: ClientConfig{
			Secret:       "test-token-1",
			Proxy:        true,
			Domains:      []string{"foo.example.org"},
			Interval:     "5m",
			IntervalTime: time.Duration(5 * time.Minute),
		},
	}
	c2 := Config{
		LogLevel: "debug",
		Server: ServerConfig{
			Port:    80,
			Domains: []string{"example.com"},
		},
		Client: ClientConfig{
			Secret:       "test-token-2",
			Proxy:        false,
			Domains:      []string{"bar.example.net"},
			Interval:     "10m",
			IntervalTime: time.Duration(10 * time.Minute),
		},
	}
	tMatrix := []struct {
		Name, Path, Mode string
		Result           Config
	}{
		{
			Name:   "EmptyConfig",
			Path:   "",
			Mode:   MODE_SERVER,
			Result: DefaultConfig(),
		},
		{
			Name:   "ServerConfig1",
			Path:   "testdata/valid-config-1.yaml",
			Mode:   MODE_SERVER,
			Result: c1,
		},
		{
			Name:   "ServerConfig2",
			Path:   "testdata/valid-config-2.yaml",
			Mode:   MODE_SERVER,
			Result: c2,
		},
		{
			Name:   "ClientConfig1",
			Path:   "testdata/valid-config-1.yaml",
			Mode:   MODE_CLIENT,
			Result: c1,
		},
		{
			Name:   "ClientConfig2",
			Path:   "testdata/valid-config-2.yaml",
			Mode:   MODE_CLIENT,
			Result: c2,
		},
	}

	for _, tCase := range tMatrix {
		t.Run(tCase.Name, func(t *testing.T) {
			c, err := LoadConfig(tCase.Path, tCase.Mode)

			assert := assert.New(t)

			if !assert.Nil(err) {
				t.Fatalf("Failed to load config: %v", err)
			}
			if tCase.Mode != MODE_CLIENT {
				// The value will only be set when mode is client
				tCase.Result.Client.IntervalTime = 0
			}
			assert.Equal(tCase.Result, c)
		})
	}
}

func TestInvalidConfig(t *testing.T) {
	tMatrix := []struct {
		Name, Path string
		Error      string
	}{
		{
			Name:  "InvalidPath",
			Path:  "file-does-not-exist.yaml",
			Error: "*fs.PathError",
		},
		{
			Name:  "NotYaml",
			Path:  "testdata/not-a-config.txt",
			Error: "*yaml.TypeError",
		},
		{
			Name:  "ClientMissingSecret",
			Path:  "testdata/invalid-config-1.yaml",
			Error: "client.MissingSecretError",
		},
		{
			Name:  "ClientNoDomain",
			Path:  "testdata/invalid-config-2.yaml",
			Error: "client.NoDomainError",
		},
		{
			Name:  "ClientWrongInterval",
			Path:  "testdata/invalid-config-3.yaml",
			Error: "*errors.errorString",
		},
	}

	for _, tCase := range tMatrix {
		t.Run(tCase.Name, func(t *testing.T) {
			_, err := LoadConfig(tCase.Path, MODE_CLIENT)

			if !assert.Equal(t, tCase.Error, reflect.TypeOf(err).String()) {
				t.Fatalf("Received invalid error: %v", err)
			}
		})
	}
}

func TestSetLogLevel(t *testing.T) {
	tMatrix := []struct {
		Name  string
		Level slog.Level
		Error error
	}{
		{"debug", slog.LevelDebug, nil},
		{"info", slog.LevelInfo, nil},
		{"warn", slog.LevelWarn, nil},
		{"error", slog.LevelError, nil},
		{"DEBUG", slog.LevelDebug, nil},
		{"INFO", slog.LevelInfo, nil},
		{"WARN", slog.LevelWarn, nil},
		{"ERROR", slog.LevelError, nil},
		{"Unknown", 0, &UnknownLogLevelError{"Unknown"}},
	}
	t.Cleanup(func() {
		err := setLogLevel(DEFAULT_LOG_LEVEL)
		if err != nil {
			t.Fatalf("Failed to cleanup after test: %v", err)
		}
	})

	for _, tCase := range tMatrix {
		t.Run(tCase.Name, func(t *testing.T) {
			err := setLogLevel(tCase.Name)

			assert := assert.New(t)

			if !assert.Equal(tCase.Error, err) {
				t.Fatalf("Received invalid error: %v", err)
			}
			if err == nil {
				assert.Equal(tCase.Level, logLevel.Level())
			}
		})
	}
}
