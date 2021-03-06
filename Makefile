# Get current directory
current_dir := $(shell pwd)

# Get current commit hash
# commit_hash := $(shell git rev-parse --short=7 HEAD)

# Targets
.PHONY: test

all: test build

test:
	@echo "Running all tests"
	go clean -testcache
	go test -v -race github.com/istherepie/dropoff/internal/webserver

build:
	@echo "Building binaries"

	mkdir $(current_dir)/build
	go build -o $(current_dir)/build/droppy $(current_dir)/cmd/main.go

clean:
	@echo "Cleaning up..."
	rm -rf $(current_dir)/build