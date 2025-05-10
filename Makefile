TEST_COVERAGE_OUT := ./.gocoverage

.PHONY: all
all:
	$(MAKE) lint
	$(MAKE) test

.PHONY: lint
lint:
	@golangci-lint run -v \
		./...

.PHONY: test-fast
test-fast:
	go test -v \
		-shuffle on \
		-failfast \
		./...

.PHONY: test
test: test-fast
	go test -v \
		-shuffle on \
		-vet=all \
		-race \
		-cover -covermode=atomic -coverprofile="${TEST_COVERAGE_OUT}" \
		./...

.PHONY: test-cover-open
test-cover-open: test
	go tool cover \
		-html="${TEST_COVERAGE_OUT}"

.PHONY: clean
clean:
	go clean -r -cache -testcache -modcache -fuzzcache
