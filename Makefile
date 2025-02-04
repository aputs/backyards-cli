# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

OS = $(shell uname)

# Project variables
PACKAGE = github.com/banzaicloud/backyards-cli
BINARY_NAME = backyards

# Helm charts
ASSETS_DIR ?= assets
CHARTS_DIR ?= ${ASSETS_DIR}/charts

# Build variables
BUILD_DIR ?= build
BUILD_PACKAGE = ${PACKAGE}/cmd/backyards
VERSION ?= $(shell git describe --tags --exact-match 2>/dev/null || git symbolic-ref -q --short HEAD)
COMMIT_HASH ?= $(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_DATE ?= $(shell date +%FT%T%z)
LDFLAGS += -X main.version=${VERSION} -X main.commitHash=${COMMIT_HASH} -X main.buildDate=${BUILD_DATE}
export CGO_ENABLED ?= 0
ifeq (${VERBOSE}, 1)
ifeq ($(filter -v,${GOARGS}),)
	GOARGS += -v
endif
TEST_FORMAT = short-verbose
endif

# Dependency versions
GOTESTSUM_VERSION = 0.3.3
GOLANGCI_VERSION = 1.18.0
LICENSEI_VERSION = 0.1.0
GORELEASER_VERSION = 0.112.2
PACKR_VERSION = 2.0.8

GOLANG_VERSION = 1.13

# Add the ability to override some variables
# Use with care
-include override.mk

bin/packr2: bin/packr2-${PACKR_VERSION}
	@ln -sf packr2-${PACKR_VERSION} bin/packr2
bin/packr2-${PACKR_VERSION}:
	@mkdir -p bin
ifeq (${OS}, Darwin)
	curl -L https://github.com/gobuffalo/packr/releases/download/v${PACKR_VERSION}/packr_${PACKR_VERSION}_darwin_amd64.tar.gz | tar -zOxf - packr2 > ./bin/packr2-${PACKR_VERSION} && chmod +x ./bin/packr2-${PACKR_VERSION}
endif
ifeq (${OS}, Linux)
	curl -L https://github.com/gobuffalo/packr/releases/download/v${PACKR_VERSION}/packr_${PACKR_VERSION}_linux_amd64.tar.gz | tar -zOxf - packr2 > ./bin/packr2-${PACKR_VERSION} && chmod +x ./bin/packr2-${PACKR_VERSION}
endif

.PHONY: pre-build
pre-build: ## Pre build bundles of static assets
	find ${ASSETS_DIR}/ -exec touch -t 201901010101 {} +
	go generate ${BUILD_PACKAGE}

.PHONY: build
build: pre-build ## Build a binary
ifeq (${VERBOSE}, 1)
	go env
endif
ifneq (${IGNORE_GOLANG_VERSION_REQ}, 1)
	@printf "${GOLANG_VERSION}\n$$(go version | awk '{sub(/^go/, "", $$3);print $$3}')" | sort -t '.' -k 1,1 -k 2,2 -k 3,3 -g | head -1 | grep -q -E "^${GOLANG_VERSION}$$" || (printf "Required Go version is ${GOLANG_VERSION}\nInstalled: `go version`" && exit 1)
endif
	go build ${GOARGS} -tags "${GOTAGS}" -ldflags "${LDFLAGS}" -o ${BUILD_DIR}/${BINARY_NAME} ${BUILD_PACKAGE}

.PHONY: build-release
build-release: LDFLAGS += -w
build-release: build ## Build a binary without debug information

.PHONY: build-debug
build-debug: ## Build a binary with remote debugging capabilities
	@${MAKE} GOARGS="${GOARGS} -gcflags \"all=-N -l\"" BUILD_DIR="${BUILD_DIR}/debug" build

.PHONY: debug
debug: build-debug
	dlv --listen=:40000 --log --headless=true --api-version=2 exec "${BUILD_DIR}/debug/${BINARY_NAME}" -- ${ARGS}

.PHONY: run
run: build
	"${BUILD_DIR}/${BINARY_NAME}" ${ARGS}


.PHONY: generate-docs
generate-docs: ## Generate documentation for Backyards CLI
	rm -rf cmd/docs/*.md
	cd cmd/docs/ && go run -v generate.go

.PHONY: check
check: test lint ## Run tests and linters

bin/gotestsum: bin/gotestsum-${GOTESTSUM_VERSION}
	@ln -sf gotestsum-${GOTESTSUM_VERSION} bin/gotestsum
bin/gotestsum-${GOTESTSUM_VERSION}:
	@mkdir -p bin
ifeq (${OS}, Darwin)
	curl -L https://github.com/gotestyourself/gotestsum/releases/download/v${GOTESTSUM_VERSION}/gotestsum_${GOTESTSUM_VERSION}_darwin_amd64.tar.gz | tar -zOxf - gotestsum > ./bin/gotestsum-${GOTESTSUM_VERSION} && chmod +x ./bin/gotestsum-${GOTESTSUM_VERSION}
endif
ifeq (${OS}, Linux)
	curl -L https://github.com/gotestyourself/gotestsum/releases/download/v${GOTESTSUM_VERSION}/gotestsum_${GOTESTSUM_VERSION}_linux_amd64.tar.gz | tar -zOxf - gotestsum > ./bin/gotestsum-${GOTESTSUM_VERSION} && chmod +x ./bin/gotestsum-${GOTESTSUM_VERSION}
endif

TEST_PKGS ?= ./...
TEST_REPORT_NAME ?= results.xml
.PHONY: test
test: TEST_REPORT ?= main
test: TEST_FORMAT ?= short
test: SHELL = /bin/bash
test: bin/gotestsum ## Run tests
	@mkdir -p ${BUILD_DIR}/test_results/${TEST_REPORT}
	bin/gotestsum --no-summary=skipped --junitfile ${BUILD_DIR}/test_results/${TEST_REPORT}/${TEST_REPORT_NAME} --format ${TEST_FORMAT} -- $(filter-out -v,${GOARGS}) $(if ${TEST_PKGS},${TEST_PKGS},./...)

bin/golangci-lint: bin/golangci-lint-${GOLANGCI_VERSION}
	@ln -sf golangci-lint-${GOLANGCI_VERSION} bin/golangci-lint
bin/golangci-lint-${GOLANGCI_VERSION}:
	@mkdir -p bin
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b ./bin/ v${GOLANGCI_VERSION}
	@mv bin/golangci-lint $@

.PHONY: lint
lint: bin/golangci-lint ## Run linter
	bin/golangci-lint run

.PHONY: lint-fix
lint-fix: bin/golangci-lint ## Run linter
	bin/golangci-lint run --fix

bin/licensei: bin/licensei-${LICENSEI_VERSION}
	@ln -sf licensei-${LICENSEI_VERSION} bin/licensei
bin/licensei-${LICENSEI_VERSION}:
	@mkdir -p bin
	curl -sfL https://raw.githubusercontent.com/goph/licensei/master/install.sh | bash -s v${LICENSEI_VERSION}
	@mv bin/licensei $@

.PHONY: license-check
license-check: bin/licensei ## Run license check
	bin/licensei check
	./scripts/check-header.sh

.PHONY: license-cache
license-cache: bin/licensei ## Generate license cache
	bin/licensei cache

bin/goreleaser: bin/goreleaser-${GORELEASER_VERSION}
	@ln -sf goreleaser-${GORELEASER_VERSION} bin/goreleaser
bin/goreleaser-${GORELEASER_VERSION}:
	@mkdir -p bin
	curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | bash -s -- -b ./bin/ v${GORELEASER_VERSION}
	@mv bin/goreleaser $@

.PHONY: release
release: bin/goreleaser # Publish a release
	bin/goreleaser release ${GORELEASERFLAGS}

release-%:
ifneq (${DRY}, 1)

ifeq (${TAG}, 1)
	git tag -m 'Release $*' ${TAG_PREFIX}$*
ifeq (${PUSH}, 1)
	git push; git push origin ${TAG_PREFIX}$*
endif
endif
endif
	@echo "Version updated to $*!"
ifneq (${PUSH}, 1)
	@echo
	@echo "Review the changes made by this script then execute the following:"
ifneq (${TAG}, 1)
	@echo
	@echo "git tag -m 'Release $*' ${TAG_PREFIX}$*"
	@echo
	@echo "Finally, push the changes:"
endif
	@echo
	@echo "git push; git push origin ${TAG_PREFIX}$*"
endif

.PHONY: patch
patch: ## Release a new patch version
	@${MAKE} release-$(shell (git describe --abbrev=0 --tags 2> /dev/null || echo "0.0.0") | sed 's/^v//' | awk -F'[ .]' '{print $$1"."$$2"."$$3+1}')

.PHONY: minor
minor: ## Release a new minor version
	@${MAKE} release-$(shell (git describe --abbrev=0 --tags 2> /dev/null || echo "0.0.0") | sed 's/^v//' | awk -F'[ .]' '{print $$1"."$$2+1".0"}')

.PHONY: major
major: ## Release a new major version
	@${MAKE} release-$(shell (git describe --abbrev=0 --tags 2> /dev/null || echo "0.0.0") | sed 's/^v//' | awk -F'[ .]' '{print $$1+1".0.0"}')

.PHONY: list
list: ## List all make targets
	@${MAKE} -pRrn : -f $(MAKEFILE_LIST) 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | sort

.PHONY: download-charts
download-charts: ## Download Helm charts
	scripts/download-charts.sh ${CHARTS_DIR}

.PHONY: check-charts-diff
check-charts-diff: ## Check whether the local charts differs
	scripts/check-charts-diff.sh

.PHONY: check-docs-diff
check-docs-diff: ## Check whether the docs differ
	scripts/check-docs-diff.sh

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Variable outputting/exporting rules
var-%: ; @echo $($*)
varexport-%: ; @echo $*=$($*)
