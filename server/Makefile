all: build run

run:
	./build/server --port=3000 --src=/Users/ilya/Yandex.Disk.localized/_KnowledgeHub_

build:
	mkdir -p build/
	GOROOT=${GOROOT} GOPATH=${GOPATH} go build -o build/server
