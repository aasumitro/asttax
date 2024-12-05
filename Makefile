# Exporting bin folder to the path for makefile
export PATH   := $(PWD)/bin:$(PATH)
# Default Shell
export SHELL  := bash
# Type of OS: Linux or Darwin.
export OSTYPE := $(shell uname -s)

# --- Tooling & Variables ----------------------------------------------------------------
include ./misc/make/tools.Makefile

install-deps: gotestsum golangci-lint
deps: $(GOTESTSUM) $(GOLANGCI)
deps:
	@ echo "Required Tools Are Available"

.Phony: run-lint
run-lint: $(GOLANGCI)
	@ echo "Applying linter"
	@ golangci-lint cache clean
	@ golangci-lint run -c .golangci.yaml ./...

.PHONY: run-tests
run-tests: $(GOTESTSUM) run-lint
	@ echo "Run tests in specific directories"
	@ gotestsum --format pkgname-and-test-fails \
		--hide-summary=skipped \
		-- -coverprofile=cover.out \
		$(shell find ./internal/repository/sql \
		./internal/repository/rest ./internal/repository/rpc \
		./internal/service ./internal/util -type d -exec go list {} \;)
	@ rm cover.out

.Phony: run-bot
run-bot: $(GOTESTSUM)
	@ go run ./cmd/bot/main.go