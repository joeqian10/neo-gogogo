package keys

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"github.com/joeqian10/neo-gogogo/crypto"
	"github.com/joeqian10/neo-gogogo/helper"
	"math/big"
)

type KeyPair struct {
	PrivateKey []byte
	PublicKey  PublicKey
}

func NewKeyPair(privateKey []byte) (key *KeyPair, err error) {
	length := len(privateKey)
	if length != 32 {
		return nil, fmt.Errorf("argument length is wrong %v", length)
	}
	priv := ToEcdsa(privateKey)
	key = &KeyPair{privateKey, PublicKey{priv.X, priv.Y}}
	return key, nil
}

// ecsda converts the key to a usable ecsda.PrivateKey for signing data.
func (p *KeyPair) ToEcdsa() *ecdsa.PrivateKey {
	return ToEcdsa(p.PrivateKey)
}

// ecsda converts the private key byte[] to a usable ecsda.PrivateKey for signing data.
func ToEcdsa(key []byte) *ecdsa.PrivateKey {
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = elliptic.P256()
	priv.D = new(big.Int).SetBytes(key)
	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(key)
	return priv
}

// export wif string
func (p *KeyPair) ExportWIF() string {
	data := make([]byte, 34)
	data[0] = 0x80
	copy(data[1:], p.PrivateKey)
	data[33] = 0x01
	wif := crypto.Base58CheckEncode(data)
	return wif
}

// export nep2 key string
func (p *KeyPair) ExportNep2(password string) (string, error) {
	nep2, err := NEP2Encrypt(p, password)
	if err != nil {
		return "", err
	}
	return nep2, nil
}

// String implements the Stringer interface.
func (p *KeyPair) String() string {
	return helper.BytesToHex(p.PrivateKey)
}
