PROJECT_NAME := test
PKG := github.com/$(PROJECT_NAME)
MOD := $(GOPATH)/pkg/mod
COMPOSE_FILE := ./developments/proto.docker-compose.yml

# build:
# 	@go build -i -v $(PKG)/cmd/server
run:
	go run cmd/server/main.go
test:
	go test ./...
install:
	@go install ./cmd/server/.
gen-proto:
	docker compose -f ./developments/proto.docker-compose.yml up

build:
	@echo "building..."
	go build ./cmd/srv/.
#  --gofullmethods_out=. \
# 	 --gofullmethods_opt=paths=source_relative