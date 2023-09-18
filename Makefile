build:
	go build -v cmd/api

run:
	cd cmd/api; ./rundev.sh

test:
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yml down --volumes
