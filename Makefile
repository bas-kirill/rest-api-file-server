build:
	go build -o cmd/api/file-server cmd/api/main.go

run:
	cd cmd/api; ./rundev.sh
