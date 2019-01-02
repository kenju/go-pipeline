NAME := pipeline
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
	-X 'main.revision=$(REVISION)'

OK_COLOR    = \033[0;32m
ERROR_COLOR = \033[0;31m
WARN_COLOR  = \033[0;33m
NO_COLOR    = \033[m

OK_STRING    = "[OK]"
ERROR_STRING = "[ERROR]"
WARN_STRING  = "[WARNING]"

## Run tests
test:
	if gotest ./... -v; then \
		echo "$(OK_COLOR)$(OK_STRING) go test succeeded$(NO_COLOR)"; \
	else \
		echo "$(ERROR_COLOR)$(ERROR_STRING) go test failed$(NO_COLOR)n"; \
	fi

## Build binaries
build:
	go build -ldflags "$(LDFLAGS)"

## Setup
setup:
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/Songmu/make2help/cmd/make2help
	go get github.com/rakyll/gotest
	go get github.com/cheekybits/genny

## Lint
lint:
	go vet $$(go list)
	for pkg in $$(go list); do \
		golint -set_exit_status $$pkg || exit $$?; \
	done

## Format source codes
fmt:
	goimports -w main.go

## Generate files with `go generate`
gen:
	go generate

## Update CHANGELOG.md with auto-changelog
changelog:
	auto-changelog -t keepachangelog && git commit -am "update CHANGELOG"

## Show help
help:
	@make2help $(MAKEFILE_LIST)
