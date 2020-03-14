package crypto

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/taiyuechain/taiyuechain/crypto/ecies"
	"github.com/taiyuechain/taiyuechain/crypto/gm/sm2"
	"github.com/taiyuechain/taiyuechain/crypto/gm/sm3"
	"github.com/taiyuechain/taiyuechain/crypto/gm/sm4"
	"github.com/taiyuechain/taiyuechain/crypto/gm/util"
	"github.com/taiyuechain/taiyuechain/crypto/secp256k1"

	"golang.org/x/crypto/sha3"
	"math/big"
)

type SmSeries struct {
	smPrivateKey sm2.SmPrivateKey
	smPublicKey  sm2.SmPublicKey
	sm3Digest    sm3.Sm3Digest
	sm4Cipher    sm4.Sm4Cipher
}
type ECDSASeries struct {
	privateKey ecdsa.PrivateKey
	publicKey  ecdsa.PublicKey
}

//type smPrivateKey sm2.SmPrivateKey
//type smPublicKey sm2.SmPublicKey

//sm generateKey method
//func (smpri *smPrivateKey)Hash256(data []byte) []byte{
//	d :=sm3.New()
//	d.Write(data)
//	digest := d.Sum(nil)
//	return digest
//}
func (smpub *SmSeries) Keccak256(data ...[]byte) []byte {
	d := sm3.New()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}
func (smpri *SmSeries) GenerateKey() error {

	sm2P256V1 := sm2.GetSm2P256V1()
	priv, x, y, err := elliptic.GenerateKey(sm2P256V1, rand.Reader)
	if err != nil {
		return err
	}

	privateKey := new(sm2.SmPrivateKey)
	privateKey.Curve = sm2P256V1
	privateKey.D = new(big.Int).SetBytes(priv)
	publicKey := new(sm2.SmPublicKey)
	publicKey.Curve = sm2P256V1
	publicKey.X = x
	publicKey.Y = y
	smpri.smPrivateKey = *privateKey
	smpri.smPublicKey = *publicKey
	return nil
}

