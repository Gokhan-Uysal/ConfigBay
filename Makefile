IS_CONTAINER ?= False

GOOS ?= darwin
GOARCH ?= amd64
COMPOSE_PROJECT_NAME ?= configbay

BASE_API_BIN = $(COMPOSE_PROJECT_NAME)-api
BASE_CLI_BIN = $(COMPOSE_PROJECT_NAME)-cli
API_BIN = $(BASE_API_BIN)-$(GOOS)-$(GOARCH)
CLI_BIN = $(BASE_CLI_BIN)-$(GOOS)-$(GOARCH)

TEST_DIR = ./test
API_TEST_DIR = $(TEST_DIR)/api_test
CLI_TEST_DIR = $(TEST_DIR)/cli_test
CORE_TEST_DIR = $(TEST_DIR)/core_test

COVERAGE_BIN = coverage.out
COVERAGE_HTML = coverage.html

install:
	@echo "Installing dependencies..."
	@go mod download && go mod verify

test-api:
	@echo "Testing..."
	@go test -v -coverprofile $(API_TEST_DIR)/$(COVERAGE_BIN) $(API_TEST_DIR)/...
	@go tool cover -html=$(API_TEST_DIR)/$(COVERAGE_BIN) -o $(API_TEST_DIR)/$(COVERAGE_HTML)
	@echo "Tested!"

test-cli:
	@echo "Testing..."
	@go test -v -cover -coverprofile $(CLI_TEST_DIR)/$(COVERAGE_BIN) $(CLI_TEST_DIR)/...
	@go tool cover -html=$(CLI_TEST_DIR)/$(COVERAGE_BIN) -o $(CLI_TEST_DIR)/$(COVERAGE_HTML)
	@echo "Tested!"

test-core:
	@echo "Testing..."
	@go test -v -cover -coverprofile $(CORE_TEST_DIR)/$(COVERAGE_BIN) $(CORE_TEST_DIR)/...
	@go tool cover -html=$(CORE_TEST_DIR)/$(COVERAGE_BIN) -o $(CORE_TEST_DIR)/$(COVERAGE_HTML)
	@echo "Tested!"

build-api:
	@echo "Building api..."
	@env CGO_ENABLED=0
	@go build -o ./$(API_BIN) -v -x ./cmd/api/.
	@echo "Built!"

build-cli:
	@echo "Building cli..."
	@env CGO_ENABLED=0
	@go build -o ./$(CLI_BIN) -v -x ./cmd/cli/.
	@echo "Built!"

run-api:
ifeq ($(IS_CONTAINER),False)
	@echo "Running $(BASE_API_BIN) dev..."
	@./$(BASE_API_BIN)* &
	@sleep 2
else
	@echo "Running $(BASE_API_BIN) prod..."
	@./$(BASE_API_BIN)*
endif

clean-api:
	@echo "Cleaning..."
	@go clean
	@rm -rf ./$(BASE_API_BIN)*
	@echo "Cleaned!"

clean-cli:
	@echo "Cleaning..."
	@go clean
	@rm -rf ./$(BASE_CLI_BIN)*
	@echo "Cleaned!"

start-api: build-api run-api

rebuild-cli: clean-cli build-cli

stop-api:
	@echo "Stopping..."
	@pkill -f $(BASE_API_BIN)
	@echo "Stopped!"

restart-api: stop-api clean-api start-api

help:
	@echo "-------------------------------------------------------------------------"
	@echo "                           CONFIGBAY MAKEFILE"
	@echo "-------------------------------------------------------------------------"
	@echo ""
	@echo "Make targets available:"
	@echo "  install         : Installs the project's dependencies."
	@echo ""
	@echo "  test-api        : Runs tests for the API and outputs a coverage report."
	@echo "  test-cli        : Runs tests for the CLI and outputs a coverage report."
	@echo "  test-core       : Runs tests for the core components and outputs a coverage report."
	@echo ""
	@echo "  build-api       : Builds the API binary for the specified GOOS and GOARCH."
	@echo "  build-cli       : Builds the CLI binary for the specified GOOS and GOARCH."
	@echo ""
	@echo "  run-api         : Runs the API binary. If IS_CONTAINER is False, it runs in dev mode. Otherwise, runs in prod mode."
	@echo ""
	@echo "  clean-api       : Cleans up generated files related to the API binary."
	@echo "  clean-cli       : Cleans up generated files related to the CLI binary."
	@echo ""
	@echo "  start-api       : Builds and then runs the API binary."
	@echo "  rebuild-cli     : Cleans, then rebuilds the CLI binary."
	@echo ""
	@echo "  stop-api        : Stops any running API binary."
	@echo "  restart-api     : Stops, cleans, builds, and restarts the API binary."
	@echo ""
	@echo "Usage:"
	@echo "  make <target> [VARIABLE=value]"
	@echo ""
	@echo "Variables available for override (with defaults):"
	@echo "  IS_CONTAINER         = $(IS_CONTAINER)"
	@echo "  GOOS                 = $(GOOS)"
	@echo "  GOARCH               = $(GOARCH)"
	@echo "  COMPOSE_PROJECT_NAME = $(COMPOSE_PROJECT_NAME)"
	@echo "-------------------------------------------------------------------------"
