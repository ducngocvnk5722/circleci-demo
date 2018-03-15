GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

test: test-all

test-all:
	@go test -timeout 20m -v $(GOPACKAGES)
test-cover-html:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(GOPACKAGES),\
		go test -coverprofile=coverage.out -covermode=count -v -p 1 $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out -o cover.html
	go tool cover -func=coverage-all.out >> func-cover.txt