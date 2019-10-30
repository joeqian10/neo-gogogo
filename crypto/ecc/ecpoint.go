package ecc

import (
	"fmt"
	hlp "github.com/joeqian10/neo-gogogo/helper"
	"math/big"
)

type ECPoint struct {
	X     *ECFieldElement
	Y     *ECFieldElement
	Curve *ECCurve
}

func (p *ECPoint) IsInfinity() bool {
	return p.X == nil && p.Y == nil
}

func NewECPoint(x *ECFieldElement, y *ECFieldElement, c *ECCurve) (*ECPoint, error) {
	if x == nil || y == nil {
		return nil, fmt.Errorf("One of the field element is nil.")
	}
	return &ECPoint{
		X:     x,
		Y:     y,
		Curve: c,
	}, nil
}

func (p *ECPoint) CompareTo(q *ECPoint) int {
	if p == q {
		return 0
	}
	result := p.X.CompareTo(q.X)
	if result != 0 {
		return result
	}
	return p.Y.CompareTo(q.Y)
}

func DecodePoint(encoded []byte, curve *ECCurve) (*ECPoint, error) {
	var p *ECPoint
	expectedLength := (hlp.GetBitLength(curve.Q) + 7) / 8
	switch encoded[0] {
	case 0x00: // infinity
		if len(encoded) != 1 {
			return nil, fmt.Errorf("Incorrect length for infinity encoding.")
		}
		p = curve.Infinity
		break
	case 0x02: // compressed
	case 0x03: // compressed
		if len(encoded) != (expectedLength + 1) {
			return nil, fmt.Errorf("Incorrect length for compressed encoding.")
		}
		yTilde := encoded[0] & 1
		b := append(hlp.ReverseBytes(encoded[1:]), byte(0))
		X1 := hlp.BigIntFromBytes(b)
		p, _ = DecompressPoint(int(yTilde), X1, curve)
	case 0x04: // uncompressed
	case 0x06: // hybrid
	case 0x07: // hybrid
		if len(encoded) != (2*expectedLength + 1) {
			return nil, fmt.Errorf("Incorrect length for uncompressed/hybrid encoding.")
		}
		X1 := hlp.BigIntFromBytes(append(hlp.ReverseBytes(encoded[1:1+expectedLength]), byte(0)))
		Y1 := hlp.BigIntFromBytes(append(hlp.ReverseBytes(encoded[1+expectedLength:]), byte(0)))
		f1, _ := NewECFieldElement(X1, curve)
		f2, _ := NewECFieldElement(Y1, curve)
		p, _ = NewECPoint(f1, f2, curve)
	default:
		return nil, fmt.Errorf("Invalid point encoding.")
	}
	return p, nil
}

func DecompressPoint(yTilde int, X1 *big.Int, curve *ECCurve) (*ECPoint, error) {
	x, err := NewECFieldElement(X1, curve)
	if err != nil {
		return nil, err
	}
	alpha := Add(Mtpl(x, Add(Square(x), curve.A)), curve.B)
	beta := Sqrt(alpha)
	//
	// if we can't find a sqrt we haven't got a point on the
	// curve - run!
	//
	if beta == nil {
		return nil, fmt.Errorf("Invalid point compression")
	}
	betaValue := beta.Value
	var bit0 int
	if hlp.IsEven(betaValue) {
		bit0 = 0
	} else {
		bit0 = 1
	}
	if bit0 != yTilde {
		// use the other root
		beta, err = NewECFieldElement(hlp.Minus(curve.Q, betaValue), curve)
		if err != nil {
			return nil, err
		}
	}
	return NewECPoint(x, beta, curve)
}
