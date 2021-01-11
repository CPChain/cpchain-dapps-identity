package identity_test

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"identity"

	"bitbucket.org/cpchain/chain/accounts/abi/bind"
	"bitbucket.org/cpchain/chain/accounts/abi/bind/backends"
	"bitbucket.org/cpchain/chain/commons/log"
	"bitbucket.org/cpchain/chain/configs"
	"bitbucket.org/cpchain/chain/core"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ownerKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	ownerAddr   = crypto.PubkeyToAddress(ownerKey.PublicKey)

	candidateKey, _ = crypto.HexToECDSA("8a1f9a8f95be41cd7ccb6168179afb4504aefe388d1e14474d32c45c72ce7b7a")
	candidateAddr   = crypto.PubkeyToAddress(candidateKey.PublicKey)

	candidate2Key, _ = crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	candidate2Addr   = crypto.PubkeyToAddress(candidate2Key.PublicKey)
)

func deploy(prvKey *ecdsa.PrivateKey, backend *backends.SimulatedBackend) (common.Address, *identity.Identity) {
	deployTransactor := bind.NewKeyedTransactor(prvKey)
	add, _, instance, err := identity.DeployIdentity(deployTransactor, backend)
	if err != nil {
		log.Fatal("deploy identity failed:", "error", err)
	}
	return add, instance

}

func TestDeployAndRegister(t *testing.T) {
	// deploy contract
	contractBackend := backends.NewDporSimulatedBackend(core.GenesisAlloc{
		ownerAddr:      {Balance: big.NewInt(1000000000000)},
		candidateAddr:  {Balance: new(big.Int).Mul(big.NewInt(1000000), big.NewInt(configs.Cpc))},
		candidate2Addr: {Balance: new(big.Int).Mul(big.NewInt(1000000), big.NewInt(configs.Cpc))}})
	_, instance := deploy(ownerKey, contractBackend)
	_ = instance
}
