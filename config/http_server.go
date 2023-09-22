package config

import (
	"rest-api-file-server/env"
	"time"
)

type HttpServerConfig struct {
	HttpAddr     string        `json:"http-addr,omitempty"`
	HttpsAddr    string        `json:"https-addr,omitempty"`
	CertFile     string        `json:"cert-file,omitempty"`
	CertKey      string        `json:"cert-key,omitempty"`
	ReadTimeout  time.Duration `json:"read-timeout,omitempty"`
	WriteTimeout time.Duration `json:"write-timeout,omitempty"`
	TlsEnabled   bool          `json:"tls-enabled,omitempty"`
}

func NewHttpServerConfig(env *env.Env) HttpServerConfig {
	return HttpServerConfig{
		HttpAddr:     env.HttpAddr,
		HttpsAddr:    env.HttpsAddr,
		CertFile:     env.HttpServerCertFile,
		CertKey:      env.HttpServerCertKey,
		ReadTimeout:  env.HttpServerReadTimeoutSeconds,
		WriteTimeout: env.HttpServerWriteTimeoutSeconds,
		TlsEnabled:   env.HttpServerTlsEnabled,
	}
}
