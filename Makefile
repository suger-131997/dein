.PHONY: install-tools
install-tools:
	@go mod download -modfile=golangci-lint.mod

.PHONY: lint
lint: install-tools
	go tool -modfile=golangci-lint.mod golangci-lint run

.PHONY: lint-fix
lint-fix: install-tools
	go tool -modfile=golangci-lint.mod golangci-lint run --fix

.PHONY: test
test:
	go test -v ./...

.PHONY: generate
generate:
	go generate ./...
	rm ./testdata/*
	WRITE_GOLDEN_FILE_MODE=TRUE go test -v ./golden_test.go