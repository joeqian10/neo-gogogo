package helper

import (
	"fmt"
	"go/constant"
	"strconv"
	"strings"
)

const PRECISION = 8
const D = 100000000

type Fixed8 struct {
	Value int64
}

var One Fixed8 = NewFixed8(int64(D))
var Satosh Fixed8 = NewFixed8(int64(1))
var Zero Fixed8 = NewFixed8(int64(0))

func NewFixed8(data int64) Fixed8 {
	return Fixed8{
		Value: data,
	}
}

// Fixed8FromInt64 returns a new Fixed8 type multiplied by decimals.
func Fixed8FromInt64(val int64) Fixed8 {
	return NewFixed8(val * D)
}

// Fixed8ToInt64 returns the original value representing Fixed8 as int64.
func Fixed8ToInt64(f Fixed8) int64 {
	return int64(f.Value) / D
}

// Fixed8FromFloat64 returns a new Fixed8 type multiplied by decimals.
func Fixed8FromFloat64(val float64) Fixed8 {
	return NewFixed8(int64(val * D))
}

// Fixed8ToFloat64 returns the decimal value of a Fixed8 type.
func Fixed8ToFloat64(f Fixed8) float64 {
	return float64(f.Value) / D
}

// Fixed8FromString parses s which must be a fixed point number with precision up to 10^-8
func Fixed8FromString(s string) (Fixed8, error) {
	parts := strings.SplitN(s, ".", 2)
	ip, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return NewFixed8(0), fmt.Errorf("Fixed8 must satisfy following regex \\d+(\\.\\d{1,8})?")
	}

	if len(parts) == 1 {
		return NewFixed8(ip * D), nil
	}

	dp, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil || dp >= D {
		return NewFixed8(0), fmt.Errorf("Fixed8 must satisfy following regex \\d+(\\.\\d{1,8})?")
	}
	for i := len(parts[1]); i < PRECISION; i++ {
		dp *= 10
	}
	if ip < 0 {
		return NewFixed8(ip*D - dp), nil
	}
	return NewFixed8(ip*D + dp), nil
}

// Fixed8ToString converts a Fixed8 to a decimal string
func Fixed8ToString(input Fixed8) string {
	buf := new(strings.Builder)
	val := input.Value
	if val < 0 {
		buf.WriteRune('-')
		val = -val
	}
	ip := strconv.FormatInt(val/D, 10)
	buf.WriteString(ip)
	val %= D
	if val > 0 {
		buf.WriteRune('.')
		dp := strconv.FormatInt(val, 10)
		for i := len(dp); i < 8; i++ {
			buf.WriteRune('0')
		}
		buf.WriteString(strings.TrimRight(dp, "0"))
	}
	return buf.String()
}

// String implements the Stringer interface.
func (f Fixed8) String() string {
	buf := new(strings.Builder)
	val := f.Value
	if val < 0 {
		buf.WriteRune('-')
		val = -val
	}
	str := strconv.FormatInt(val/D, 10)
	buf.WriteString(str)
	val %= D
	if val > 0 {
		buf.WriteRune('.')
		str = strconv.FormatInt(val, 10)
		for i := len(str); i < 8; i++ {
			buf.WriteRune('0')
		}
		buf.WriteString(strings.TrimRight(str, "0"))
	}
	return buf.String()
}

// Add implements Fixed8 + operator.
func (f Fixed8) Add(g Fixed8) Fixed8 {
	return NewFixed8(f.Value + g.Value)
}

// Sub implements Fixed8 - operator
func (f Fixed8) Sub(g Fixed8) Fixed8 {
	return NewFixed8(f.Value - g.Value)
}

// Mul implements Fixed8 * operator
func (f Fixed8) Mul(g Fixed8) (Fixed8, error) {
	var QUO uint64 = (1 << 63) / (D >> 1)
	var REM uint64 = ((1 << 63) % (D >> 1)) << 1

	fv := constant.MakeInt64(f.Value)
	gv := constant.MakeInt64(g.Value)
	sign := constant.Sign(fv) * constant.Sign(gv)

	ux := uint64(f.Abs().Value)
	uy := uint64(g.Abs().Value)
	xh := ux >> 32
	xl := ux & 0x00000000ffffffff
	yh := uy >> 32
	yl := uy & 0x00000000ffffffff
	rh := xh * yh
	rm := xh*yl + xl*yh
	rl := xl * yl
	rmh := rm >> 32
	rml := rm << 32
	rh += rmh
	rl += rml
	if rl < rml {
		rh++
	}
	if rh >= D {
		return Fixed8{}, fmt.Errorf("overflow error")
	}
	rd := rh*REM + rl
	if rd < rl {
		rh++
	}
	r := rh*QUO + rd/D
	result := int64(r) * int64(sign)
	return NewFixed8(result), nil
}

// Div implements Fixed8 / operator.
func (f Fixed8) Div(g Fixed8) Fixed8 {
	return NewFixed8(f.Value / g.Value)
}

// GreaterThan implements Fixed8 > operator.
func (f Fixed8) GreaterThan(g Fixed8) bool {
	return f.Value > g.Value
}

// LessThan implements Fixed8 < operator.
func (f Fixed8) LessThan(g Fixed8) bool {
	return f.Value < g.Value
}

// Equal implements Fixed8 == operator.
func (f Fixed8) Equal(g Fixed8) bool {
	return f.Value == g.Value
}

// CompareTo returns the difference between the f and g.
// difference < 0 implies f < g.
// difference = 0 implies f = g.
// difference > 0 implies f > g.
func (f Fixed8) CompareTo(g Fixed8) int {
	return int(f.Value - g.Value)
}

// Abs returns the absolute value of a Fixed8 type
func (f Fixed8) Abs() Fixed8 {
	if f.Value >= 0 {
		return f
	} else {
		return NewFixed8(-f.Value)
	}
}

// Ceiling returns the ceiling value of a Fixed8 type
func (f Fixed8) Ceiling() Fixed8 {
	var remainder int64 = f.Value % D
	if remainder > 0 {
		return NewFixed8(f.Value - remainder + D)
	} else {
		return NewFixed8(f.Value - remainder)
	}
}
