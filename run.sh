#!/bin/bash
source ./cmd/api/env.sh
docker-compose -f docker-compose.yml up --build
# todo: add --no-cache for security reasons