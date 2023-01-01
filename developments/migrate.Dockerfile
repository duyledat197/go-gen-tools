FROM golang:1.19.1 AS migrate

ENV GO111MODULE=on

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest