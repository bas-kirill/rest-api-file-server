FROM golang:1.20.4 as build

WORKDIR /go/src/app
COPY go.mod go.sum /go/src/app/
RUN go mod download


COPY . /go/src/app
WORKDIR /go/src/app/cmd/api
RUN CGO_ENBALED=0 GOOS=linux go build -o file-server main.go

EXPOSE 8080

# TODO: use scratch image for optimization
FROM debian:12.1
COPY --from=build /go/src/app/cmd/api /app/cmd/api
COPY --from=build /go/src/app/store/pg/migrations /app/store/pg/migrations
WORKDIR /app/cmd/api
CMD ["./file-server"]
