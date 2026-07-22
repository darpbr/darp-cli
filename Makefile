.PHONY: help build install test lint fmt verify package clean run cross-compile

GO ?= go
BINARY_NAME ?= darp
BINARY_DIR ?= dist
INSTALL_DIR ?= $(HOME)/.local/bin
VERSION_BASE ?= 0.1.0
GIT_TAG := $(shell git describe --tags --exact-match 2>/dev/null)
GIT_SHA := $(shell git rev-parse --short HEAD 2>/dev/null || echo unknown)
GIT_DIRTY := $(shell test -n "$$(git status --porcelain 2>/dev/null)" && echo .dirty)
VERSION ?= $(if $(GIT_TAG),$(patsubst v%,%,$(GIT_TAG)),$(VERSION_BASE)-dev+$(GIT_SHA)$(GIT_DIRTY))
LDFLAGS := -X main.version=$(VERSION)
PLATFORMS ?= linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64 windows/arm64

help:
	@echo "Usage:"
	@echo "  make build           Build the CLI binary for the current platform"
	@echo "  make install         Install the CLI binary to $(INSTALL_DIR)"
	@echo "  make version         Print the computed CLI version"
	@echo "  make test            Run Go tests"
	@echo "  make lint            Run linting (golangci-lint if available, otherwise go vet)"
	@echo "  make fmt             Format Go source files"
	@echo "  make package         Build and package the current-platform binary"
	@echo "  make cross-compile   Build binaries for Linux, macOS and Windows"
	@echo "  make verify          Run fmt, test and lint"
	@echo "  make clean           Remove generated artifacts"
	@echo "  make run             Run the CLI locally"

version:
	@echo $(VERSION)

build:
	@mkdir -p $(BINARY_DIR)
	$(GO) build -trimpath -ldflags "$(LDFLAGS)" -o $(BINARY_DIR)/$(BINARY_NAME) .

install:
	@mkdir -p $(INSTALL_DIR)
	$(GO) build -trimpath -ldflags "$(LDFLAGS)" -o $(INSTALL_DIR)/$(BINARY_NAME) .
	@echo "Installed $(BINARY_NAME) $(VERSION) to $(INSTALL_DIR)/$(BINARY_NAME)"

run:
	$(GO) run -ldflags "$(LDFLAGS)" .

test:
	$(GO) test ./...

fmt:
	$(GO)fmt -w $(shell find . -name '*.go' -not -path './vendor/*')

lint:
	@command -v golangci-lint >/dev/null 2>&1 && golangci-lint run || (echo "golangci-lint not installed; falling back to go vet" && $(GO) vet ./...)

package: build
	@mkdir -p $(BINARY_DIR)/package
	@tar -czf $(BINARY_DIR)/package/$(BINARY_NAME)-$(VERSION)-linux-amd64.tar.gz -C $(BINARY_DIR) $(BINARY_NAME)
	@echo "Built package: $(BINARY_DIR)/package/$(BINARY_NAME)-$(VERSION)-linux-amd64.tar.gz"

cross-compile:
	@mkdir -p $(BINARY_DIR)/cross
	@for pair in $(PLATFORMS); do \
		os=$${pair%/*}; \
		arch=$${pair#*/}; \
		out="$(BINARY_DIR)/cross/$(BINARY_NAME)-$${os}-$${arch}"; \
		if [ "$$os" = "windows" ]; then out="$$out.exe"; fi; \
		echo "Building $$os/$$arch -> $$out"; \
		GOOS=$$os GOARCH=$$arch $(GO) build -trimpath -ldflags "$(LDFLAGS)" -o "$$out" .; \
	done

verify: fmt test lint

clean:
	rm -rf $(BINARY_DIR)
