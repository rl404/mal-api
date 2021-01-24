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
	@rm -f *.txt
	@rm -f *.out
	@rm -f *.html

# Generate swagger docs.
.PHONY: swagger
swagger:
	@${SWAG_GEN} -g cmd/malscraper/main.go -o ./api --parseDependency

# Build the project executable binary.
.PHONY: build
build: fmt swagger
	@cd $(CMD_PATH); \
	$(GO_BUILD) -o $(BINARY_NAME) -v

# Test and create coverage output (.html file).
# Output can be viewed on browser.
.PHONY: test
test: build
	$(GO_TEST) -cover ./... -coverprofile=$(COVER_NAME).txt
	$(GO_COVER) -html=$(COVER_NAME).txt -o $(COVER_NAME).html

# Build and run the binary.
.PHONY: run
run: build
	@cd $(CMD_PATH); \
	./$(BINARY_NAME)
