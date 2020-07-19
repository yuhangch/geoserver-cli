package config

import (
	"fmt"
)

// Server to define a GeoServer server
type Server struct {
	Alias    string
	URL      string
	Username string
	Password string
}

// Servers slice of server.
type Servers []Server

// ServersFmt to format servers.
func (svs Servers) ServersFmt() string {
	if len(svs) < 1 {
		return "No Server founded"
	}
	s := "Servers:\n"
	for key, sv := range svs {
		s += fmt.Sprintf("- %d %s %s \n", key, sv.Alias, sv.URL)
	}
	return s

}

// nameOf get server by alias.
func (svs Servers) nameOf(alias string) *Server {
	for _, v := range svs {
		if v.Alias == alias {
			return &v
		}
	}
	return nil
}

// indexOf get server index by alias.
func (svs Servers) IndexOf(alias string) int {

	for k, v := range svs {
		if v.Alias == alias {
			return k
		}
	}
	return -1
}
