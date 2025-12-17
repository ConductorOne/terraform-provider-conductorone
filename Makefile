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
	gofmt -s -w -e .

# vendor: Updates go.mod/go.sum and rebuilds the vendor directory.
#
# Run this after 'make gen' or any change that affects dependencies.
# The vendor directory must stay in sync with go.mod for builds to work.
.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

.PHONY: gen
gen:
	@DATE=$$(date +%Y%m%d%H%M%S); \
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
	@echo "Note: golangci-lint must be installed for 'make lint' to work."
	@echo "      Install: https://golangci-lint.run/welcome/install/"

.PHONY: generate
generate: fmt
	cd tools; go generate ./...
