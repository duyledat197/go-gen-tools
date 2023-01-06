#!/bin/sh

protoc \
    ./proto/*.proto \
		-I=/usr/local/include \
		--proto_path=./proto \
		--validate_out=lang=go:. \
		--go-grpc_out=:. \
		--grpc-gateway_out=:. \
		--openapiv2_out=:docs/swagger \
		--custom_out=:. \
		--fieldmask_out=lang=go:. \
		--doc_out=:./docs/html --doc_opt=html,index.html 

protoc \
    ./proto/*.proto \
		-I=/usr/local/include \
		--proto_path=./proto \
		--nrpc_out=:./pb \
		--struct-transformer_out=package=transform,debug=true,goimports=true,helper-package=transformhelpers:. \
		--doc_out=:./docs/markdown --doc_opt=markdown,docs.md

gofumpt -l -w ./transform/*.go
