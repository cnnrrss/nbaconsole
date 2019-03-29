VERSION = $$(git describe --abbrev=0 --tags)
COMMIT_REV = $$(git rev-list -n 1 $(VERSION))

all: build

version:
	@echo $(VERSION)

run:
	go run main.go

build:
	@go build main.go

clean:
	go clean && \
	rm -rf bin/

test:
	go test ./...