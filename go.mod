module identity

go 1.11

require (
	bitbucket.org/cpchain/chain v0.0.0-00010101000000-000000000000
	github.com/cheekybits/is v0.0.0-20150225183255-68e9c0620927 // indirect
	github.com/ethereum/go-ethereum v1.8.14
	github.com/ethereum/go-ethereum/crypto/randentropy v0.0.0-00010101000000-000000000000 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/naoina/go-stringutil v0.1.0 // indirect
	github.com/onsi/ginkgo v1.14.2 // indirect
	github.com/onsi/gomega v1.10.4 // indirect
)

replace bitbucket.org/cpchain/chain => ../chain/src/bitbucket.org/cpchain/chain

replace github.com/ethereum/go-ethereum/crypto/randentropy => ../chain/src/bitbucket.org/cpchain/chain/vendor/github.com/ethereum/go-ethereum/crypto/randentropy

replace gopkg.in/fatih/set.v0 => ../chain/src/bitbucket.org/cpchain/chain/vendor/gopkg.in/fatih/set.v0

replace github.com/ethereum/go-ethereum => ../chain/src/bitbucket.org/cpchain/chain/vendor/github.com/ethereum/go-ethereum
