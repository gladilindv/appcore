LOCAL_BIN:=$(CURDIR)/bin

# Check global GOLANGCI-LINT
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
GOLANGCI_TAG:=1.48.0

.PHONY: build
build:
	go env
	go version
	go build -tags netgo `go list ./...`

.PHONY: test
test:
	go test -v -coverprofile .testCoverage.txt `go list ./...`
	go tool cover -func=.testCoverage.txt


.PHONY: install-lint
install-lint:
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info Downloading golangci-lint v$(GOLANGCI_TAG))
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_TAG)
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
endif

lint: install-lint
	$(info Running lint...)
	$(GOLANGCI_BIN) run --new-from-rev=origin/master --config=.cfg/lint.yaml ./...
