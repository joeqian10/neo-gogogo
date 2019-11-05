package keys

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
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

func GenerateKeyPair() (key *KeyPair, err error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	key = &KeyPair{priv.D.Bytes(), PublicKey{priv.X, priv.Y}}
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

// sign message with KeyPair
func (p *KeyPair) Sign(message []byte) ([]byte, error) {
	privateKey := p.ToEcdsa()
	hash := sha256.Sum256(message)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])

	if err != nil {
		return nil, err
	}

	params := privateKey.Curve.Params()
	curveOrderByteSize := params.P.BitLen() / 8
	rBytes, sBytes := r.Bytes(), s.Bytes()
	signature := make([]byte, curveOrderByteSize*2)
	copy(signature[curveOrderByteSize-len(rBytes):], rBytes)
	copy(signature[curveOrderByteSize*2-len(sBytes):], sBytes)

	return signature, nil
}

// Verify returns true if the signature is valid and corresponds
// to the hash and public key
func VerifySignature(message []byte, signature []byte, p PublicKey) bool {
	hash := sha256.Sum256(message)
	publicKey := &ecdsa.PublicKey{}
	publicKey.Curve = elliptic.P256()
	publicKey.X = p.X
	publicKey.Y = p.Y
	if p.X == nil || p.Y == nil {
		return false
	}
	rBytes := new(big.Int).SetBytes(signature[0:32])
	sBytes := new(big.Int).SetBytes(signature[32:64])
	return ecdsa.Verify(publicKey, hash[:], rBytes, sBytes)
}
