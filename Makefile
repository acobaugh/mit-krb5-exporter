VERSION = $(shell git describe)

.PHONY: compile
compile:
	CGO_ENABLED=0 go build -ldflags "-X main.version=${VERSION}"

.PHONY: deps
deps:
	curl -L -s https://github.com/golang/dep/releases/download/v0.5.1/dep-linux-amd64 -o ${GOPATH}/bin/dep
	chmod +x ${GOPATH}/bin/dep
	dep ensure
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: lint
lint:
	golangci-lint run
