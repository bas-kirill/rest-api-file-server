package config

import "rest-api-file-server/env"

type FileServerConfig struct {
	BaseSystemPath string `json:"base-system-path,omitempty"`
}

func NewFileServerConfig(env *env.Env) *FileServerConfig {
	return &FileServerConfig{
		BaseSystemPath: env.FileServerBasePath,
	}
}