//bytes to privatekey
func (smpri *SmSeries) toECDSA(d []byte) error {
	if len(d) != sm2.KeyBytes {
		return errors.New("Private key raw bytes length must be " + string(sm2.KeyBytes))
	}
	privateKey := new(sm2.SmPrivateKey)
	privateKey.Curve = smpri.smPrivateKey.Curve
	privateKey.D = new(big.Int).SetBytes(d)
	smpri.smPrivateKey = *privateKey
	return nil
}
func (smpri *SmSeries) FromECDSA() []byte {
	dBytes := smpri.smPrivateKey.D.Bytes()
	dl := len(dBytes)
	if dl > sm2.KeyBytes {
		raw := make([]byte, sm2.KeyBytes)
		copy(raw, dBytes[dl-sm2.KeyBytes:])
		return raw
	} else if dl < sm2.KeyBytes {
		raw := make([]byte, sm2.KeyBytes)
		copy(raw[sm2.KeyBytes-dl:], dBytes)
		return raw
	} else {
		return dBytes
	}
}
func (smpri *SmSeries) HexToECDSA(hexkey string) error {
	b, err := hex.DecodeString(hexkey)
	if err != nil {
		return errors.New("invalid hex string")
	}
	smpri.toECDSA(b)
	return nil
}
func (smpri *SmSeries) Sign(in []byte) ([]byte, error) {
	r, s, err := sm2.SignToRS(&smpri.smPrivateKey, nil, in)
	if err != nil {
		return nil, err
	}
	return sm2.MarshalSign(r, s)
}
func (smpub *SmSeries) UnmarshalPubkey(bytes []byte) error {
	if len(bytes) != sm2.KeyBytes*2 {
		return errors.New("Public key raw bytes length must be " + string(sm2.KeyBytes*2))
	}
	publicKey := new(sm2.SmPublicKey)
	publicKey.Curve = sm2.GetSm2P256V1()
	publicKey.X = new(big.Int).SetBytes(bytes[:sm2.KeyBytes])
	publicKey.Y = new(big.Int).SetBytes(bytes[sm2.KeyBytes:])
	return nil
}
func (smpub *SmSeries) FromECDSAPub() []byte {
	xBytes := smpub.smPublicKey.X.Bytes()
	yBytes := smpub.smPublicKey.Y.Bytes()
	xl := len(xBytes)
	yl := len(yBytes)

	raw := make([]byte, 1+sm2.KeyBytes*2)
	raw[0] = sm2.UnCompress
	if xl > sm2.KeyBytes {
		copy(raw[1:1+sm2.KeyBytes], xBytes[xl-sm2.KeyBytes:])
	} else if xl < sm2.KeyBytes {
		copy(raw[1+(sm2.KeyBytes-xl):1+sm2.KeyBytes], xBytes)
	} else {
		copy(raw[1:1+sm2.KeyBytes], xBytes)
	}

	if yl > sm2.KeyBytes {
		copy(raw[1+sm2.KeyBytes:], yBytes[yl-sm2.KeyBytes:])
	} else if yl < sm2.KeyBytes {
		copy(raw[1+sm2.KeyBytes+(sm2.KeyBytes-yl):], yBytes)
	} else {
		copy(raw[1+sm2.KeyBytes:], yBytes)
	}
	return raw[1:]
}
func (smpub *SmSeries) PubkeyToAddress() common.Address {
	pubBytes := smpub.FromECDSAPub()
	return common.BytesToAddress(smpub.Keccak256(pubBytes[1:])[12:])
}
func (smpub *SmSeries) Verify(userId []byte, src []byte, sign []byte) bool {
	r, s, err := sm2.UnmarshalSign(sign)
	if err != nil {
		return false
	}

	return sm2.VerifyByRS(&smpub.smPublicKey, userId, src, r, s)
}
func (smpub *SmSeries) Encrypt(in []byte) ([]byte, error) {
	cipherTextType := sm2.C1C2C3
	c2 := make([]byte, len(in))
	copy(c2, in)
	var c1 []byte
	digest := sm3.New()
	var kPBx, kPBy *big.Int
	for {
		k, err := sm2.NextK(rand.Reader, smpub.smPublicKey.Curve.N)
		if err != nil {
			return nil, err
		}
		kBytes := k.Bytes()
		c1x, c1y := smpub.smPublicKey.Curve.ScalarBaseMult(kBytes)
		c1 = elliptic.Marshal(smpub.smPublicKey.Curve, c1x, c1y)
		kPBx, kPBy = smpub.smPublicKey.Curve.ScalarMult(smpub.smPublicKey.X, smpub.smPublicKey.Y, kBytes)
		sm2.Kdf(digest, kPBx, kPBy, c2)

		if !sm2.NotEncrypted(c2, in) {
			break
		}
	}

	digest.Reset()
	digest.Write(kPBx.Bytes())
	digest.Write(in)
	digest.Write(kPBy.Bytes())
	c3 := digest.Sum(nil)

	c1Len := len(c1)
	c2Len := len(c2)
	c3Len := len(c3)
	result := make([]byte, c1Len+c2Len+c3Len)
	if cipherTextType == sm2.C1C2C3 {
		copy(result[:c1Len], c1)
		copy(result[c1Len:c1Len+c2Len], c2)
		copy(result[c1Len+c2Len:], c3)
	} else if cipherTextType == sm2.C1C3C2 {
		copy(result[:c1Len], c1)
		copy(result[c1Len:c1Len+c3Len], c3)
		copy(result[c1Len+c3Len:], c2)
	} else {
		return nil, errors.New("unknown cipherTextType:" + string(cipherTextType))
	}
	return result, nil
}
func (smpri *SmSeries) Decrypt(in []byte) ([]byte, error) {
	cipherTextType := sm2.C1C2C3
	c1Len := ((smpri.smPrivateKey.Curve.BitSize+7)/8)*2 + 1
	c1 := make([]byte, c1Len)
	copy(c1, in[:c1Len])
	c1x, c1y := elliptic.Unmarshal(smpri.smPrivateKey.Curve, c1)
	sx, sy := smpri.smPrivateKey.Curve.ScalarMult(c1x, c1y, sm2.Sm2H.Bytes())
	if util.IsEcPointInfinity(sx, sy) {
		return nil, errors.New("[h]C1 at infinity")
	}
	c1x, c1y = smpri.smPrivateKey.Curve.ScalarMult(c1x, c1y, smpri.smPrivateKey.D.Bytes())

	digest := sm3.New()
	c3Len := digest.Size()
	c2Len := len(in) - c1Len - c3Len
	c2 := make([]byte, c2Len)
	c3 := make([]byte, c3Len)
	if cipherTextType == sm2.C1C2C3 {
		copy(c2, in[c1Len:c1Len+c2Len])
		copy(c3, in[c1Len+c2Len:])
	} else if cipherTextType == sm2.C1C3C2 {
		copy(c3, in[c1Len:c1Len+c3Len])
		copy(c2, in[c1Len+c3Len:])
	} else {
		return nil, errors.New("unknown cipherTextType:" + string(cipherTextType))
	}

	sm2.Kdf(digest, c1x, c1y, c2)

	digest.Reset()
	digest.Write(c1x.Bytes())
	digest.Write(c2)
	digest.Write(c1y.Bytes())
	newC3 := digest.Sum(nil)

	if !bytes.Equal(newC3, c3) {
		return nil, errors.New("invalid cipher text")
	}
	return c2, nil
}

