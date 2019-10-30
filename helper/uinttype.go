package helper

import (
	"bytes"
	"fmt"
)

// UInt256 is a type of hash value in neo. It's an uint256 big number with a length of 32 byte
// PublicKey, BlockHash, TransactionHash and other hashes are in UInt256 format
type UInt256 struct {
	Data []byte
}

func (value UInt256) IsValid() bool {
	return len(value.Data) == 32
}

func NewUInt256(b []byte) (UInt256, error) {
	if len(b) != 32 {
		return UInt256{}, fmt.Errorf("Invalid data length.")
	}
	return UInt256{Data: b}, nil
}

func (value UInt256) Equal(other UInt256) bool {
	return bytes.Equal(value.Data, other.Data)
}

// UInt160 is a type of script hash value in neo. Usually used as an account address.
// The address in string format is Base58 coded script hash with some checksum bytes.
type UInt160 struct {
	Data []byte
}

func (value UInt160) IsValid() bool {
	return len(value.Data) == 20
}

func NewUInt160(b []byte) (UInt160, error) {
	if len(b) != 20 {
		return UInt160{}, fmt.Errorf("Invalid data length.")
	}
	return UInt160{Data: b}, nil
}
