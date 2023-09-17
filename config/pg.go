package config

import (
	"rest-api-file-server/env"
	"time"
)

type PostgresConfig struct {
	DSN                       string        `json:"url,omitempty"`
	MaxIdleConn               int           `json:"max-idle-conn,omitempty"`
	MaxOpenConn               int           `json:"max-open-conn,omitempty"`
	ConnMaxLifetime           time.Duration `json:"conn-max-lifetime,omitempty"`
	ConnMaxIdleTime           time.Duration `json:"conn-max-idle-time,omitempty"`
	Database                  string        `json:"name,omitempty"`
	User                      string        `json:"user,omitempty"`
	Pass                      string        `json:"password,omitempty"`
	Host                      string        `json:"host,omitempty"`
	Port                      int           `json:"port,omitempty"`
	PreparedStatementsEnabled bool          `json:"prepared-statements-enabled,omitempty"`
	MigrationsUrl             string        `json:"migrationsUrl,omitempty"`
}

func NewPostgresConfig(env *env.Env) *PostgresConfig {
	return &PostgresConfig{
		DSN:                       env.PostgresDSN,
		MaxIdleConn:               env.PostgresMaxIdleConnections,
		MaxOpenConn:               env.PostgresMaxOpenConnections,
		ConnMaxLifetime:           env.PostgresConnectionMaxLifetimeSeconds,
		ConnMaxIdleTime:           env.PostgresConnectionMaxIdleTimeSeconds,
		Database:                  env.PostgresDatabase,
		User:                      env.PostgresDatabaseUser,
		Pass:                      env.PostgresDatabasePassword,
		Host:                      env.PostgresHost,
		Port:                      env.PostgresPort,
		PreparedStatementsEnabled: true,
		MigrationsUrl:             env.PostgresMigrationsUrl,
	}
}
