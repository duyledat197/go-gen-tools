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
	docker compose -f ${COMPOSE_FILE} up postgres -d
migrate:
	docker compose -f ${COMPOSE_FILE} up migrate
gen-layer:
	go run ./cmd/gen-layer/.
	go fmt ./internal
gen-mock:
	docker compose -f ${COMPOSE_FILE} up generate_mock
bdd-test:
	docker compose -f ${COMPOSE_FILE} up bdd_test
build:
	@echo "building..."
	go build ./cmd/srv/.
#  --gofullmethods_out=. \
# 	 --gofullmethods_opt=paths=source_relative