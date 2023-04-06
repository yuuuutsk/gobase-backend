SHELL=/bin/bash
PROJECTNAME=gobase-backend

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

build:
	go build -gcflags "all=-N -l" -o ./bin/cli ./cmd/cli
	go build -gcflags "all=-N -l" -o ./bin/server ./cmd/graphql

lint:
	goimports -w .
	go vet ./...

.PHONY: wire
wire: ## wire
	wire ./cmd/di

.PHONY: install-tools
install-tools: ## install-tools
	go install github.com/volatiletech/sqlboiler/v4@v4.14.2
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
	go install golang.org/x/tools/cmd/goimports@v0.1.3
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.41.1
	go install github.com/google/wire/cmd/wire@v0.5.0

db-migrate-all: ## db-migrate-all
	make db-migrate
	make test-db-migrate

db-migrate-dryrun-all: ## db-migrate-dryrun-all
	make db-migrate-dryrun
	make test-db-migrate-dryrun

# gobase-backend migrate with local mysqldef
db-migrate:
	 cat ./dbschema/*.sql | mysqldef -u testuser -p password -h 127.0.0.1 gobase-backend_local
db-migrate-dryrun:
	 cat ./dbschema/*.sql | mysqldef -u testuser -p password -h 127.0.0.1 --dry-run gobase-backend_local

prd-db-migrate:
	 cat ./dbschema/*.sql | mysqldef -u ${DB_USER} -p ${DB_PASSWORD} -P ${DB_PORT} -h ${DB_HOST} ${DB_NAME}
prd-db-migrate-dryrun:
	 cat ./dbschema/*.sql | mysqldef -u ${DB_USER} -p ${DB_PASSWORD} -P ${DB_PORT} -h ${DB_HOST} --dry-run ${DB_NAME}


# gobase-backend migrate with local mysqldef
test-db-migrate:
	 cat ./dbschema/*.sql | mysqldef -u testuser -p password -h 127.0.0.1 gobase-backend_test
test-db-migrate-dryrun:
	 cat ./dbschema/*.sql | mysqldef -u testuser -p password -h 127.0.0.1 --dry-run gobase-backend_test

.PHONY: test
test: ## test
	make test-db-migrate-all
	go test -v ./...

.PHONY: gen
gen: ## gen
	sqlboiler mysql -o app/infrastracture/dao -p dao --no-driver-templates --wipe --templates ${GOPATH}/pkg/mod/github.com/volatiletech/sqlboiler/v4@v4.14.2/templates/main --templates templates/sqlboiler/main
	sqlboiler mysql -o app/infrastracture/dto -p models --no-driver-templates --wipe --templates templates/sqlboiler/models
	rm -fr app/models/*.base.go
	goimports -w app/infrastracture/dao
	goimports -w app/infrastracture/dto
	mv app/infrastracture/dto/*.base.go app/domain/models
	mv -n app/infrastracture/dto/*.go app/domain/models

.PHONY: gqlgen
gqlgen: ## gqlgen
	go generate ./graph/...
