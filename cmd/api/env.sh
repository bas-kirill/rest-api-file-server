#!/bin/bash

# HTTP server settings
export HTTP_ADDR=:80
export HTTPS_ADDR=:443
export HTTP_SERVER_CERT_FILE=./tls/localhost.crt
export HTTP_SERVER_CERT_KEY=./tls/localhost.key
export HTTP_SERVER_READ_TIMEOUT_SECONDS=60
export HTTP_SERVER_WRITE_TIMEOUT_SECONDS=60
export HTTP_SERVER_TLS_ENABLED=True

# File server settings
export FILE_SERVER_BASE_PATH=/tmp/rest-api-file-server

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
