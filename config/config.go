package config

import (
	"encoding/base64"
	"fmt"
)

type Config struct {
	Servers Servers
	Server  int
}

// BasicAuth to return current server's basic auth.
func (cfg *Config) BasicAuth() (string, error) {
	i := cfg.Server
	if i >= len(cfg.Servers) {
		return "", fmt.Errorf("Can't get current server info")
	}
	server := cfg.Servers[i]
	auth := server.Username + ":" + server.Password
	return base64.StdEncoding.EncodeToString([]byte(auth)), nil
}

// ServerURL get current server's url.
func (cfg *Config) ServerURL() string {
	return cfg.currentServer().URL
}

// currentServer get current server.
func (cfg *Config) currentServer() Server {
	i := cfg.Server
	if i >= len(cfg.Servers) {
		panic("Can't get current server info")
	}
	return cfg.Servers[i]
}
