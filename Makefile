help:
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

depupdate:  ## Update all vendored dependencies
	dep ensure -update

build:  ## Build the provider
	go build -o terraform-provider-elephantsql

install: build  ## Install the provider into terraform plugin directory
	mv terraform-provider-elephantsql ~/.terraform.d/plugins/

init: install  ## Run terraform init for local testing
	terraform init

.PHONY: help build install init
.DEFAULT_GOAL := help
