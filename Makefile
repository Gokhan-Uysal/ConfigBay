IS_CONTAINER ?= False

GOOS ?= linux
GOARCH ?= amd64

APP_NAME = configbay
BASE_API_BIN = $(APP_NAME)-api
BASE_CLI_BIN = $(APP_NAME)-cli
API_BIN = $(BASE_API_BIN)-$(GOOS)-$(GOARCH)
CLI_BIN = $(BASE_CLI_BIN)-$(GOOS)-$(GOARCH)

install:
	@echo "Installing dependencies..."
	@go mod download && go mod verify

test:
	@echo "Testing..."
	@go test -v -cover ${PWD}/test
	@echo "Tested!"

build-api:
	@echo "Building api..."
	@env CGO_ENABLED=0
	@go build -o ${PWD}/$(API_BIN) -v -x ${PWD}/cmd/api/.
	@echo "Built!"

build-cli:
	@echo "Building cli..."
	@env CGO_ENABLED=0
	@go build -o ${PWD}/$(CLI_BIN) -v -x ${PWD}/cmd/cli/.
	@echo "Built!"

run-api:
ifeq ($(IS_CONTAINER),False)
	@echo "Running $(API_BIN) dev..."
	@./$(API_BIN) &
	@sleep 2
else
	@echo "Running $(API_BIN) prod..."
	@./$(API_BIN)
endif

clean-api:
	@echo "Cleaning..."
	@go clean
	@rm -rf ${PWD}/$(BASE_API_BIN)*
	@echo "Cleaned!"

clean-cli:
	@echo "Cleaning..."
	@go clean
	@rm -rf ${PWD}/$(BASE_CLI_BIN)*
	@echo "Cleaned!"

start-api: build-api run-api

rebuild-cli: clean-cli build-cli

stop-api:
	@echo "Stopping..."
	@pkill -f $(API_BIN)
	@echo "Stopped!"

restart-api: stop-api clean-api start-api