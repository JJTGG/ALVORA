package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	defaultPort = 8080
	minPort     = 1
	maxPort     = 65535
)

// Config contains process-level configuration for an ALVORA service.
type Config struct {
	Address string
	Port    int
}

// Load reads configuration from the environment and applies safe defaults.
func Load(getenv func(string) string) (Config, error) {
	if getenv == nil {
		getenv = os.Getenv
	}

	port := defaultPort
	if raw := strings.TrimSpace(getenv("ALVORA_PORT")); raw != "" {
		parsed, err := strconv.Atoi(raw)
		if err != nil || parsed < minPort || parsed > maxPort {
			return Config{}, fmt.Errorf("ALVORA_PORT must be an integer between %d and %d", minPort, maxPort)
		}
		port = parsed
	}

	address := strings.TrimSpace(getenv("ALVORA_ADDRESS"))
	if address == "" {
		address = "0.0.0.0"
	}

	return Config{Address: address, Port: port}, nil
}
