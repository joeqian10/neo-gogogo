package helper

// wrapper of big.Int

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func BigIntFromBytes(b []byte) *big.Int {
	t := &big.Int{}
	return t.SetBytes(b) // big endian
}

func BitLen(n int) int {
	var b int
	b = 0
	for n>>1 > 0 {
		b++
	}
	return b
}

// TODO unit testing
func GetBitLength(n *big.Int) int {
	b := n.Bytes() // big endian, may cause error
	x := (len(b) - 1) * 8
	if n.Sign() > 0 {
		return x + BitLen(int(b[len(b)-1]))
	} else {
		return x + BitLen(int(255-b[len(b)-1]))
	}
}

// TODO unit testing
func GetLowestSetBit(n *big.Int) int {
	if n.Sign() == 0 {
		return -1
	}
	b := n.Bytes() // big endian, may cause error
	w := 0
	for b[w] == 0 {
		w++
	}
	for i := 0; i < 8; i++ {
		if (b[w] & 1 << uint(i)) > 0 {
			return i + w*8
		}
	}
	return -2
}

// TestBit tests if a bit of a bigInt is set
func TestBit(n *big.Int, index int) bool {
	//i := index / 8
	//j := index % 8
	//b := n.Bytes()
	//return b[i] & (1 << uint(j)) > 0
	b := n.Bit(index)
	return b == 1
}

func IsEven(n *big.Int) bool {
	if Mod(n, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		return true
	}
	return false
}

func Negate(n *big.Int) *big.Int {
	t := &big.Int{}
	return t.Neg(n)
}

func Add(x *big.Int, y *big.Int) *big.Int {
	t := &big.Int{}
	return t.Add(x, y)
}

func Minus(x *big.Int, y *big.Int) *big.Int {
	t := &big.Int{}
	return t.Sub(x, y)
}

func Mtpl(x *big.Int, y *big.Int) *big.Int {
	t := &big.Int{}
	return t.Mul(x, y)
}

func Mod(x *big.Int, y *big.Int) *big.Int {
	t := &big.Int{}
	return t.Mod(x, y)
}

func ModInverse(g *big.Int, n *big.Int) *big.Int {
	t := &big.Int{}
	return t.ModInverse(g, n)
}

func ModPow(x, y, m *big.Int) *big.Int {
	t := &big.Int{}
	return t.Exp(x, y, m)
}

func LeftShift(n *big.Int, i uint) *big.Int {
	t := &big.Int{}
	return t.Lsh(n, i)
}

func RightShift(n *big.Int, i uint) *big.Int {
	t := &big.Int{}
	return t.Rsh(n, i)
}

func NextBigInt(sizeInBit int) (n *big.Int, err error) {
	if sizeInBit < 0 {
		return nil, fmt.Errorf("sizeInBits must be non-negative.")
	}
	if sizeInBit == 0 {
		return big.NewInt(0), nil
	}
	b := make([]byte, sizeInBit/8+1)
	_, err = rand.Read(b)
	if err != nil {
		return nil, err
	}
	if sizeInBit%8 == 0 {
		b[len(b)-1] = 0
	} else {
		b[len(b)-1] &= byte((1 << uint(sizeInBit) % 8) - 1)
	}
	return BigIntFromBytes(b), nil
}
