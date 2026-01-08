.PHONY: run
run:
	go run main.go --debug

.PHONY: build
build:
	go build -v ./...

.PHONY: install
install: build
	go install -v ./...

.PHONY: lint
lint:
	golangci-lint run --timeout=3m

.PHONY: fmt
fmt:
	find . -name '*.go' -not -path './vendor/*' | xargs gofmt -s -w -e

# vendor: Updates go.mod/go.sum and rebuilds the vendor directory.
#
# Run this after 'make gen' or any change that affects dependencies.
# The vendor directory must stay in sync with go.mod for builds to work.
.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

# gen: Regenerates the SDK from the live OpenAPI spec.
#
# IMPORTANT: Speakeasy version is pinned in .speakeasy/workflow.yaml.
# This target enforces the version check before generation. See KNOWN_ISSUES.md
# for details on the perpetual diff regression in newer Speakeasy versions.
#
# After regeneration, verify the pagination fix is intact (see CLAUDE.md).
.PHONY: gen
gen:
	@PINNED_VERSION=$$(grep 'speakeasyVersion:' .speakeasy/workflow.yaml | head -1 | awk '{print $$2}'); \
	INSTALLED_VERSION=$$(speakeasy --version 2>/dev/null | grep -oE '[0-9]+\.[0-9]+\.[0-9]+' | head -1); \
	if [ -z "$$INSTALLED_VERSION" ]; then \
		echo "ERROR: speakeasy not found. Install from https://speakeasy.com/docs/speakeasy-cli/getting-started"; \
		exit 1; \
	fi; \
	if [ "$$INSTALLED_VERSION" != "$$PINNED_VERSION" ]; then \
		echo "ERROR: Speakeasy version mismatch"; \
		echo "  Installed: $$INSTALLED_VERSION"; \
		echo "  Required:  $$PINNED_VERSION"; \
		echo ""; \
		echo "See KNOWN_ISSUES.md for why version matters."; \
		echo "Either install the correct version or use 'speakeasy run' which respects workflow.yaml"; \
		exit 1; \
	fi; \
	DATE=$$(date +%Y%m%d%H%M%S); \
	FILENAME="openapi_$$DATE.yaml"; \
	FILENAME2="combined_$$DATE.yaml"; \
	trap 'rm -f $$FILENAME $$FILENAME2' EXIT; \
	curl -sSL -o $$FILENAME https://insulator.conductor.one/api/v1/openapi.yaml && \
	speakeasy overlay apply -s $$FILENAME -o overlay.yaml >> $$FILENAME2 && \
	speakeasy generate sdk -s $$FILENAME2 -o . -l terraform -d
	$(MAKE) vendor
	

.PHONY: test
test:
	go test -v -cover -timeout=120s -parallel=4 ./...

# pre-commit: Validates the codebase before committing changes.
#
# This target runs the essential validation checks that should pass before
# any code is committed, particularly after running 'make gen' to regenerate
# from the OpenAPI spec.
#
# Checks performed (in order):
#   1. build    - Ensures all generated code compiles successfully
#   2. lint     - Runs golangci-lint for static analysis and code quality
#   3. test     - Runs unit tests (not acceptance tests, which need credentials)
#   4. generate - Regenerates tfplugindocs to ensure docs match the provider schema
#
# After completion, review 'git diff' to see any changes made by 'make generate'.
# If docs were updated, stage and include them in your commit.
#
# Note: This does NOT run 'make testacc' (acceptance tests) as those require
# CONDUCTORONE_CLIENT_ID, CONDUCTORONE_CLIENT_SECRET, and CONDUCTORONE_SERVER_URL
# environment variables to be configured.
.PHONY: pre-commit
pre-commit: build lint test generate

.PHONY: testacc
testacc:
	TF_ACC=1 go test -v -cover -timeout=5m ./...

# install-hooks: Configures git to use the project's hooks from .githooks/
#
# This sets the git hooksPath to use the hooks in .githooks/, which includes:
#   - pre-push: Runs 'make pre-commit' before pushing to remote
#
# Run this once after cloning the repository to enable the hooks.
# To bypass hooks temporarily, use --no-verify (e.g., 'git push --no-verify').
.PHONY: install-hooks
install-hooks:
	git config core.hooksPath .githooks
	chmod +x .githooks/*
	@echo "Git hooks installed. Pre-push validation is now enabled."
	@echo "Note: golangci-lint v1.x (>= v1.63.0) is required for 'make lint' to work."
	@echo "      Install: go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8"

.PHONY: generate
generate: fmt
	cd tools; go generate ./...
