GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

test: test-all

test-all:
	@go test -timeout 20m -v $(GOPACKAGES)