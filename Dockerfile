FROM golang:1.20.4

COPY . /go/src/app

WORKDIR /go/src/app/cmd/api

RUN go build -o file-server main.go

EXPOSE 8080

CMD ["./file-server"]
