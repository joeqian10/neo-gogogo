package helper

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/joeqian10/neo-gogogo/crypto"
)

//BytesToHex bytes to hex string
func BytesToHex(b []byte) string {
	return hex.EncodeToString(b)
}

//HexToBytes Simple hex string to bytes
func HexToBytes(hexstring string) (b []byte) {
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

//ToNibbles ..
func ToNibbles(data []byte) []byte {
	r := make([]byte, len(data)*2)
	for i := 0; i < len(data); i++ {
		r[i*2] = data[i] >> 4
		r[i*2+1] = data[i] & 0x0f
	}
	return r
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
	return BytesToHex(ReverseBytes(HexToBytes(input)))
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

func Abs(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -x
}
