package config

type UnknownLogLevelError struct {
	Level string
}

func (e *UnknownLogLevelError) Error() string {
	return "Unknown log level " + e.Level
}
