/*
Copyright © 2020 Yuhang Chen <i@yuhang.ch>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
