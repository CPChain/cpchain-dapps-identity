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

	// register
	identity := "{\"pub_key\":\"hello\", \"name\": \"value\"}"
	register(t, instance, candidateKey, identity)
	contractBackend.Commit()

	checkNum(t, instance, 1)
	checkIdentity(t, instance, candidateAddr, identity)
	if _, err := instance.Get(nil, candidate2Addr); err == nil {
		t.Error("should get an error because the address have not already register its identity")
	}

	// add more
	identity2 := "{\"pub_key\":\"hello2\", \"name\": \"value\"}"
	register(t, instance, candidate2Key, identity2)
	contractBackend.Commit()

	checkNum(t, instance, 2)
	checkIdentity(t, instance, candidateAddr, identity)
	checkIdentity(t, instance, candidate2Addr, identity2)

	// remove
	remove(t, instance, candidateKey)
	contractBackend.Commit()

	checkNum(t, instance, 1)
	if _, err := instance.Get(nil, candidateAddr); err == nil {
		t.Error("should get an error because the address have not already register its identity")
	}
	checkIdentity(t, instance, candidate2Addr, identity2)

}

func register(t *testing.T, instance *identity.Identity, key *ecdsa.PrivateKey, identity string) {
	txOpts := bind.NewKeyedTransactor(key)
	txOpts.GasLimit = uint64(50000000)
	txOpts.Value = big.NewInt(0)
	_, err := instance.Register(txOpts, identity)
	if err != nil {
		t.Fatal("register failed:", err)
	}
}

func remove(t *testing.T, instance *identity.Identity, key *ecdsa.PrivateKey) {
	txOpts := bind.NewKeyedTransactor(key)
	txOpts.GasLimit = uint64(50000000)
	txOpts.Value = big.NewInt(0)
	_, err := instance.Remove(txOpts)
	if err != nil {
		t.Fatal("register failed:", err)
	}
}

func checkError(t *testing.T, title string, err error) {
	if err != nil {
		t.Fatal(title, ":", err)
	}
}

func checkIdentity(t *testing.T, instance *identity.Identity, addr common.Address, expect string) {
	got, err := instance.Get(nil, addr)
	checkError(t, "get identity", err)
	if got != expect {
		t.Errorf("got identity do not equal to the expect: '%v' != '%v'", got, expect)
	}
}

func checkNum(t *testing.T, instance *identity.Identity, amount int) {
	num, err := instance.Count(nil)
	checkError(t, "get num", err)

	if num.Cmp(new(big.Int).SetInt64(int64(amount))) != 0 {
		t.Errorf("rnode'num %d != %d", num, amount)
	}
}
