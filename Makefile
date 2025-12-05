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

.PHONY: gen
gen:
	@DATE=$$(date +%Y%m%d%H%M%S); \
	FILENAME="openapi_$$DATE.yaml"; \
	FILENAME2="combined_$$DATE.yaml"; \
	trap 'rm -f $$FILENAME $$FILENAME2' EXIT; \
	curl -sSL -o $$FILENAME https://insulator.conductor.one/api/v1/openapi.yaml && \
	speakeasy overlay apply -s $$FILENAME -o overlays.yml >> $$FILENAME2 && \
	speakeasy generate sdk -s $$FILENAME2 -o . -l terraform -d
	

.PHONY: test
test:
	go test -v -cover -timeout=120s -parallel=4 ./...

.PHONY: testacc
testacc:
	TF_ACC=1 go test -v -cover -timeout=5m ./...

.PHONY: generate
generate: fmt
	cd tools; go generate ./...
