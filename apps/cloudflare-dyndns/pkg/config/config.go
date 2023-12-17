package config

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/heathcliff26/containers/apps/cloudflare-dyndns/pkg/dyndns"
	"gopkg.in/yaml.v3"
)

const (
	DEFAULT_LOG_LEVEL       = "info"
	DEFAULT_SERVER_PORT     = 8080
	DEFAULT_CLIENT_INTERVAL = time.Duration(5 * time.Minute)

	MODE_SERVER = "server"
	MODE_CLIENT = "client"
	MODE_RELAY  = "relay"
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
	Port    int       `yaml:"port"`
	Domains []string  `yaml:"domains,omitempty"`
	SSL     SSLConfig `yaml:"ssl,omitempty"`
}

type SSLConfig struct {
	Enabled bool   `yaml:"enabled,omitempty"`
	Cert    string `yaml:"cert,omitempty"`
	Key     string `yaml:"key,omitempty"`
}

// Yaml configuration for dyndns client
type ClientConfig struct {
	Token    string        `yaml:"token"`
	Proxy    bool          `yaml:"proxy,omitempty"`
	Domains  []string      `yaml:"domains"`
	Interval time.Duration `yaml:"interval,omitempty"`
	Endpoint string        `yaml:"endpoint,omitempty"`
}

// Validate the server part of the config
func (c *Config) validateServer() error {
	if c.Server.SSL.Enabled && (c.Server.SSL.Cert == "" || c.Server.SSL.Key == "") {
		return ErrIncompleteSSLConfig{}
	}
	slog.Info("Loaded server config",
		slog.Int("port", c.Server.Port),
		slog.String("domains", fmt.Sprintf("%v", c.Server.Domains)),
		slog.Bool("ssl.enabled", c.Server.SSL.Enabled),
		slog.String("ssl.cert", c.Server.SSL.Cert),
		slog.String("ssl.key", c.Server.SSL.Key),
	)
	return nil
}

// Validate the client part of the config
func (c *Config) validateClient() error {
	if c.Client.Token == "" {
		return dyndns.ErrMissingToken{}
	}

	if c.Client.Domains == nil || len(c.Client.Domains) < 1 {
		return dyndns.ErrNoDomain{}
	}

	if c.Client.Interval < time.Duration(30*time.Second) {
		return &ErrInvalidInterval{c.Client.Interval}
	}

	slog.Info("Loaded client config",
		slog.Bool("proxy", c.Client.Proxy),
		slog.String("domains", fmt.Sprintf("%v", c.Client.Domains)),
		slog.String("interval", c.Client.Interval.String()),
		slog.String("endpoint", c.Client.Endpoint),
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
			Interval: DEFAULT_CLIENT_INTERVAL,
		},
	}
}

// Loads config from file, returns error if config is invalid
// Arguments:
//
//	path: Path to config file
//	mode: Mode used, determines how the config will be validated and which values will be processed
//	env: Determines if enviroment variables in the file will be expanded before decoding
func LoadConfig(path string, mode string, env bool) (Config, error) {
	c := DefaultConfig()

	if path == "" && mode == MODE_SERVER {
		_ = setLogLevel(DEFAULT_LOG_LEVEL)
		return c, nil
	}

	f, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	if env {
		f = []byte(os.ExpandEnv(string(f)))
	}

	err = yaml.Unmarshal(f, &c)
	if err != nil {
		return Config{}, err
	}

	err = setLogLevel(c.LogLevel)
	if err != nil {
		return Config{}, err
	}

	switch mode {
	case MODE_SERVER:
		err = c.validateServer()
	case MODE_CLIENT, MODE_RELAY:
		err = c.validateClient()
	}
	if err != nil {
		return Config{}, err
	}

	if mode == MODE_RELAY && c.Client.Endpoint == "" {
		return Config{}, dyndns.ErrMissingEndpoint{}
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
