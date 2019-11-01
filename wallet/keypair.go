package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"github.com/joeqian10/neo-gogogo/crypto"
	"math/big"
)

type KeyPair struct {
	PrivateKey []byte
	PublicKey  ecdsa.PublicKey
}

func NewKeyPair(privateKey []byte) (key *KeyPair, err error) {
	length := len(privateKey)
	if length != 32 {
		return nil, fmt.Errorf("argument length is wrong %v", length)
	}
	priv := ToEcdsa(privateKey)
	key = &KeyPair{privateKey, priv.PublicKey}
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

func (p *KeyPair) ExportWIF() string {
	data := make([]byte, 34)
	data[0] = 0x80
	copy(data[1:], p.PrivateKey)
	data[33] = 0x01
	wif := crypto.Base58CheckEncode(data)
	return wif
}
