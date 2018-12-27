BUILDCOMMIT := $(shell git describe --dirty --always)
BUILDDATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
VER_FLAGS=-X main.commit=$(BUILDCOMMIT) -X main.date=$(BUILDDATE)

.PHONY: build
build: ## Build the operator
	@go build -ldflags "$(VER_FLAGS)" .

.PHONY: release
release: ## Build a release version of operator
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags "-w -s $(VER_FLAGS)" -o $(GOPATH)/bin/itsrandom .
