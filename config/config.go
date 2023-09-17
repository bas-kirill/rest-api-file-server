package config

// Config is a config :)
type Config struct {
	Log        LogConfig        `json:"log,omitempty"`
	HttpServer HttpServerConfig `json:"server,omitempty"`
	FileServer FileServerConfig `json:"fileServer,omitempty"`
}
