package config

import (
	"rest-api-file-server/env"
	"time"
)

type HttpServerConfig struct {
	Host         string        `json:"host,omitempty"`
	Protocol     string        `json:"protocol,omitempty"`
	HTTPPort     int           `json:"http-port,omitempty"`
	HTTPSPort    int           `json:"https-port,omitempty"`
	CertFile     string        `json:"cert-file,omitempty"`
	CertKey      string        `json:"cert-key,omitempty"`
	ReadTimeout  time.Duration `json:"read-timeout,omitempty"`
	WriteTimeout time.Duration `json:"write-timeout,omitempty"`
	TlsEnabled   bool          `json:"tls-enabled,omitempty"`
}

func NewHttpServerConfig(env *env.Env) HttpServerConfig {
	return HttpServerConfig{
		Host:         env.HttpServerHost,
		Protocol:     env.HttpServerProtocol,
		HTTPPort:     env.HttpServerPort,
		HTTPSPort:    env.HttpServerTLSPort,
		CertFile:     env.HttpServerCertFile,
		CertKey:      env.HttpServerCertKey,
		ReadTimeout:  env.HttpServerReadTimeoutSeconds,
		WriteTimeout: env.HttpServerWriteTimeoutSeconds,
		TlsEnabled:   env.HttpServerTlsEnabled,
	}
}
