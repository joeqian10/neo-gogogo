package wallet

import (
	"fmt"
	"github.com/joeqian10/neo-gogogo/crypto/ecc"
)

type KeyPair struct {
	PrivateKey []byte
	PublicKey  ecc.ECPoint
}

func NewKeyPair(privateKey []byte) (key *KeyPair, err error) {
	length := len(privateKey)
	if length != 32 && length != 94 && length != 104 {
		return nil, fmt.Errorf("argument length is wrong %v", length)
	}

	key = &KeyPair{privateKey[:32], nil}
	if length == 32 {
		//key.PublicKey =  ecc.Secp256r1.G
	} else {

	}

	return key, nil
}
