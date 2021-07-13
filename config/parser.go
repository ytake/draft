package config

import "github.com/BurntSushi/toml"

type TomlConfig struct {
	Driver    string
	Memcached memcached
}

type memcached struct {
	Servers []string
}

func Parse(file string) (*TomlConfig, error) {
	var config TomlConfig
	if _, err := toml.DecodeFile(file, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
