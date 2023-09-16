build:
	go build -v cmd/api

run:
	cd cmd/api; ./rundev.sh

test:
	go test -v ./... -coverpkg=./...
