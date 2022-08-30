.PHONY: all
all: build run

.PHONY: run
run:
	./build/server --port=3333 --src=/Users/ilya/Yandex.Disk.localized/_KnowledgeHub_

.PHONY: build
build:
	mkdir -p build/
	GOROOT=${GOROOT} GOPATH=${GOPATH} go build -o build/server cli/cli.go
