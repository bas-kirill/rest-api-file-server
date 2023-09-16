package config

type Config struct {
	Log        LogConfig        `json:"log,omitempty"`
	HttpServer HttpServerConfig `json:"server,omitempty"`
}
