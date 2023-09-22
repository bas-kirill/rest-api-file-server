# todo: create multi-stage build + USER
FROM golang:1.20.4

COPY . /go/src/app

WORKDIR /go/src/app/cmd/api

RUN CGO_ENBALED=0 GOOS=linux go build -o file-server main.go

EXPOSE 8080

CMD ["./file-server"]
