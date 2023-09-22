package env

const (
	HttpAddr                      = "HTTP_ADDR"
	HttpsAddr                     = "HTTPS_ADDR"
	HttpServerCertFile            = "HTTP_SERVER_CERT_FILE"
	HttpServerCertKey             = "HTTP_SERVER_CERT_KEY"
	HttpServerReadTimeoutSeconds  = "HTTP_SERVER_READ_TIMEOUT_SECONDS"
	HttpServerWriteTimeoutSeconds = "HTTP_SERVER_WRITE_TIMEOUT_SECONDS"
	HttpServerTlsEnabled          = "HTTP_SERVER_TLS_ENABLED"

	FileServerBasePath = "FILE_SERVER_BASE_PATH"

	PostgresDSN                          = "POSTGRES_DSN"
	PostgresHost                         = "POSTGRES_HOST"
	PostgresPort                         = "POSTGRES_PORT"
	PostgresDatabaseName                 = "POSTGRES_DATABASE"
	PostgresDatabaseUser                 = "POSTGRES_DATABASE_USER"
	PostgresDatabasePassword             = "POSTGRES_DATABASE_PASSWORD"
	PostgresMaxIdleConnections           = "POSTGRES_MAX_IDLE_CONNECTIONS"
	PostgresMaxOpenConnections           = "POSTGRES_MAX_OPEN_CONNECTIONS"
	PostgresConnectionMaxLifetimeSeconds = "POSTGRES_CONNECTION_MAX_LIFETIME_SECONDS"
	PostgresConnectionMaxIdleTimeSeconds = "POSTGRES_CONNECTION_MAX_IDLE_TIME_SECONDS"
	PostgresMigrationsUrl                = "POSTGRES_MIGRATIONS_URL"
)
