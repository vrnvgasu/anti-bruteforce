BIN_APP := "./bin/antibruteforce"
DOCKER_IMG_APP="antibruteforce:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)
SOURCE_CONFIG_FILE := ./configs/app-dev.yml

.PHONY: integration-tests
integration-tests:
	SOURCE_CONFIG_FILE="./configs/app-test.yml" docker compose --project-directory ./deployments up -d
	go test -tags=integration ./tests/integration -v
	docker compose --project-directory ./deployments down

.PHONY: build
build:
	go build -v -o $(BIN_APP) -ldflags "$(LDFLAGS)" ./cmd/app

.PHONY: build-img
build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		--build-arg=SOURCE_CONFIG_FILE="$(SOURCE_CONFIG_FILE)" \
		-t $(DOCKER_IMG_APP) \
		-f build/Dockerfile_app .

.PHONY: run
run:
	SOURCE_CONFIG_FILE="$(SOURCE_CONFIG_FILE)" docker compose --project-directory ./deployments up -d

.PHONY: stop
stop:
	docker compose --project-directory ./deployments down

.PHONY:
version: build
	$(BIN_APP) version

.PHONY: test
test: generate
	go test -race ./... -v

.PHONY: install-lint-deps
install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.62.0

.PHONY: lint
lint: generate install-lint-deps
	golangci-lint run ./...

.PHONY: generate
generate: generate
	@echo "Run go:generate"
	@go generate ./...

.PHONY: generate-grpc
generate-grpc:
	protoc api/*.proto --go_out=./internal/server/grpc/pb --go-grpc_out=./internal/server/grpc/pb
	protoc api/*.proto --go_out=./tests/grpc/pb --go-grpc_out=./tests/grpc/pb