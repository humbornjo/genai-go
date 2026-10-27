REF ?= main

.PHONY: all vendor test
all: test

vendor:
	./generate.sh "$(REF)"

test:
	go test ./...
	go vet ./...
	go mod tidy -diff
