.PHONY: install-tools
install-tools:
	@go mod download -modfile=golangci-lint.mod

.PHONY: lint
lint: install-tools
	@go tool -modfile=golangci-lint.mod golangci-lint run

.PHONY: lint-fix
lint-fix: install-tools
	@go tool -modfile=golangci-lint.mod golangci-lint run --fix