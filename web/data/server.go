package data

import (
	"fmt"
	"net/url"
)

type Server struct {
	Hostname string
	Port     int
}

type ServerLocation struct {
	Uri    string
	Server Server
}

type RequestData struct {
	Params url.Values
	Body   []byte
}

func (server Server) GetServer() string {
	return fmt.Sprintf("%s:%d", server.Hostname, server.Port)
}

func (location ServerLocation) GetFullURI() string {
	return fmt.Sprintf("%s%s", location.Server.GetServer(), location.Uri)
}