//type privateKey ecdsa.PrivateKey
//type publicKey ecdsa.PublicKey

//generateKey method
func (pri *ECDSASeries) GenerateKey() error {
	p, err := ecdsa.GenerateKey(S256(), rand.Reader)
	if err != nil {
		return err
	}
	pri.privateKey = *p
	pri.publicKey = p.PublicKey
	return nil
}
func (pri *ECDSASeries) toECDSA(d []byte) error {
	strict := false
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = S256()
	if strict && 8*len(d) != priv.Params().BitSize {
		return fmt.Errorf("invalid length, need %d bits", priv.Params().BitSize)
	}
	priv.D = new(big.Int).SetBytes(d)

	// The priv.D must < N
	if priv.D.Cmp(secp256k1N) >= 0 {
		return fmt.Errorf("invalid private key, >=N")
	}
	// The priv.D must not be zero or negative.
	if priv.D.Sign() <= 0 {
		return fmt.Errorf("invalid private key, zero or negative")
	}

	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(d)
	if priv.PublicKey.X == nil {
		return errors.New("invalid private key")
	}
	pri.privateKey = *(priv)
	return nil
}
func (pri *ECDSASeries) FromECDSA() []byte {
	if pri == nil {
		return nil
	}
	return math.PaddedBigBytes(pri.privateKey.D, pri.privateKey.Params().BitSize/8)
}
func (pri *ECDSASeries) HexToECDSA(hexkey string) error {
	b, err := hex.DecodeString(hexkey)
	if err != nil {
		return errors.New("invalid hex string")
	}
	pri.toECDSA(b)
	return nil
}
func (pub *ECDSASeries) UnmarshalPubkey(pubbytes []byte) error {
	x, y := elliptic.Unmarshal(S256(), pubbytes)
	if x == nil {
		return errInvalidPubkey
	}
	pub.publicKey = *(&ecdsa.PublicKey{Curve: S256(), X: x, Y: y})
	return nil
}
func (pub *ECDSASeries) FromECDSAPub() []byte {
	if pub == nil || pub.publicKey.X == nil || pub.publicKey.Y == nil {
		return nil
	}
	return elliptic.Marshal(S256(), pub.publicKey.X, pub.publicKey.Y)
}
func (pub *ECDSASeries) PubkeyToAddress() common.Address {
	pubBytes := pub.FromECDSAPub()
	return common.BytesToAddress(pub.Keccak256(pubBytes[1:])[12:])
}
func (pub *ECDSASeries) Keccak256(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}
func (pri *ECDSASeries) Sign(in []byte) ([]byte, error) {
	if len(in) != DigestLength {
		return nil, fmt.Errorf("hash is required to be exactly %d bytes (%d)", DigestLength, len(in))
	}
	seckey := math.PaddedBigBytes(pri.privateKey.D, pri.privateKey.Params().BitSize/8)
	defer zeroBytes(seckey)
	return secp256k1.Sign(in, seckey)
}
func (pub *ECDSASeries) Verify(userId []byte, src []byte, sign []byte) bool {
	return secp256k1.VerifySignature(pub.FromECDSAPub(), src, sign)
}
func (pub *ECDSASeries) Encrypt(in []byte) ([]byte, error) {
	key := ecies.PublicKey{pub.publicKey.X, pub.publicKey.Y, pub.publicKey.Curve, ecies.ECIES_AES256_SHA256}
	return ecies.Encrypt(rand.Reader, &key, in, nil, nil)
}
func (pri *ECDSASeries) Decrypt(pin []byte) ([]byte, error) {
	key := ecies.PublicKey{pri.privateKey.PublicKey.X, pri.privateKey.PublicKey.Y, pri.privateKey.PublicKey.Curve, ecies.ECIES_AES256_SHA256}
	//prie:=ecies.PrivateKey{key,pri.D}
	return ecies.PrivateKey{key, pri.privateKey.D}.Decrypt(pin, nil, nil)
}
