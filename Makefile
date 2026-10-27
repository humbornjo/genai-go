.PHONY: all vendor update test
all: test

vendor:
	./generate.sh

update:
	git submodule update --init --remote upstream
	./generate.sh

test:
	go test ./...
	go vet ./...
	go mod tidy -diff
