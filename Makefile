# Make this makefile self-documented with target `help`
.PHONY: help
.DEFAULT_GOAL := help
help: ## Show help
	@grep -Eh '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: download
download: ## Download go modules
	go mod download

.PHONY: lint
lint: ## Lint the application
	golangci-lint run ./...

.PHONY: lint-fix
lint-fix: ## Lint and auto-fix the application
	golangci-lint run --fix ./...

.PHONY: generate
generate: ## Update generated files
	go generate ./...

.PHONY: build
build: ## Build the application
	go build -o the-ethereum-fetcher

.PHONY: test
test: ## Test the application
	go test ./...

.PHONY: acquirer
acquirer: ## Run acquirer (env: theta)
	go run . acquirer
