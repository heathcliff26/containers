package config

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/heathcliff26/containers/apps/cloudflare-dyndns/pkg/client"
	"gopkg.in/yaml.v3"
)

const (
	DEFAULT_LOG_LEVEL   = "info"
	DEFAULT_SERVER_PORT = 8080
	MODE_CLIENT         = "client"
	MODE_SERVER         = "server"
)

var logLevel *slog.LevelVar

// Initialize the logger
func init() {
	logLevel = &slog.LevelVar{}
	opts := slog.HandlerOptions{
		Level: logLevel,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &opts))
	slog.SetDefault(logger)
}

type Config struct {
	LogLevel string       `yaml:"logLevel,omitempty"`
	Server   ServerConfig `yaml:"server,omitempty"`
	Client   ClientConfig `yaml:"client,omitempty"`
}

// Yaml configuration for dyndns server
type ServerConfig struct {
	Port    int      `yaml:"port"`
	Domains []string `yaml:"domains,omitempty"`
}

// Yaml configuration for dyndns client
type ClientConfig struct {
	Secret       string        `yaml:"secret"`
	Proxy        bool          `yaml:"proxy,omitempty"`
	Domains      []string      `yaml:"domains"`
	Interval     string        `yaml:"interval,omitempty"`
	IntervalTime time.Duration `yaml:"-"`
}

// Validate the client part of the config
func (c *Config) validateClient() error {
	if c.Client.Secret == "" {
		return client.ErrMissingSecret{}
	}

	if c.Client.Domains == nil || len(c.Client.Domains) < 1 {
		return client.ErrNoDomain{}
	}

	// Interval should be a valid duration
	var err error
	c.Client.IntervalTime, err = time.ParseDuration(c.Client.Interval)
	if err != nil {
		return err
	}

	slog.Info("Loaded client config",
		slog.Bool("proxy", c.Client.Proxy),
		slog.String("domains", fmt.Sprintf("%v", c.Client.Domains)),
		slog.String("interval", c.Client.Interval),
	)

	return nil
}

// Returns a Config with default values set
func DefaultConfig() Config {
	return Config{
		LogLevel: DEFAULT_LOG_LEVEL,
		Server:   ServerConfig{Port: DEFAULT_SERVER_PORT},
		Client: ClientConfig{
			Proxy:    true,
			Interval: "5m",
		},
	}
}

// Loads config from file, returns error if config is invalid
func LoadConfig(path string, mode string) (Config, error) {
	c := DefaultConfig()

	if path == "" && mode == MODE_SERVER {
		_ = setLogLevel(DEFAULT_LOG_LEVEL)
		return c, nil
	}

	f, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()

	d := yaml.NewDecoder(f)

	err = d.Decode(&c)
	if err != nil {
		return Config{}, err
	}

	err = setLogLevel(c.LogLevel)
	if err != nil {
		return Config{}, err
	}

	if mode == MODE_CLIENT {
		err = c.validateClient()
		if err != nil {
			return Config{}, err
		}
	}

	return c, nil
}

// Parse a given string and set the resulting log level
func setLogLevel(level string) error {
	switch strings.ToLower(level) {
	case "debug":
		logLevel.Set(slog.LevelDebug)
	case "info":
		logLevel.Set(slog.LevelInfo)
	case "warn":
		logLevel.Set(slog.LevelWarn)
	case "error":
		logLevel.Set(slog.LevelError)
	default:
		return &ErrUnknownLogLevel{level}
	}
	return nil
}
