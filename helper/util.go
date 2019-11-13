package helper

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/joeqian10/neo-gogogo/crypto"
)

// bytes to hex string
func BytesToHex(b []byte) string {
	return hex.EncodeToString(b)
}

// Simple hex string to bytes
func HexTobytes(hexstring string) (b []byte) {
	b, _ = hex.DecodeString(hexstring)
	return b
}

// ConcatBytes ...
func ConcatBytes(b1 []byte, b2 []byte) []byte {
	var buffer bytes.Buffer //Buffer: length changeable, writable, readable
	buffer.Write(b1)
	buffer.Write(b2)
	return buffer.Bytes()
}

// ReverseBytes without change original slice
func ReverseBytes(data []byte) []byte {
	b := make([]byte, len(data))
	copy(b, data)
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}
	return b
}

func ScriptHashToAddress(scriptHash UInt160) string {
	var addressVersion byte = 0x17
	data := append([]byte{addressVersion}, scriptHash.Bytes()...)
	return crypto.Base58CheckEncode(data)
}

func AddressToScriptHash(address string) (UInt160, error) {
	data, err := crypto.Base58CheckDecode(address)
	var u UInt160
	if err != nil {
		return u, err
	}
	if data == nil || len(data) != 21 || data[0] != 0x17 {
		return u, fmt.Errorf("invalid address string")
	}
	return UInt160FromBytes(data[1:])
}

// ReverseString
func ReverseString(input string) string {
	return BytesToHex(ReverseBytes(HexTobytes(input)))
}

// UInt32ToBytes ...
func UInt32ToBytes(n uint32) []byte {
	var buff = make([]byte, 4)
	binary.LittleEndian.PutUint32(buff, n)
	return buff
}

// Int64ToBytes ...
func Int64ToBytes(n int64) []byte {
	var buff = make([]byte, 8)
	binary.LittleEndian.PutUint64(buff, uint64(n))
	return buff
}
