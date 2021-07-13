package record

import "github.com/ytake/draft/config"

func NewRecorder(config *config.TomlConfig) Documenter {
	switch config.Driver {
	case "memcahced":
		return NewMemcachedConnect(config.Memcached.Servers...)
	}
	return nil
}
