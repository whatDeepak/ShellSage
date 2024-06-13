# Define variables
GOBIN := ./bin/shellsage
GOPACKAGES := $(shell go list ./... | grep -v /vendor)

# Build command
build:
	-mkdir -p ./bin
	go build -o $(GOBIN) ./main.go

# Run command
run: build
	$(GOBIN)

# Test command (assuming you have unit tests)
test:
	go test -v $(GOPACKAGES)

# Clean command
clean:
	rm -f $(GOBIN)

# Live reload command (optional)
livereload:
	# Replace with your live reload command (e.g., go-watcher)
	# go-watcher -n -o $(GOBIN)/shellsage ./...

gomod:
	go mod tidy

vendor:
	go mod vendor

# Help command
help:
	@echo "Available commands:"
	@echo "  build     - Build the ShellSage executable"
	@echo "  run       - Run the ShellSage program"
	@echo "  test      - Run unit tests"
	@echo "  clean     - Remove built files"
	@echo "  livereload - Run with live reloading (optional)"
	@echo "  help      - Display this help message"
	@echo "  gomod     - Update go.mod file"
	@echo "  vendor    - Update vendor directory"
