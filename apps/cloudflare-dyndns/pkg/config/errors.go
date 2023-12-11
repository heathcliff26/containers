package config

type ErrUnknownLogLevel struct {
	Level string
}

func (e *ErrUnknownLogLevel) Error() string {
	return "Unknown log level " + e.Level
}
