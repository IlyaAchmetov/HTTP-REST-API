package apiserver

import "github.com/IlyaAchmetov/HTTP-REST-API/internal/app/store"

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	Loglevel string `toml:"log_level"`
	Store    *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
		Loglevel: "debug",
		Store:    store.NewConfig(),
	}
}
