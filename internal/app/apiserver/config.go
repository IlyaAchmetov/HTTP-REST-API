package apiserver

// Config ...
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	Loglevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: "8080",
		Loglevel: "debug",
	}
}
