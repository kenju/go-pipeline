NAME := pipeline
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
	-X 'main.revision=$(REVISION)'

## Run tests
test:
	gotest ./... -v -parallel=4

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
	go get github.com/kenju/go-pipeline # for CI

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
