PROJECT_NAME := duyledat197/go-template
PKG := github.com/$(PROJECT_NAME)
MOD := $(GOPATH)/pkg/mod
build:
	@go build -i -v $(PKG)/cmd/server
run:
	go run cmd/server/main.go
test:
	go test ./...
install:
	@go install ./cmd/server/.
gen-proto:
	protoc \
		proto/*.proto \
		-I $(MOD)/github.com/envoyproxy/protoc-gen-validate@v0.9.1 \
		-I $(MOD)/github.com/googleapis/googleapis@v0.0.0-20221209211743-f7f499371afa \
		--proto_path=proto \
		--go_out=:. \
		--validate_out=lang=go:. \
		--go-grpc_out=:. \
		--grpc-gateway_out=:. \
		--openapiv2_out=:docs/swagger \
		--struct-transformer_out=package=transform,debug=true,goimports=true:. \
		--syllabus_out=:. \
		--doc_out=./docs/html --doc_opt=html,index.html 

build:
	@echo "building..."
	go build ./cmd/srv/.
#  --gofullmethods_out=. \
# 	 --gofullmethods_opt=paths=source_relative