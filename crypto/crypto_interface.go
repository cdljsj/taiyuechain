package crypto

import (
	"github.com/TaiChain/common"
)

/*
caoliang
Crypto interface

*/
//hash interface

type HashAlgorithm interface {
}

//asymmetric cryptography interface
type Asymmetric interface {
	GenerateKey() error
	toECDSA(d []byte) error
	FromECDSA() []byte
	UnmarshalPubkey([]byte) error
	FromECDSAPub() []byte
	HexToECDSA(hexkey string) error
	PubkeyToAddress() common.Address
	Keccak256(data ...[]byte) []byte
	Sign(in []byte) ([]byte, error)
	Verify(userId []byte, src []byte, sign []byte) bool
	Encrypt(in []byte) ([]byte, error)
	Decrypt(pin []byte) ([]byte, error)
}

//Symmetric cryptography interface
type Symmetric interface {
	DesEncrypt(dst, src []byte)
	DesDecrypt(dst, src []byte)
}
