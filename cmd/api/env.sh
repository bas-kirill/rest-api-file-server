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