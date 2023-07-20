TF_PLUGIN_DOCS_VERSION=v0.16.0

.PHONY: run
run:
	go run main.go --debug

.PHONY: gen
gen:
	speakeasy generate sdk -s openapi.yaml -o . -l conductorone -d

.PHONY: generate
generate:
	cd tools; go generate ./...