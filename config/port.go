package config

import (
	"flag"
	"strconv"
)

const (
	DefaultPort = 8080
)

type (
	// Bootstrapper for app server option
	Bootstrapper interface {
		PortOption() string
	}
	// Config for port
	Config struct {
		Port string
	}
)

// NewConfig make config
func NewConfig() *Config {
	return &Config{portOption()}
}

// BootOption server port
func portOption() string {
	port := flag.Int("port", DefaultPort, "specified HTTP server port")
	flag.Parse()
	return strconv.Itoa(*port)
}
