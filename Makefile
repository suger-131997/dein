.PHONY: install-tools
install-tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(GOPATH)/bin v1.64.8

.PHONY: lint
lint:
	golangci-lint run

.PHONY: lint-fix
lint-fix:
	golangci-lint --fix

.PHONY: test
test:
	go test -v ./...

.PHONY: generate
generate:
	go generate ./...
	rm ./testdata/*
	WRITE_GOLDEN_FILE_MODE=TRUE go test -v ./golden_test.go