#? build stage
FROM golang:1.20.0-alpine AS build-stage
ENV GOBIN=/usr/local/bin/
ENV GO111MODULE=on
RUN go install github.com/ethereum/go-ethereum/cmd/abigen@latest

#* main stage
FROM ethereum/solc:0.8.19-alpine AS abi_gen_contract
COPY --from=build-stage /usr/local/bin/. /usr/local/bin/.
