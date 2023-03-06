
FROM golang:1.19.1 

WORKDIR /app

ENV GO111MODULE=on
RUN go install github.com/cucumber/godog/cmd/godog@latest


