package utils

import (
	"fmt"
	"net/url"
)

type Host struct {
	Hostname string
	Port     int
}

type URI struct {
	Protocol string
	Uri      string
	Host     Host
}

type RequestData struct {
	Params url.Values
	Body   []byte
}

func (server Host) GetHostString() string {
	return fmt.Sprintf("%s:%d", server.Hostname, server.Port)
}

func (location URI) ToUrl() string {
	if len(location.Protocol) > 0 {
		return fmt.Sprintf("%s://%s%s", location.Protocol, location.Host.GetHostString(), location.Uri)
	}
	return fmt.Sprintf("%s%s", location.Host.GetHostString(), location.Uri)
}
