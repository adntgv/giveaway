GOMOD= GO111MOD=on

TARGET_SERVER= ./build/server
PATH_SERVER= ./cmd/server


TARGET_CLI= ./build/cli
PATH_CLI= ./cmd/cli

all: build

build: build-cli build-server

build-cli:
	${GOMOD} go build -o ${TARGET_CLI} ${PATH_CLI}

build-server:
	${GOMOD} go build -o ${TARGET_SERVER} ${PATH_SERVER}

clear:
	rm -rf ./build

PHONY: all