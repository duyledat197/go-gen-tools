#!/bin/sh

protoc \
    ./proto/*.proto \
		-I=/usr/local/include \
		--proto_path=./proto \
		--go_out=:. \
		--validate_out=lang=go:. \
		--go-grpc_out=:. \
		--grpc-gateway_out=:. \
		--openapiv2_out=:docs/swagger \
		--struct-transformer_out=package=transform,debug=true,goimports=true,helper-package=transformhelpers:. \
		--custom_out=:. \
		--doc_out=./docs/html --doc_opt=html,index.html 

gofumpt -l -w ./transform/*.go
