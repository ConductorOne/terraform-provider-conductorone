.PHONY: run
run:
	go run main.go --debug

.PHONY: gen
gen:
	speakeasy generate sdk -s openapi.yaml -o . -l conductorone -d