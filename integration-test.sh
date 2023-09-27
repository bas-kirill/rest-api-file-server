docker-compose -f docker-compose.test.yml build --no-cache  # by security reasons do not use cache
docker-compose -f docker-compose.test.yml up --abort-on-container-exit
docker-compose -f docker-compose.test.yml down --volumes