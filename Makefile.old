SHELL := /bin/bash

.PHONY: all
all: \
	commitlint \
	buf-lint \
	buf-generate \
	go-lint \
	go-review \
	go-test \
	go-mod-tidy \
	git-verify-nodiff

include tools/buf/rules.mk
include tools/commitlint/rules.mk
include tools/git-verify-nodiff/rules.mk
include tools/golangci-lint/rules.mk
include tools/goreview/rules.mk
include tools/semantic-release/rules.mk

build/protoc-gen-go: go.mod
	$(info [$@] building binary...)
	@go build -o $@ google.golang.org/protobuf/cmd/protoc-gen-go

.PHONY: clean
clean:
	$(info [$@] removing build files...)
	@rm -rf build

.PHONY: go-test
go-test:
	$(info [$@] running Go tests...)
	@mkdir -p build/coverage
	@go test -cover -race -coverprofile=build/coverage/$@.txt -covermode=atomic ./...

.PHONY: go-mod-tidy
go-mod-tidy:
	$(info [$@] tidying Go module files...)
	@go mod tidy -v

.PHONY: buf-lint
buf-lint: $(buf)
	$(info [$@] linting protobuf schemas...)
	@cd internal/examples/proto && $(buf) lint

.PHONY: buf-generate
buf-generate: $(buf) build/protoc-gen-go
	$(info [$@] generating protobuf stubs...)
	@rm -rf internal/examples/proto/gen
	@cd internal/examples/proto && $(buf) generate --path einride
