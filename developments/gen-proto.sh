#!/bin/sh

#* variables
PROTO_PATH=./api/proto
PROTO_OUT=./idl/pb
DOC_OUT=./docs

#! create nat and cobra folder
mkdir -p ${PROTO_OUT}/nats_pb
mkdir -p ${PROTO_OUT}/cobra_pb
mkdir -p ${DOC_OUT}/html
mkdir -p ${DOC_OUT}/markdown
mkdir -p ${DOC_OUT}/swagger

#* gen normal proto
protoc \
	${PROTO_PATH}/*.proto \
	-I=/usr/local/include \
	--proto_path=${PROTO_PATH} \
	--go_out=:${PROTO_OUT} \
	--validate_out=lang=go:${PROTO_OUT} \
	--go-grpc_out=:${PROTO_OUT} \
	--grpc-gateway_out=:${PROTO_OUT} \
	--openapiv2_out=:${DOC_OUT}/swagger \
	--custom_out=:${PROTO_OUT} \
	--fieldmask_out=lang=go:${PROTO_OUT} \
	--doc_out=:${DOC_OUT}/html --doc_opt=html,index.html

#* gen markdown and tranformer
protoc \
	${PROTO_PATH}/*.proto \
	-I=/usr/local/include \
	--proto_path=${PROTO_PATH} \
	--struct-transformer_out=package=transform,debug=true,goimports=true,helper-package=transformhelpers:. \
	--doc_out=:${DOC_OUT}/markdown --doc_opt=markdown,docs.md

#* gen nrpc(nats)
protoc \
	${PROTO_PATH}/nats/*.proto \
	-I=/usr/local/include \
	--proto_path=${PROTO_PATH}/nats \
	--go_out=:${PROTO_OUT}/nats_pb \
	--nrpc_out=:${PROTO_OUT}/nats_pb

#* gen cobra(cmd)
protoc \
	${PROTO_PATH}/cobra/*.proto \
	-I=/usr/local/include \
	--go_out=:${PROTO_OUT}/cobra_pb \
	--go-grpc_out=:${PROTO_OUT}/cobra_pb \
	--proto_path=${PROTO_PATH}/cobra \
	--experimental_allow_proto3_optional=:true \
	--cobra_out=plugins=client:${PROTO_OUT}/cobra_pb

#! remove permission folders
chmod -R 777 ${PROTO_OUT}
chmod -R 777 ${DOC_OUT}
