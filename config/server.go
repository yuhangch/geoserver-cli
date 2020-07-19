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

// ServersFmt to format servers.
func ServersFmt(servrers []Server) string {
	if len(servrers) < 1 {
		return "No Server founded"
	}
	s := "Servers:\n"
	for key, sv := range servrers {
		s += fmt.Sprintf("- %d %s %s \n", key, sv.Alias, sv.URL)
	}
	return s

}
