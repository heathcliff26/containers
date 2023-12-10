package config

type UnknownLogLevelError struct {
	level string
}

func (e *UnknownLogLevelError) Error() string {
	return "Unknown log level " + e.level
}
