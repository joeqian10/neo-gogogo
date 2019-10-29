package ecc

import (
	"fmt"
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


