include .envrc

MIGRATIONS_PATH=./cmd/migrate/migrations

.PHONY: migration
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DSN) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DSN) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: build
build:
	@go build -o morae ./cmd/api

.PHONY: run
run:
	@echo "Hot reload not enabled. To run the project with hot-reloading, run 'air' in root directory"
	@go run ./cmd/api 

.PHONY: test
test:
	@go test -v ./...

.PHONY: clean
clean:
	@rm -f morae
	@go clean

.PHONY: debug
debug:
	@env | grep DSN
