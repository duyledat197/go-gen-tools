PROJECT_NAME := test
PKG := github.com/$(PROJECT_NAME)
MOD := $(GOPATH)/pkg/mod
COMPOSE_FILE := ./developments/docker-compose.yml

# build:
# 	@go build -i -v $(PKG)/cmd/server
run:
	./developments/start.sh
test:
	go test ./...
gen-sql:
	docker compose -f ${COMPOSE_FILE} up generate_sqlc --build
gen-proto:
	docker compose -f ${COMPOSE_FILE} up generate_pb_go --build
gen-contract:
	docker compose -f ${COMPOSE_FILE} up generate_contract --build
start-postgres:
	docker compose -f ${COMPOSE_FILE} up postgres -d
migrate:
	docker compose -f ${COMPOSE_FILE} up migrate
gen-layer:
	go run ./tools/gen-layer/.
	go fmt ./internal
gen-mock:
	docker compose -f ${COMPOSE_FILE} up generate_mock
bdd-test:
	docker compose -f ${COMPOSE_FILE} up bdd_test --build
build:
	@echo "building..."
	go build ./cmd/srv/.
#  --gofullmethods_out=. \
# 	 --gofullmethods_opt=paths=source_relative
	