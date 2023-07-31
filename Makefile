TARGET_BIN = go-api-template
TARGET_ARCH = amd64
SOURCE_MAIN = cmd/app/main.go
LDFLAGS = -s -w

export GOOSE_DBSTRING=postgresql://admin:admin@localhost:5432/postgres
export GOOSE_DRIVER=postgres

all: build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=$(TARGET_ARCH) go build -ldflags "$(LDFLAGS)" -o bin/$(TARGET_BIN)_linux-amd64 $(SOURCE_MAIN)

build-linux-noarch:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "$(LDFLAGS)" -o bin/$(TARGET_BIN) $(SOURCE_MAIN)

start:
	go run $(SOURCE_MAIN)

install-dependencies:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.4
	go install github.com/pressly/goose/v3/cmd/goose@latest


.PHONY: run
mod-download:
	go mod download

generate: install-dependencies mod-download
	go generate ./...

generate-mocks:
	@mockery --output user/mocks --dir user --all

.PHONY: build
build: ## Build app
	go build -o bin/app cmd/app/main.go


migrate: ## run database migrations
	goose -dir migrations up


migrate-rollback: ## run database migrations
	goose -dir migrations down
