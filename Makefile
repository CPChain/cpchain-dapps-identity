
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

compile:
	@docker run --rm -v $(ROOT_DIR):/go/src/bitbucket.org/cpchain/chain/identity -it cpchain2018/abigen abigen --sol ./identity/identity.sol --pkg identity --out ./identity/identity.go

test:
	@docker run --rm -v $(ROOT_DIR):/go/src/bitbucket.org/cpchain/chain/identity -it cpchain2018/abigen go test ./identity
