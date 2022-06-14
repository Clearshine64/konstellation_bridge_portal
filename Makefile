VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
BINDIR ?= $(GOPATH)/bin
BUILDDIR ?= $(CURDIR)/build
SIMAPP = ./app

# for dockerized protobuf tools
DOCKER := $(shell which docker)
PROTO_CONTAINER := cosmwasm/prototools-docker:v0.1.0
BUF_IMAGE=bufbuild/buf@sha256:9dc5d6645f8f8a2d5aaafc8957fbbb5ea64eada98a84cb09654e8f49d6f73b3e
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(BUF_IMAGE)
#DOCKER_BUF := docker run --rm -v $(shell pwd)/buf.yaml:/workspace/buf.yaml -v $(shell go list -f "{{ .Dir }}" -m github.com/cosmos/cosmos-sdk):/workspace/cosmos_sdk_dir -v $(shell pwd):/workspace/wasmd  --workdir /workspace $(PROTO_CONTAINER)
HTTPS_GIT := https://github.com/konstellation/swap.git

export GO111MODULE = on

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq ($(WITH_CLEVELDB),yes)
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
empty = $(whitespace) $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(empty),$(comma),$(build_tags))

BUILD_FLAGS := -tags "$(build_tags_comma_sep)" -trimpath

build: go.sum
ifeq ($(OS),Windows_NT)
	exit 1
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/portal ./cmd/portal
endif

build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/

gen:
	solc --abi contracts/DRC.sol > contracts/DRC_sol_DRC.abi
	solc --bin contracts/DRC.sol > contracts/DRC_sol_DRC.bin
compile:
	abigen --bin contracts/DRC_sol_DRC.bin --abi contracts/DRC_sol_DRC.abi --pkg BEP20Token --alias _symbol=privateSymbol,_name=privateName,_decimals=privateDecimals --out internal/types/token.go

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/portal

########################################
### Tools & dependencies

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

draw-deps:
	@# requires brew install graphviz or apt-get install graphviz
	go get github.com/RobotsAndPencils/goviz
	@goviz -i ./cmd/portal -d 2 | dot -Tpng -o dependency-graph.png

clean:
	rm -rf snapcraft-local.yaml build/

distclean: clean
	rm -rf vendor/

########################################
### Testing


test: test-unit
test-all: check test-race test-cover

test-unit:
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' ./...

test-race:
	@VERSION=$(VERSION) go test -mod=readonly -race -tags='ledger test_ledger_mock' ./...

test-cover:
	@go test -mod=readonly -timeout 30m -race -coverprofile=coverage.txt -covermode=atomic -tags='ledger test_ledger_mock' ./...


benchmark:
	@go test -mod=readonly -bench=. ./...
