package ecc

import (
	"github.com/joeqian10/neo-gogogo/helper"
	"math/big"
)

type ECCurve struct {
	Q        *big.Int
	A        *ECFieldElement
	B        *ECFieldElement
	N        *big.Int
	Infinity *ECPoint
	G        *ECPoint
}

func NewECCurve(Q, A, B, N *big.Int, G []byte) *ECCurve {
	curve := &ECCurve{}
	curve.Q = Q
	curve.A, _ = NewECFieldElement(A, curve)
	curve.B, _ = NewECFieldElement(B, curve)
	curve.N = N
	curve.Infinity, _ = NewECPoint(nil, nil, curve)
	curve.G, _ = DecodePoint(G, curve)
	return curve
}

var Secp256k1, Secp256r1 *ECCurve

func init() {
	/* See Certicom's SEC2 2.7.1, pg.15 */
	/* secp256k1 elliptic curve parameters */
	q, _ := new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
	a := big.NewInt(0)
	b := big.NewInt(7)
	n, _ := new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
	g := helper.HexTobytes("04" + "79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798" + "483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8")
	Secp256k1 = NewECCurve(q, a, b, n, g)

	//NEO uses SECP256R1
	q1, _ := new(big.Int).SetString("FFFFFFFF00000001000000000000000000000000FFFFFFFFFFFFFFFFFFFFFFFF", 16)
	a1, _ := new(big.Int).SetString("FFFFFFFF00000001000000000000000000000000FFFFFFFFFFFFFFFFFFFFFFFC", 16)
	b1, _ := new(big.Int).SetString("5AC635D8AA3A93E7B3EBBD55769886BC651D06B0CC53B0F63BCE3C3E27D2604B", 16)
	n1, _ := new(big.Int).SetString("FFFFFFFF00000000FFFFFFFFFFFFFFFFBCE6FAADA7179E84F3B9CAC2FC632551", 16)
	g1 := helper.HexTobytes("04" + "6B17D1F2E12C4247F8BCE6E563A440F277037D812DEB33A0F4A13945D898C296" + "4FE342E2FE1A7F9B8EE7EB4A7C0F9E162BCE33576B315ECECBB6406837BF51F5")
	Secp256r1 = NewECCurve(q1, a1, b1, n1, g1)
}
