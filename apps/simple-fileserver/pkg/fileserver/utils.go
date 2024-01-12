package fileserver

import (
	"net/http"
)

func ReadUserIP(req *http.Request) string {
	IPAddress := req.Header.Get("x-real-ip")
	if IPAddress == "" {
		IPAddress = req.Header.Get("x-forwarded-for")
	}
	if IPAddress == "" {
		IPAddress = req.RemoteAddr
	}
	return IPAddress
}
