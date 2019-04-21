VERSION = $$(git describe --abbrev=0 --tags)
COMMIT_REV = $$(git rev-list -n 1 $(VERSION))
# SRCPATH := $(pwd)/

all: build

version:
	@echo $(VERSION)

build:
	@go build -o bin/nbaconsole *.go 

run:
	./bin/nbaconsole

clean:
	@go clean && \
	rm -rf bin/

test:
	go test ./...