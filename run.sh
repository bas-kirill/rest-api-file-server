#!/bin/bash
source ./cmd/api/env.sh
docker-compose -f docker-compose.yml build --no-cache  # by security reasons do not use cache
docker-compose -f docker-compose.yml up
