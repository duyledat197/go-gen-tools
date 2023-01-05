
FROM golang:1.19.1 

WORKDIR /app

ENV GO111MODULE=on
RUN apt update && apt install -y --no-install-recommends curl make git unzip apt-utils
RUN go install github.com/cucumber/godog/cmd/godog@latest


