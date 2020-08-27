export GO111MODULE=on

UNAME := $(shell uname)
BINARY_NAME = drone_simulation
DIST_DIR = $(CURDIR)/dist
DIST_BIN = $(DIST_DIR)/$(BINARY_NAME)
CMD_FILE= $(CURDIR)/cmd/$(BINARY_NAME)/*.go
DEBUG_PORT=40000

.PHONY: help
help:	## Show a list of available commands
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: make-debug
make-debug:	## Debug Makefile itself
	@echo $(UNAME)

.PHONY: fmt
fmt:	## Format code
	gofmt -w .
	goimports -w -l ./

bench:	## Run benchmarks
	$(GOCMD) test -bench=. -benchmem ./...

.PHONY: tidy
tidy:	## Prune any no-longer-needed dependencies from go.mod and add any dependencies needed
	go mod tidy -v

.PHONY: test
test:	## Run unitary test
	go test -p 1 -cover -v ./... -timeout 5m

.PHONY: build
build:	## Build application
	CGO_ENABLED=0 go build -o ${DIST_BIN} -ldflags="-s -w" $(CMD_FILE)

.PHONY: build-debug
build-debug: ## Build application for debug proposes
	go build -o ${DIST_BIN} -gcflags '-N' $(CMD_FILE)

.PHONY: run
run:	## Run application
	go run $(CMD_FILE)

.PHONY: download-tools
download-tools:  ## Download all required tools to generate documentation, code analysis...
	GO111MODULE=off go get -u golang.org/x/lint/golint
	GO111MODULE=off go get -u github.com/fzipp/gocyclo
	GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger
	GO111MODULE=off go get -u github.com/go-openapi/runtime
	GO111MODULE=off go get -u github.com/jessevdk/go-flags

.PHONY: remote-debug
remote-debug:	## Debug application application [WIP]
	dlv debug $(CMD_FILE) --headless --listen=:$(DEBUG_PORT) --api-version=2 --log
