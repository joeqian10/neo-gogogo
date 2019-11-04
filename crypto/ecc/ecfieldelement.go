package ecc

import (
	"fmt"
	hlp "github.com/joeqian10/neo-gogogo/helper"
	"math/big"
)

type ECFieldElement struct {
	Value *big.Int
	curve *ECCurve
}

func NewECFieldElement(v *big.Int, c *ECCurve) (*ECFieldElement, error) {
	if v.Cmp(c.Q) > 0 {
		return nil, fmt.Errorf("x value is too large in field element.")
	}

	return &ECFieldElement{
		Value: v,
		curve: c,
	}, nil
}

func (x *ECFieldElement) CompareTo(y *ECFieldElement) int {
	if x == y {
		return 0
	}
	return x.Value.Cmp(y.Value)
}

func (x *ECFieldElement) Equals(y *ECFieldElement) bool {
	if x.Value.Cmp(y.Value) == 0 {
		return true
	}
	return false
}

func Square(x *ECFieldElement) *ECFieldElement {
	return &ECFieldElement{
		Value: hlp.Mod(hlp.Mtpl(x.Value, x.Value), x.curve.Q),
		curve: x.curve,
	}
}

func (x *ECFieldElement) Negate() *ECFieldElement {
	return &ECFieldElement{
		Value: hlp.Mod(hlp.Negate(x.Value), x.curve.Q),
		curve: x.curve,
	}
}

func Add(x, y *ECFieldElement) *ECFieldElement {
	return &ECFieldElement{
		Value: hlp.Mod(hlp.Add(x.Value, y.Value), x.curve.Q),
		curve: x.curve,
	}
}

func Minus(x, y *ECFieldElement) *ECFieldElement {
	return &ECFieldElement{
		Value: hlp.Mod(hlp.Minus(x.Value, y.Value), x.curve.Q),
		curve: x.curve,
	}
}

func Mtpl(x, y *ECFieldElement) *ECFieldElement {
	return &ECFieldElement{
		Value: hlp.Mod(hlp.Mtpl(x.Value, y.Value), x.curve.Q),
		curve: x.curve,
	}
}

func Divide(x, y *ECFieldElement) *ECFieldElement {
	return &ECFieldElement{
		Value: hlp.Mod(hlp.Mtpl(x.Value, hlp.ModInverse(y.Value, x.curve.Q)), x.curve.Q),
		curve: x.curve,
	}
}

func Sqrt(x *ECFieldElement) *ECFieldElement {
	if hlp.TestBit(x.curve.Q, 1) {
		z, _ := NewECFieldElement(hlp.ModPow(x.Value, hlp.Add(hlp.RightShift(x.curve.Q, 2), big.NewInt(1)), x.curve.Q), x.curve)
		if Square(z).Equals(x) {
			return z
		}
		return nil
	}
	var qMinusOne, legendreExponent, u, k, Q, fourQ, U, V *big.Int
	qMinusOne = hlp.Minus(x.curve.Q, big.NewInt(1))
	legendreExponent = hlp.RightShift(qMinusOne, 1)
	if hlp.ModPow(x.Value, legendreExponent, x.curve.Q).Cmp(big.NewInt(1)) != 0 {
		return nil
	}
	u = hlp.RightShift(qMinusOne, 2)
	k = hlp.Add(hlp.LeftShift(u, 1), big.NewInt(1))
	Q = x.Value
	fourQ = hlp.Mod(hlp.LeftShift(Q, 2), x.curve.Q)
	for ok := true; ok; ok = U.Cmp(big.NewInt(1)) == 0 || U.Cmp(qMinusOne) == 0 {
		var P *big.Int
		for ok2 := true; ok2; ok2 = P.Cmp(x.curve.Q) >= 0 ||
			hlp.ModPow(hlp.Minus(hlp.Mtpl(P, P), fourQ), legendreExponent, x.curve.Q).Cmp(qMinusOne) != 0 {
			P, _ = hlp.NextBigInt(hlp.GetBitLength(x.curve.Q))
		}
		result := fastLucasSequence(x.curve.Q, P, Q, k)
		U = result[0]
		V = result[1]
		if hlp.Mod(hlp.Mtpl(V, V), x.curve.Q).Cmp(fourQ) == 0 {
			if hlp.TestBit(V, 0) {
				V = hlp.Add(V, x.curve.Q)
			}
			V = hlp.RightShift(V, 1)
			return &ECFieldElement{Value: V, curve: x.curve} // need testing
		}
	}
	return nil
}

//
func fastLucasSequence(p *big.Int, P *big.Int, Q *big.Int, k *big.Int) []*big.Int {
	n := hlp.GetBitLength(k)
	s := hlp.GetLowestSetBit(k)

	var Uh, Vl, Vh, Ql, Qh *big.Int
	Uh = big.NewInt(1)
	Vl = big.NewInt(2)
	Vh = P
	Ql = big.NewInt(1)
	Qh = big.NewInt(1)

	for j := n - 1; j >= s+1; j-- {
		Ql = hlp.Mod(hlp.Mtpl(Ql, Qh), p)

		if hlp.TestBit(k, j) {
			Qh = hlp.Mod(hlp.Mtpl(Ql, Q), p)
			Uh = hlp.Mod(hlp.Mtpl(Uh, Vh), p)
			Vl = hlp.Mod(hlp.Minus(hlp.Mtpl(Vh, Vl), hlp.Mtpl(P, Ql)), p)
			Vh = hlp.Mod(hlp.Minus(hlp.Mtpl(Vh, Vh), hlp.LeftShift(Qh, 1)), p)
		} else {
			Qh = Ql
			Uh = hlp.Mod(hlp.Minus(hlp.Mtpl(Uh, Vl), Ql), p)
			Vh = hlp.Mod(hlp.Minus(hlp.Mtpl(Vh, Vl), hlp.Mtpl(P, Ql)), p)
			Vl = hlp.Mod(hlp.Minus(hlp.Mtpl(Vl, Vl), hlp.LeftShift(Ql, 1)), p)
		}
	}
	Ql = hlp.Mod(hlp.Mtpl(Ql, Qh), p)
	Qh = hlp.Mod(hlp.Mtpl(Ql, Q), p)
	Uh = hlp.Mod(hlp.Minus(hlp.Mtpl(Uh, Vl), Ql), p)
	Vl = hlp.Mod(hlp.Minus(hlp.Mtpl(Vh, Vl), hlp.Mtpl(P, Ql)), p)
	Ql = hlp.Mod(hlp.Mtpl(Ql, Qh), p)

	for j := 1; j <= s; j++ {
		Uh = hlp.Mtpl(hlp.Mtpl(Uh, Vl), p)
		Vl = hlp.Mod(hlp.Minus(hlp.Mtpl(Vl, Vl), hlp.LeftShift(Ql, 1)), p)
		Ql = hlp.Mod(hlp.Mtpl(Ql, Ql), p)
	}

	return []*big.Int{Uh, Vl}
}
