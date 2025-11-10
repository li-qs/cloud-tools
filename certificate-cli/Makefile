NAME=certificate-cli
VERSION=$(shell git rev-parse --short HEAD)
BIN_DIR=$(shell pwd)/build
MAIN_DIR=$(shell pwd)/

all: linux-amd64 darwin-amd64 windows-amd64 darwin-arm64

linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(NAME)-linux-amd64-$(VERSION) $(MAIN_DIR)

darwin-amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/$(NAME)-darwin-amd64-$(VERSION) $(MAIN_DIR)

windows-amd64:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/$(NAME)-windows-amd64-$(VERSION).exe $(MAIN_DIR)

darwin-arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o $(BIN_DIR)/$(NAME)-darwin-arm64-$(VERSION) $(MAIN_DIR)
