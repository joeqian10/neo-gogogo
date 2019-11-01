package helper

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

const uint256Size = 32

// UInt256 is a 32 byte long unsigned integer.
type UInt256 [uint256Size]uint8

// UInt256DecodeReverseString attempts to decode the given string (in LE representation) into an UInt256.
func UInt256DecodeReverseString(s string) (u UInt256, err error) {
	if len(s) != uint256Size*2 {
		return u, fmt.Errorf("expected string size of %d got %d", uint256Size*2, len(s))
	}
	b, err := hex.DecodeString(s)
	if err != nil {
		return u, err
	}
	return UInt256DecodeReverseBytes(b)
}

// UInt256DecodeBytes attempts to decode the given string (in BE representation) into an UInt256.
func UInt256DecodeBytes(b []byte) (u UInt256, err error) {
	if len(b) != uint256Size {
		return u, fmt.Errorf("expected []byte of size %d got %d", uint256Size, len(b))
	}
	copy(u[:], b)
	return u, nil
}

// UInt256DecodeReverseBytes attempts to decode the given string (in LE representation) into an UInt256.
func UInt256DecodeReverseBytes(b []byte) (u UInt256, err error) {
	b = ReverseBytes(b)
	return UInt256DecodeBytes(b)
}

// Bytes returns a byte slice representation of u.
func (u UInt256) Bytes() []byte {
	return u[:]
}

// Reverse reverses the UInt256 object
func (u UInt256) Reverse() UInt256 {
	res, _ := UInt256DecodeReverseBytes(u.Bytes())
	return res
}

// BytesReverse return a reversed byte representation of u.
func (u UInt256) BytesReverse() []byte {
	return ReverseBytes(u.Bytes())
}

// Equals returns true if both UInt256 values are the same.
func (u UInt256) Equals(other UInt256) bool {
	return u == other
}

// String implements the stringer interface.
func (u UInt256) String() string {
	return hex.EncodeToString(u.Bytes())
}

// ReverseString produces string representation of UInt256 with LE byte order.
func (u UInt256) ReverseString() string {
	return hex.EncodeToString(u.BytesReverse())
}

// UnmarshalJSON implements the json unmarshaller interface.
func (u *UInt256) UnmarshalJSON(data []byte) (err error) {
	var js string
	if err = json.Unmarshal(data, &js); err != nil {
		return err
	}
	js = strings.TrimPrefix(js, "0x")
	*u, err = UInt256DecodeReverseString(js)
	return err
}

// MarshalJSON implements the json marshaller interface.
func (u UInt256) MarshalJSON() ([]byte, error) {
	return []byte(`"0x` + u.ReverseString() + `"`), nil
}

// CompareTo compares two UInt256 with each other. Possible output: 1, -1, 0
//  1 implies u > other.
// -1 implies u < other.
//  0 implies  u = other.
func (u UInt256) CompareTo(other UInt256) int { return bytes.Compare(u[:], other[:]) }
