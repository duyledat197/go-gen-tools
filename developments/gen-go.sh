#!/bin/sh

#* gen normal proto
protoc \
	./proto/*.proto \
	-I=/usr/local/include \
	--proto_path=./proto \
	--go_out=:. \
	--validate_out=lang=go:. \
	--go-grpc_out=:. \
	--grpc-gateway_out=:. \
	--openapiv2_out=:docs/swagger \
	--custom_out=:. \
	--fieldmask_out=lang=go:. \
	--doc_out=:./docs/html --doc_opt=html,index.html

#! create nat and cobra folder
mkdir ./pb/natspb
mkdir ./pb/cobra
mkdir docs

#* gen markdown and tranformer
protoc \
	./proto/*.proto \
	-I=/usr/local/include \
	--proto_path=./proto \
	--struct-transformer_out=package=transform,debug=true,goimports=true,helper-package=transformhelpers:. \
	--doc_out=:./docs/markdown --doc_opt=markdown,docs.md

#* gen nrpc(nats)
protoc \
	./proto/nats/*.proto \
	-I=/usr/local/include \
	--proto_path=./proto/nats \
	--go_out=:. \
	--nrpc_out=:./pb/natspb

#* gen cobra(cmd)
protoc \
	./proto/cobra/*.proto \
	-I=/usr/local/include \
	--go_out=:./pb/cobra \
	--go-grpc_out=:./pb/cobra \
	--proto_path=./proto/cobra \
	--experimental_allow_proto3_optional=:true \
	--cobra_out=plugins=client:./pb/cobra

#! remove permission folders
chmod -R 777 ./pb
