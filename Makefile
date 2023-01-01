PROJECT_NAME := test
PKG := github.com/$(PROJECT_NAME)
MOD := $(GOPATH)/pkg/mod
COMPOSE_FILE := ./developments/docker-compose.yml

# build:
# 	@go build -i -v $(PKG)/cmd/server
run:
	go run cmd/server/main.go
test:
	go test ./...
install:
	@go install ./cmd/server/.
gen-sql:
	docker compose -f ${COMPOSE_FILE} up generate_sqlc
gen-proto:
	docker compose -f ${COMPOSE_FILE} up generate_pb_go --build
start-postgres:
	docker compose -f ${COMPOSE_FILE} up postgres -d --build
migrate:
	docker compose -f ${COMPOSE_FILE} up migrate
build:
	@echo "building..."
	go build ./cmd/srv/.
#  --gofullmethods_out=. \
# 	 --gofullmethods_opt=paths=source_relative