GO := GO111MODULE=on go
BINARY_NAME = gosrg
BINARY_PATH = /usr/local/bin/

default: build

build: main.go go.sum go.mod
	@$(GO) build -ldflags "-s -w" -o $(BINARY_NAME)
	@echo "$(BINARY_NAME) build success"

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