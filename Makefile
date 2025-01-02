BIN_APP := "./bin/antibruteforce"
DOCKER_IMG_APP="antibruteforce:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

.PHONY: integration-tests
integration-tests:
	echo "integration-tests"

.PHONY: build
build:
	go build -v -o $(BIN_APP) -ldflags "$(LDFLAGS)" ./cmd/app

.PHONY: build-img
build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG_APP) \
		-f build/Dockerfile_app .

.PHONY: build-img
run-img: build-img
	docker run $(DOCKER_IMG_APP)

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