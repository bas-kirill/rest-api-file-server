#!/bin/bash
source ./cmd/api/env.sh
docker-compose up --build
# todo: add --no-cache for security reasons