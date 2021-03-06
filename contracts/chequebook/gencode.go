// +build none

// This program generates contract/code.go, which contains the chequebook code
// after deployment.
package main

import (
	"fmt"
	"github.com/taiyuechain/taiyuechain/crypto/taiCrypto"
	"io/ioutil"
	"math/big"

	"github.com/taiyuechain/taiyuechain/accounts/abi/bind"
	"github.com/taiyuechain/taiyuechain/accounts/abi/bind/backends"
	"github.com/taiyuechain/taiyuechain/contracts/chequebook/contract"
	"github.com/taiyuechain/taiyuechain/core"
	//"github.com/taiyuechain/taiyuechain/crypto"
)

var (
	//caoliang modify
	//testKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	taiprivate taiCrypto.TaiPrivateKey
	testKey, _ = taiprivate.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	testAlloc  = core.GenesisAlloc{
		crypto.PubkeyToAddress(testKey.PublicKey): {Balance: big.NewInt(500000000000)},
	}
)

func main() {
	backend := backends.NewSimulatedBackend(testAlloc)
	auth := bind.NewKeyedTransactor(testKey)

	// Deploy the contract, get the code.
	addr, _, _, err := contract.DeployChequebook(auth, backend)
	if err != nil {
		panic(err)
	}
	backend.Commit()
	code, err := backend.CodeAt(nil, addr, nil)
	if err != nil {
		panic(err)
	}
	if len(code) == 0 {
		panic("empty code")
	}

	// Write the output file.
	content := fmt.Sprintf(`package contract

// ContractDeployedCode is used to detect suicides. This constant needs to be
// updated when the contract code is changed.
const ContractDeployedCode = "%#x"
`, code)
	if err := ioutil.WriteFile("contract/code.go", []byte(content), 0644); err != nil {
		panic(err)
	}
}
