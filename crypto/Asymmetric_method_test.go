package crypto

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/taiyuechain/taiyuechain/crypto/gm/sm2"
	"github.com/taiyuechain/taiyuechain/crypto/gm/sm3"
	"github.com/taiyuechain/taiyuechain/crypto/gm/sm4"
	"reflect"
	"testing"
)

//func TestECDSASeries_Decrypt(t *testing.T) {
//	type fields struct {
//		privateKey ecdsa.PrivateKey
//		publicKey  ecdsa.PublicKey
//	}
//	type args struct {
//		pin []byte
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    []byte
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			pri := &ECDSASeries{
//				privateKey: tt.fields.privateKey,
//				publicKey:  tt.fields.publicKey,
//			}
//			got, err := pri.Decrypt(tt.args.pin)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Decrypt() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestECDSASeries_Encrypt(t *testing.T) {
//	type fields struct {
//		privateKey ecdsa.PrivateKey
//		publicKey  ecdsa.PublicKey
//	}
//	type args struct {
//		in []byte
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    []byte
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			pub := &ECDSASeries{
//				privateKey: tt.fields.privateKey,
//				publicKey:  tt.fields.publicKey,
//			}
//			got, err := pub.Encrypt(tt.args.in)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Encrypt() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestECDSASeries_FromECDSA(t *testing.T) {
	type fields struct {
		privateKey ecdsa.PrivateKey
		publicKey  ecdsa.PublicKey
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pri := &ECDSASeries{
				privateKey: tt.fields.privateKey,
				publicKey:  tt.fields.publicKey,
			}
			if got := pri.FromECDSA(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromECDSA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestECDSASeries_FromECDSAPub(t *testing.T) {
	type fields struct {
		privateKey ecdsa.PrivateKey
		publicKey  ecdsa.PublicKey
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pub := &ECDSASeries{
				privateKey: tt.fields.privateKey,
				publicKey:  tt.fields.publicKey,
			}
			if got := pub.FromECDSAPub(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromECDSAPub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestECDSASeries_GenerateKey(t *testing.T) {
	type fields struct {
		privateKey ecdsa.PrivateKey
		publicKey  ecdsa.PublicKey
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pri := &ECDSASeries{
				privateKey: tt.fields.privateKey,
				publicKey:  tt.fields.publicKey,
			}
			if err := pri.GenerateKey(); (err != nil) != tt.wantErr {
				t.Errorf("GenerateKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestECDSASeries_HexToECDSA(t *testing.T) {
	type fields struct {
		privateKey ecdsa.PrivateKey
		publicKey  ecdsa.PublicKey
	}
	type args struct {
		hexkey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pri := &ECDSASeries{
				privateKey: tt.fields.privateKey,
				publicKey:  tt.fields.publicKey,
			}
			if err := pri.HexToECDSA(tt.args.hexkey); (err != nil) != tt.wantErr {
				t.Errorf("HexToECDSA() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestECDSASeries_Keccak256(t *testing.T) {
	type fields struct {
		privateKey ecdsa.PrivateKey
		publicKey  ecdsa.PublicKey
	}
	type args struct {
		data [][]byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pub := &ECDSASeries{
				privateKey: tt.fields.privateKey,
				publicKey:  tt.fields.publicKey,
			}
			if got := pub.Keccak256(tt.args.data...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keccak256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestECDSASeries_PubkeyToAddress(t *testing.T) {
	type fields struct {
		privateKey ecdsa.PrivateKey
		publicKey  ecdsa.PublicKey
	}
	tests := []struct {
		name   string
		fields fields
		want   common.Address
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pub := &ECDSASeries{
				privateKey: tt.fields.privateKey,
				publicKey:  tt.fields.publicKey,
			}
			if got := pub.PubkeyToAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PubkeyToAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestECDSASeries_Sign(t *testing.T) {
	type fields struct {
		privateKey ecdsa.PrivateKey
		publicKey  ecdsa.PublicKey
	}
	type args struct {
		in []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pri := &ECDSASeries{
				privateKey: tt.fields.privateKey,
				publicKey:  tt.fields.publicKey,
			}
			got, err := pri.Sign(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sign() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestECDSASeries_UnmarshalPubkey(t *testing.T) {
	type fields struct {
		privateKey ecdsa.PrivateKey
		publicKey  ecdsa.PublicKey
	}
	type args struct {
		pubbytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pub := &ECDSASeries{
				privateKey: tt.fields.privateKey,
				publicKey:  tt.fields.publicKey,
			}
			if err := pub.UnmarshalPubkey(tt.args.pubbytes); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalPubkey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestECDSASeries_Verify(t *testing.T) {
	type fields struct {
		privateKey ecdsa.PrivateKey
		publicKey  ecdsa.PublicKey
	}
	type args struct {
		userId []byte
		src    []byte
		sign   []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pub := &ECDSASeries{
				privateKey: tt.fields.privateKey,
				publicKey:  tt.fields.publicKey,
			}
			if got := pub.Verify(tt.args.userId, tt.args.src, tt.args.sign); got != tt.want {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestECDSASeries_toECDSA(t *testing.T) {
	type fields struct {
		privateKey ecdsa.PrivateKey
		publicKey  ecdsa.PublicKey
	}
	type args struct {
		d []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pri := &ECDSASeries{
				privateKey: tt.fields.privateKey,
				publicKey:  tt.fields.publicKey,
			}
			if err := pri.toECDSA(tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("toECDSA() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSmSeries_Decrypt(t *testing.T) {
	type fields struct {
		smPrivateKey sm2.SmPrivateKey
		smPublicKey  sm2.SmPublicKey
		sm3Digest    sm3.Sm3Digest
		sm4Cipher    sm4.Sm4Cipher
	}
	type args struct {
		in []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			smpri := &SmSeries{
				smPrivateKey: tt.fields.smPrivateKey,
				smPublicKey:  tt.fields.smPublicKey,
				sm3Digest:    tt.fields.sm3Digest,
				sm4Cipher:    tt.fields.sm4Cipher,
			}
			got, err := smpri.Decrypt(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmSeries_Encrypt(t *testing.T) {
	type fields struct {
		smPrivateKey sm2.SmPrivateKey
		smPublicKey  sm2.SmPublicKey
		sm3Digest    sm3.Sm3Digest
		sm4Cipher    sm4.Sm4Cipher
	}
	type args struct {
		in []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			smpub := &SmSeries{
				smPrivateKey: tt.fields.smPrivateKey,
				smPublicKey:  tt.fields.smPublicKey,
				sm3Digest:    tt.fields.sm3Digest,
				sm4Cipher:    tt.fields.sm4Cipher,
			}
			got, err := smpub.Encrypt(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmSeries_FromECDSA(t *testing.T) {
	type fields struct {
		smPrivateKey sm2.SmPrivateKey
		smPublicKey  sm2.SmPublicKey
		sm3Digest    sm3.Sm3Digest
		sm4Cipher    sm4.Sm4Cipher
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			smpri := &SmSeries{
				smPrivateKey: tt.fields.smPrivateKey,
				smPublicKey:  tt.fields.smPublicKey,
				sm3Digest:    tt.fields.sm3Digest,
				sm4Cipher:    tt.fields.sm4Cipher,
			}
			if got := smpri.FromECDSA(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromECDSA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmSeries_FromECDSAPub(t *testing.T) {
	type fields struct {
		smPrivateKey sm2.SmPrivateKey
		smPublicKey  sm2.SmPublicKey
		sm3Digest    sm3.Sm3Digest
		sm4Cipher    sm4.Sm4Cipher
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			smpub := &SmSeries{
				smPrivateKey: tt.fields.smPrivateKey,
				smPublicKey:  tt.fields.smPublicKey,
				sm3Digest:    tt.fields.sm3Digest,
				sm4Cipher:    tt.fields.sm4Cipher,
			}
			if got := smpub.FromECDSAPub(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromECDSAPub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmSeries_GenerateKey(t *testing.T) {
	type Fields struct {
		smPrivateKey sm2.SmPrivateKey
		smPublicKey  sm2.SmPublicKey
		sm3Digest    sm3.Sm3Digest
		sm4Cipher    sm4.Sm4Cipher
	}
	tests := []struct {
		name    string
		fields  Fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "caoliang",
			fields:  Fields{smPrivateKey: sm2.SmPrivateKey{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			smpri := &SmSeries{
				smPrivateKey: tt.fields.smPrivateKey,
				smPublicKey:  tt.fields.smPublicKey,
				sm3Digest:    tt.fields.sm3Digest,
				sm4Cipher:    tt.fields.sm4Cipher,
			}
			if err := smpri.GenerateKey(); (err != nil) != tt.wantErr {
				t.Errorf("GenerateKey() error = %v, wantErr %v", err, tt.wantErr)

			}
			fmt.Println("caoliang")
			fmt.Println(smpri.smPrivateKey)
			fmt.Println(smpri.smPrivateKey.SmPublicKey)
			fmt.Println(smpri.smPrivateKey.GetRawBytes())

		})
	}

}

func TestSmSeries_HexToECDSA(t *testing.T) {
	type fields struct {
		smPrivateKey sm2.SmPrivateKey
		smPublicKey  sm2.SmPublicKey
		sm3Digest    sm3.Sm3Digest
		sm4Cipher    sm4.Sm4Cipher
	}
	type args struct {
		hexkey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			smpri := &SmSeries{
				smPrivateKey: tt.fields.smPrivateKey,
				smPublicKey:  tt.fields.smPublicKey,
				sm3Digest:    tt.fields.sm3Digest,
				sm4Cipher:    tt.fields.sm4Cipher,
			}
			if err := smpri.HexToECDSA(tt.args.hexkey); (err != nil) != tt.wantErr {
				t.Errorf("HexToECDSA() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSmSeries_Keccak256(t *testing.T) {
	type fields struct {
		smPrivateKey sm2.SmPrivateKey
		smPublicKey  sm2.SmPublicKey
		sm3Digest    sm3.Sm3Digest
		sm4Cipher    sm4.Sm4Cipher
	}
	type args struct {
		data [][]byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			smpub := &SmSeries{
				smPrivateKey: tt.fields.smPrivateKey,
				smPublicKey:  tt.fields.smPublicKey,
				sm3Digest:    tt.fields.sm3Digest,
				sm4Cipher:    tt.fields.sm4Cipher,
			}
			if got := smpub.Keccak256(tt.args.data...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keccak256() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmSeries_PubkeyToAddress(t *testing.T) {
	type fields struct {
		smPrivateKey sm2.SmPrivateKey
		smPublicKey  sm2.SmPublicKey
		sm3Digest    sm3.Sm3Digest
		sm4Cipher    sm4.Sm4Cipher
	}
	tests := []struct {
		name   string
		fields fields
		want   common.Address
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			smpub := &SmSeries{
				smPrivateKey: tt.fields.smPrivateKey,
				smPublicKey:  tt.fields.smPublicKey,
				sm3Digest:    tt.fields.sm3Digest,
				sm4Cipher:    tt.fields.sm4Cipher,
			}
			if got := smpub.PubkeyToAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PubkeyToAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmSeries_Sign(t *testing.T) {
	type fields struct {
		smPrivateKey sm2.SmPrivateKey
		smPublicKey  sm2.SmPublicKey
		sm3Digest    sm3.Sm3Digest
		sm4Cipher    sm4.Sm4Cipher
	}
	type args struct {
		in []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			smpri := &SmSeries{
				smPrivateKey: tt.fields.smPrivateKey,
				smPublicKey:  tt.fields.smPublicKey,
				sm3Digest:    tt.fields.sm3Digest,
				sm4Cipher:    tt.fields.sm4Cipher,
			}
			got, err := smpri.Sign(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sign() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmSeries_UnmarshalPubkey(t *testing.T) {
	type fields struct {
		smPrivateKey sm2.SmPrivateKey
		smPublicKey  sm2.SmPublicKey
		sm3Digest    sm3.Sm3Digest
		sm4Cipher    sm4.Sm4Cipher
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			smpub := &SmSeries{
				smPrivateKey: tt.fields.smPrivateKey,
				smPublicKey:  tt.fields.smPublicKey,
				sm3Digest:    tt.fields.sm3Digest,
				sm4Cipher:    tt.fields.sm4Cipher,
			}
			if err := smpub.UnmarshalPubkey(tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalPubkey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSmSeries_Verify(t *testing.T) {
	type fields struct {
		smPrivateKey sm2.SmPrivateKey
		smPublicKey  sm2.SmPublicKey
		sm3Digest    sm3.Sm3Digest
		sm4Cipher    sm4.Sm4Cipher
	}
	type args struct {
		userId []byte
		src    []byte
		sign   []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			smpub := &SmSeries{
				smPrivateKey: tt.fields.smPrivateKey,
				smPublicKey:  tt.fields.smPublicKey,
				sm3Digest:    tt.fields.sm3Digest,
				sm4Cipher:    tt.fields.sm4Cipher,
			}
			if got := smpub.Verify(tt.args.userId, tt.args.src, tt.args.sign); got != tt.want {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmSeries_toECDSA(t *testing.T) {
	type fields struct {
		smPrivateKey sm2.SmPrivateKey
		smPublicKey  sm2.SmPublicKey
		sm3Digest    sm3.Sm3Digest
		sm4Cipher    sm4.Sm4Cipher
	}
	type args struct {
		d []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			smpri := &SmSeries{
				smPrivateKey: tt.fields.smPrivateKey,
				smPublicKey:  tt.fields.smPublicKey,
				sm3Digest:    tt.fields.sm3Digest,
				sm4Cipher:    tt.fields.sm4Cipher,
			}
			if err := smpri.toECDSA(tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("toECDSA() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
