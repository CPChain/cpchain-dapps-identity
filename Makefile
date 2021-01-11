
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: build

compile:
	@docker run --rm -v $(ROOT_DIR):/go/src/bitbucket.org/cpchain/chain/identity -it cpchain2018/abigen abigen --sol ./identity/identity.sol --pkg identity --out ./identity/identity.go

test:
	@go test

build:
	@mkdir -p build
	@go build -o build/main cmd/main.go
