package config

type LogConfig struct {
	Level string `json:"level,omitempty"`
	File  string `json:"file,omitempty"`
}
