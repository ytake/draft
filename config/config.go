package config

import (
	"flag"
	"strconv"
)

const (
	DefaultPort = 8080
)

type (
	// Config for port
	Config struct {
		Port string
		File string
	}
)

// NewConfig make config
func NewConfig() *Config {
	port := flag.Int("port", DefaultPort, "specified HTTP server port")
	file := flag.String("config", "", "specify the toml file")
	flag.Parse()
	return &Config{
		Port: strconv.Itoa(*port),
		File: *file,
	}
}
