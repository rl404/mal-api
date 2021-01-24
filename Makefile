# Base Go commands.
GO_CMD   := go
GO_FMT   := $(GO_CMD) fmt
GO_CLEAN := $(GO_CMD) clean
GO_BUILD := $(GO_CMD) build -mod vendor
GO_RUN   := $(GO_CMD) run -mod vendor

# Base swagger commands.
SWAG     := swag
SWAG_GEN := ${SWAG} init

# Project executable file, and its binary.
CMD_PATH    := ./cmd/malscraper
BINARY_NAME := malscraper

# Default makefile target.
.DEFAULT_GOAL := run

# Standarize go coding style for the whole project.
.PHONY: fmt
fmt:
	@$(GO_FMT) ./...

# Lint go source code.
.PHONY: lint
lint: fmt
	@golint `go list ./... | grep -v /vendor/`

# Clean project binary, test, and coverage file.
.PHONY: clean
clean:
	@$(GO_CLEAN) ./...

# Generate swagger docs.
.PHONY: swagger
swagger:
	@${SWAG_GEN} -g cmd/malscraper/main.go -o ./api --parseDependency

# Build the project executable binary.
.PHONY: build
build: clean fmt swagger
	@cd $(CMD_PATH); \
	$(GO_BUILD) -o $(BINARY_NAME) -v

# Build and run the binary.
.PHONY: run
run: build
	@cd $(CMD_PATH); \
	./$(BINARY_NAME)

# Docker base command.
DOCKER_CMD   := docker
DOCKER_IMAGE := $(DOCKER_CMD) image
DOCKER_LOGS  := $(DOCKER_CMD) logs --follow

# Docker-compose base command and docker-compose.yml path.
COMPOSE_CMD  := docker-compose
COMPOSE_PATH := ./deployment/docker-compose.yml

# Container names.
CONTAINER_MAL := mal_api

# Build docker images and container for the project
# then delete builder image.
.PHONY: docker-build
docker-build: clean fmt swagger
	@$(COMPOSE_CMD) -f $(COMPOSE_PATH) -p $(CONTAINER_MAL) build
	@$(DOCKER_IMAGE) prune -f --filter label=stage=mal_api_builder

# Start built docker containers.
.PHONY: docker-up
docker-up:
	@$(COMPOSE_CMD) -f $(COMPOSE_PATH) -p $(CONTAINER_MAL) up -d
	@$(DOCKER_LOGS) $(CONTAINER_MAL)

# Build and start docker container for the project.
.PHONY: docker
docker: docker-build docker-up

# Stop docker container.
.PHONY: docker-stop
docker-stop:
	@$(COMPOSE_CMD) -f $(COMPOSE_PATH) -p $(CONTAINER_MAL) stop