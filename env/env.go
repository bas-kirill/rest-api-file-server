package env

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Env struct {
	HttpAddr                      string
	HttpsAddr                     string
	HttpServerCertFile            string
	HttpServerCertKey             string
	HttpServerReadTimeoutSeconds  time.Duration
	HttpServerWriteTimeoutSeconds time.Duration
	HttpServerTlsEnabled          bool

	FileServerBasePath string

	PostgresDSN                          string
	PostgresHost                         string
	PostgresPort                         int
	PostgresDatabase                     string
	PostgresDatabaseUser                 string
	PostgresDatabasePassword             string
	PostgresMaxIdleConnections           int
	PostgresMaxOpenConnections           int
	PostgresConnectionMaxLifetimeSeconds time.Duration
	PostgresConnectionMaxIdleTimeSeconds time.Duration
	PostgresMigrationsUrl                string
}

func NewEnv() *Env {
	return &Env{
		HttpAddr:                             getEnvStr(HttpAddr),
		HttpsAddr:                            getEnvStr(HttpsAddr),
		HttpServerCertFile:                   getEnvStr(HttpServerCertFile),
		HttpServerCertKey:                    getEnvStr(HttpServerCertKey),
		HttpServerReadTimeoutSeconds:         getEnvSeconds(HttpServerReadTimeoutSeconds),
		HttpServerWriteTimeoutSeconds:        getEnvSeconds(HttpServerWriteTimeoutSeconds),
		HttpServerTlsEnabled:                 getEnvBool(HttpServerTlsEnabled),
		FileServerBasePath:                   getEnvStr(FileServerBasePath),
		PostgresDSN:                          getEnvStr(PostgresDSN),
		PostgresHost:                         getEnvStr(PostgresHost),
		PostgresPort:                         getEnvInt(PostgresPort),
		PostgresDatabase:                     getEnvStr(PostgresDatabaseName),
		PostgresDatabaseUser:                 getEnvStr(PostgresDatabaseUser),
		PostgresDatabasePassword:             getEnvStr(PostgresDatabasePassword),
		PostgresMaxIdleConnections:           getEnvInt(PostgresMaxIdleConnections),
		PostgresMaxOpenConnections:           getEnvInt(PostgresMaxOpenConnections),
		PostgresConnectionMaxLifetimeSeconds: getEnvSeconds(PostgresConnectionMaxLifetimeSeconds),
		PostgresConnectionMaxIdleTimeSeconds: getEnvSeconds(PostgresMaxIdleConnections),
		PostgresMigrationsUrl:                getEnvStr(PostgresMigrationsUrl),
	}
}

func getEnvStr(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	panic(fmt.Sprintf("can not find %s in environment", key))
}

func getEnvInt(key string) int {
	if value, exists := os.LookupEnv(key); exists {
		i, err := strconv.Atoi(value)
		if err != nil {
			panic(fmt.Errorf("fail to parse value=%s to integer", err))
		}
		return i
	}
	panic(fmt.Sprintf("can not find %s usecases environment", key))
}

func getEnvBool(key string) bool {
	envStr := getEnvStr(key)
	b, err := strconv.ParseBool(envStr)
	if err != nil {
		panic(fmt.Errorf("failed parse to bool value `%s`", envStr))
	}
	return b
}

func getEnvSeconds(key string) time.Duration {
	return time.Duration(getEnvInt(key)) * time.Second
}
