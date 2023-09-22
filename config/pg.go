package config

import (
	"rest-api-file-server/env"
	"time"
)

type PostgresConfig struct {
	DSN             string        `json:"url,omitempty"`
	MaxIdleConn     int           `json:"max-idle-conn,omitempty"`
	MaxOpenConn     int           `json:"max-open-conn,omitempty"`
	ConnMaxLifetime time.Duration `json:"conn-max-lifetime,omitempty"`
	ConnMaxIdleTime time.Duration `json:"conn-max-idle-time,omitempty"`
	MigrationsUrl   string        `json:"migrationsUrl,omitempty"`
}

func NewPostgresConfig(env *env.Env) *PostgresConfig {
	return &PostgresConfig{
		DSN:             env.PostgresDSN,
		MaxIdleConn:     env.PostgresMaxIdleConnections,
		MaxOpenConn:     env.PostgresMaxOpenConnections,
		ConnMaxLifetime: env.PostgresConnectionMaxLifetimeSeconds,
		ConnMaxIdleTime: env.PostgresConnectionMaxIdleTimeSeconds,
		MigrationsUrl:   env.PostgresMigrationsUrl,
	}
}
