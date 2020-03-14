package crypto

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taiyuechain/taiyuechain/crypto/gm/sm4"
	"github.com/taiyuechain/taiyuechain/crypto/gm/util"
)

var key = []byte{0x7b, 0xea, 0x0a, 0xa5, 0x45, 0x8e, 0xd1, 0xa3, 0x7d, 0xb1, 0x65, 0x2e, 0xfb, 0xc5, 0x95, 0x05}

func (sm4en *SmSeries) DesEncrypt(dst, src []byte) {
	c, err := sm4.NewCipher(key)
	if err != nil {
		log.Error(err.Error())
		return
	}

	dst = make([]byte, sm4.BlockSize+8)
	c.Decrypt(dst, util.PKCS5Padding(src, sm4.BlockSize))
	fmt.Printf("decrypt result:%s\n", hex.EncodeToString(dst))

}
func (sm4en *SmSeries) DesDecrypt(dst, src []byte) {
	c, err := sm4.NewCipher(key)
	if err != nil {
		log.Error(err.Error())
		return
	}
	dst = make([]byte, len(dst)+8)
	c.Decrypt(dst, src)
	fmt.Printf("decrypt result:%s\n", hex.EncodeToString(dst))
}
func (pri *ECDSASeries) DesEncrypt(dst, src []byte) {
	pri.aes.Put(string(key), string(src))
}
func (pri *ECDSASeries) DesDecrypt(dst, src []byte) {
	pri.aes.Get(string(key))
}
