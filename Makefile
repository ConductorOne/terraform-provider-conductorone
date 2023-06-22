.PHONY: gen
gen:
	speakeasy generate sdk -s openapi.yaml -o . -l terraform -d