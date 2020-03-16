package crypto

import (
	"encoding/hex"
	"github.com/taiyuechain/crypto/gm/sm3"
)

func HashAlgorithm(key string) ([]byte, error) {
	switch config.HashType {
	case SmHash:
		smHash := sm3.New()
		smHash.Write([]byte(key))
		sh := smHash.Sum(nil)
		return sh, nil

	case ECDSAhash:
		b, err := hex.DecodeString(key)
		if err != nil {
			return nil, err
		}
		return Keccak256(b), err
	}
}
