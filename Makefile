GO := GO111MODULE=on go
BINARY_NAME = gosrg
BINARY_PATH = /usr/local/bin/
BUILD_VERSION := v0.1.4
BUILD_TIME    := $(shell date "+%F %T")
COMMIT_SHA1   := $(shell git rev-parse HEAD )

default: build

build: main.go go.sum go.mod
	@$(GO) build -ldflags \
	"-w -s                \
	-X '$(BINARY_NAME)/config.Version=$(BUILD_VERSION)' \
	-X '$(BINARY_NAME)/config.GitCommit=$(COMMIT_SHA1)' \
	-X '$(BINARY_NAME)/config.BuildTime=$(BUILD_TIME)' \
	" \
	-o $(BINARY_NAME)
	@echo "$(BINARY_NAME) build success"
	@echo "Version: $(BUILD_VERSION)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Commit_SHA1: $(COMMIT_SHA1)"


test:
	@# 日志权限检查
	@# redis 检查
	@echo "todo"

install:
	@cp ./$(BINARY_NAME) $(BINARY_PATH)
	@if [ -f $(BINARY_NAME) ] ; then rm $(BINARY_NAME) ; fi
	@echo "$(BINARY_NAME) has installed at $(BINARY_PATH)$(BINARY_NAME)m"

uninstall:
	@if [ -f $(BINARY_PATH)$(BINARY_NAME) ] ; then rm $(BINARY_PATH)$(BINARY_NAME) ; fi
	@echo "uninstall $(BINARY_NAME) success"

start: build
	@./$(BINARY_NAME)

.PHONY: default build test install uninstall start