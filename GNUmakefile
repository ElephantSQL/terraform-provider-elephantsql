GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
PKG_NAME=elephantsql
PROVIDER_VERSION = 0.0.1

default: build

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	GOOS += linux
else ifeq ($(UNAME_S),Darwin)
	GOOS += darwin
endif

UNAME_M := $(shell uname -m)
ifeq ($(UNAME_M),x86_64)
	GOARCH += amd64
else ifeq ($(UNAME_M),arm64)
	GOARCH += arm64
else ifeq ($(UNAME_M),aarch64)
	GOARCH += arm64
endif
PROVIDER_ARCH = $(GOOS)_$(GOARCH)

tools:
	GO111MODULE=on go install github.com/client9/misspell/cmd/misspell
	GO111MODULE=on go install github.com/golangci/golangci-lint/cmd/golangci-lint

build: fmtcheck
	go install -ldflags "-X 'main.version=$(PROVIDER_VERSION)'"

local-clean:
	rm -rf ~/.terraform.d/plugins/localhost/elephantsql/elephantsql/$(PROVIDER_VERSION)/$(PROVIDER_ARCH)/terraform-provider-elephantsql_v$(PROVIDER_VERSION)

local-build: local-clean
	@echo $(GOOS);
	@echo $(GOARCH);
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X 'main.version=$(PROVIDER_VERSION)'" -o terraform-provider-elephantsql_v$(PROVIDER_VERSION)

local-install: local-build
	mkdir -p ~/.terraform.d/plugins/localhost/elephantsql/elephantsql/$(PROVIDER_VERSION)/$(PROVIDER_ARCH)
	cp $(CURDIR)/terraform-provider-elephantsql_v$(PROVIDER_VERSION) ~/.terraform.d/plugins/localhost/elephantsql/elephantsql/$(PROVIDER_VERSION)/$(PROVIDER_ARCH)

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

lint:
	@echo "==> Checking source code against linters..."
	golangci-lint run ./...

.PHONY: fmtcheck lint tools
