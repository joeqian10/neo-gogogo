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

// ReverseBytes ...
func ReverseBytes(b []byte) []byte {
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}
	return b
}

// Converts big endian human readable hex string format scripthash to little endian byte array
func ScriptHashToBytes(scriptHash string) []byte {
	return ReverseBytes(HexTobytes(scriptHash))
}

// Converts little endian byte array to big endian human readable hex string format scripthash
func BytesToScriptHash(ba []byte) string {
	return BytesToHex(ReverseBytes(ba))
}

func ScriptHashToAddress(scriptHash UInt256) string {
	var addressVersion byte = 0x17
	data := append([]byte{addressVersion}, scriptHash.Data...)
	return crypto.Base58CheckEncode(data)
}

func AddressToScriptHash(address string) (UInt160, error) {
	data, err := crypto.Base58CheckDecode(address)
	if err != nil {
		u, _ := NewUInt160([]byte{})
		return u, err
	}
	if data == nil || len(data) != 21 || data[0] != 0x17 {
		u, _ := NewUInt160([]byte{})
		return u, fmt.Errorf("Invalid address string.")
	}
	return NewUInt160(data[1:])
}

// ReverseString
func ReverseString(input string) string {
	return BytesToHex(ReverseBytes(HexTobytes(input)))
}

// Uint32ToBytes ...
func Uint32ToBytes(n uint32) []byte {
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
