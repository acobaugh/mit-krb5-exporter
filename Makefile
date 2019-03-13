VERSION = $(shell git describe)

.PHONY: compile
compile:
	CGO_ENABLED=0 go build -ldflags "-X main.version=${VERSION}"

.PHONY: deps
deps:
	dep ensure
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: lint
lint:
	golangci-lint run
