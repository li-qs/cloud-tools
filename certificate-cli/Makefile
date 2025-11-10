APP_NAME = certificate-cli
SRC_FILES = main.go
BUILD_DIR = build
PLATFORMS = linux/amd64 linux/arm64 windows/amd64 darwin/amd64 darwin/arm64

GO=go

.PHONY: all
all: clean build

.PHONY: build
build:
	@mkdir -p $(BUILD_DIR)
	@for platform in $(PLATFORMS); do \
		OS=$$(echo $$platform | cut -d/ -f1); \
		ARCH=$$(echo $$platform | cut -d/ -f2); \
		OUTPUT=$(BUILD_DIR)/$(APP_NAME)_$${OS}_$${ARCH}; \
		[ "$$OS" = "windows" ] && OUTPUT=$$OUTPUT.exe; \
		echo "Building for $$OS/$$ARCH..."; \
		GOOS=$$OS GOARCH=$$ARCH $(GO) build -o $$OUTPUT $(SRC_FILES) || exit 1; \
		echo "Built $$OUTPUT"; \
	done

.PHONY: clean
clean:
	@echo "Cleaning up build directory..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete."

.PHONY: run
run: build
	./$(BUILD_DIR)/$(APP_NAME)
