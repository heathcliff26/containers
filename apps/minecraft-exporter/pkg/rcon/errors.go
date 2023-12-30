package rcon

type ErrRCONMissingHost struct{}

func (e ErrRCONMissingHost) Error() string {
	return "Missing target host for RCON"
}

type ErrRCONMissingPort struct{}

func (e ErrRCONMissingPort) Error() string {
	return "Missing target port for RCON"
}

type ErrRCONMissingPassword struct{}

func (e ErrRCONMissingPassword) Error() string {
	return "Missing password for RCON"
}

type ErrRCONConnectionTimeout struct{}

func (e ErrRCONConnectionTimeout) Error() string {
	return "Timed out waiting for a response"
}
