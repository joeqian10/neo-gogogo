package ecc

import "math/big"

type ECCurve struct {
	Q        *big.Int
	A        *ECFieldElement
	B        *ECFieldElement
	N        *big.Int
	Infinity *ECPoint
	G        *ECPoint
}


