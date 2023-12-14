package config

import "time"

type ErrUnknownLogLevel struct {
	Level string
}

func (e *ErrUnknownLogLevel) Error() string {
	return "Unknown log level " + e.Level
}

type ErrInvalidInterval struct {
	Interval time.Duration
}

func (e *ErrInvalidInterval) Error() string {
	return "Interval is to short, needs to be at least 30s, current " + e.Interval.String()
}

type ErrIncompleteSSLConfig struct{}

func (e ErrIncompleteSSLConfig) Error() string {
	return "SSL is enabled but certificate and/or private key are missing"
}
