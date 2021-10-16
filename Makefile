.DEFAULT_GOAL := help

.PHONY: format
format: ## format code with gofmt
	gofmt -l -w .

.PHONY: lint
lint: ## run lint
	golangci-lint run --enable=gofmt,govet,gosec,prealloc,gocognit

.PHONY: lintv
lintv: ## run lint verbosely
	golangci-lint run -v --enable=gofmt,govet,gosec,prealloc,gocognit

.PHONY: test
test: ## run test
	go test -v ./... -count=1

.PHONY: test-with-coverage
test-with-coverage: ## run test with coverage for Codecov
	go test -v ./... -count=1 -covermode=atomic -coverprofile=profile.out

.PHONY: release
release:  ## release by git tag
	git commit --allow-empty -m "Release $(VERSION):tada:"
	git tag -a $(VERSION) -m "Release $(VERSION):tada:"
	git push origin $(VERSION)

.PHONY: unrelease
unrelease:  ## unrelease by git tag
	git tag -d $(VERSION)
	git push origin :$(VERSION)

.PHONY: help
help:
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
