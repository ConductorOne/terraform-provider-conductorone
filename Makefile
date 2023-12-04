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
	golangci-lint run

.PHONY: fmt
fmt:
	gofmt -s -w -e .

.PHONY: gen
gen:
	speakeasy generate sdk -s openapi.yaml -o . -l terraform -d

.PHONY: yaml
yaml:
	curl -sSL -o openapi.yaml https://insulator.conductor.one/api/v1/openapi.yaml
	python3 .github/workflows/clean_yaml.py openapi.yaml

.PHONY: cleangen
cleangen:
	make yaml
	make gen

.PHONY: test
test:
	go test -v -cover -timeout=120s -parallel=4 ./...

.PHONY: testacc
testacc:
	TF_ACC=1 go test -v -cover -timeout=5m ./...

.PHONY: generate
generate: fmt
	cd tools; go generate ./...

