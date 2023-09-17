#!/bin/bash

# HTTP server settings
export HTTP_SERVER_HOST=localhost
export HTTP_SERVER_PROTOCOL=http
export HTTP_SERVER_PORT=36000
export HTTP_SERVER_TLS_PORT=37000
export HTTP_SERVER_CERT_FILE=rest-api-file-server.crt
export HTTP_SERVER_CERT_KEY=rest-api-file-server.key
export HTTP_SERVER_READ_TIMEOUT_SECONDS=60
export HTTP_SERVER_WRITE_TIMEOUT_SECONDS=60
export HTTP_SERVER_TLS_ENABLED=False

# File server settings
export FILE_SERVER_BASE_PATH=/Users/eertree_work/experiments/rest-api-file-server/test-data

# Postgres settings
export POSTGRES_DSN=postgres://fileserver:fileserver@localhost:5432/fileserver?sslmode=disable
export POSTGRES_HOST=localhost
export POSTGRES_PORT=5432
export POSTGRES_DATABASE=fileserver
export POSTGRES_DATABASE_USER=fileserver
export POSTGRES_DATABASE_PASSWORD=fileserver
export POSTGRES_MAX_IDLE_CONNECTIONS=20
export POSTGRES_MAX_OPEN_CONNECTIONS=100
export POSTGRES_CONNECTION_MAX_LIFETIME_SECONDS=300
export POSTGRES_CONNECTION_MAX_IDLE_TIME_SECONDS=300
export POSTGRES_MIGRATIONS_URL=file://../../store/pg/migrations
