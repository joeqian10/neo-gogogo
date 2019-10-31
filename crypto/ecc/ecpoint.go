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

func (p *ECPoint) EncodePoint(compressed bool) (data []byte) {
	if p.IsInfinity() {
		return []byte{0}
	}
	if compressed {
		data = make([]byte, 32)
	} else {
		data = make([]byte, 65)
		yBytes := p.Y.Value.Bytes()
		copy(data[65-len(yBytes):], yBytes)
	}
	xBytes := p.X.Value.Bytes()
	copy(data[33-len(xBytes):], xBytes)
	data[0] = 0x04
	if compressed {
		if hlp.IsEven(p.Y.Value) {
			data[0] = 0x02
		} else {
			data[0] = 0x03
		}
	}
	return data
}

func (p *ECPoint) Negate() (result *ECPoint) {
	return &ECPoint{p.X, p.Y.Negate(), p.Curve}
}

func (p *ECPoint) Add(a, b *ECPoint) (result *ECPoint) {
	if a.IsInfinity() {
		return b
	}
	if b.IsInfinity() {
		return a
	}
	if a.X.Equals(b.X) {
		if a.Y.Equals(b.Y) {
			return a.Twice()
		}
		return a.Curve.Infinity
	}
	gamma := Divide(Minus(b.Y, a.Y), Minus(b.X, a.X))
	x3 := Minus(Minus(Square(gamma), a.X), b.X)
	y3 := Minus(Mtpl(gamma, Minus(a.X, x3)), a.Y)
	result, _ = NewECPoint(x3, y3, a.Curve)
	return result
}

func (p *ECPoint) Minus(a, b *ECPoint) (result *ECPoint) {
	if b.IsInfinity() {
		return a
	}
	return ECPoint{}.Add(a, b.Negate())
}

func (p *ECPoint) Multiply(k *big.Int) (result *ECPoint) {
	// floor(log2(k))
	m := k.BitLen()

	// width of the Window NAF, Required length of precomputation array
	width, reqPreCompLen := 0, 0
	if m < 13 {
		width = 2
		reqPreCompLen = 1
	} else if m < 41 {
		width = 3
		reqPreCompLen = 2
	} else if m < 121 {
		width = 4
		reqPreCompLen = 4
	} else if m < 337 {
		width = 5
		reqPreCompLen = 8
	} else if m < 897 {
		width = 6
		reqPreCompLen = 16
	} else if m < 2305 {
		width = 7
		reqPreCompLen = 32
	} else {
		width = 8
		reqPreCompLen = 127
	}

	// The length of the precomputation array
	preCompLen := 1
	preComp := []*ECPoint{p}
	twiceP := p.Twice()

	if preCompLen < reqPreCompLen {
		oldPreComp := preComp
		preComp = make([]*ECPoint, reqPreCompLen)
		copy(preComp, oldPreComp)
		for i := preCompLen; i < reqPreCompLen; i++ {
			preComp[i] = ECPoint{}.Add(twiceP, preComp[i-1])
		}
	}
	// Compute the Window NAF of the desired width
	wnaf := windowNaf(uint8(width), k)
	l := len(wnaf)
	// Apply the Window NAF to p using the precomputed ECPoint values.
	q := p.Curve.Infinity
	for i := l - 1; i >= 0; i-- {
		q = q.Twice()
		if wnaf[i] != 0 {
			if wnaf[i] > 0 {
				q = ECPoint{}.Add(q, preComp[(wnaf[i]-1)/2])
			} else {
				q = ECPoint{}.Minus(q, preComp[(wnaf[i]-1)/2])
			}
		}
	}
	return q
}

func (p *ECPoint) Twice() (result *ECPoint) {
	if p.IsInfinity() {
		return p
	}
	if p.Y.Value.Sign() == 0 {
		return p.Curve.Infinity
	}
	TWO, _ := NewECFieldElement(big.NewInt(2), p.Curve)
	THREE, _ := NewECFieldElement(big.NewInt(2), p.Curve)
	gamma := Divide(Add(Mtpl(Square(p.X), THREE), p.Curve.A), Mtpl(p.Y, TWO))
	x3 := Minus(Square(gamma), Mtpl(p.X, TWO))
	y3 := Minus(Mtpl(gamma, Minus(p.X, x3)), p.Y)
	result, _ = NewECPoint(x3, y3, p.Curve)
	return result
}

func windowNaf(width uint8, k *big.Int) []int8 {
	wnaf := make([]int8, k.BitLen()+1)
	pow2wB := big.NewInt(1 << width)
	i, length := 0, 0
	for k.Sign() > 0 {
		if !hlp.IsEven(k) {
			remainder := big.Int{}.Mod(k, pow2wB)
			if hlp.TestBit(remainder, int(width-1)) {
				wnaf[i] = int8(big.Int{}.Sub(remainder, pow2wB).Int64())
			} else {
				wnaf[i] = int8(remainder.Int64())
			}
			k = hlp.Minus(k, big.NewInt(int64(wnaf[i])))
			length = i
		} else {
			wnaf[i] = 0
		}
		k = k.Rsh(k, 1)
		i++
	}
	length++
	wnafShort := make([]int8, length)
	copy(wnafShort, wnaf)
	return wnafShort
}
