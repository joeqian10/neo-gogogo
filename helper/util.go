package helper

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/joeqian10/neo-gogogo/crypto"
	"math/big"
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

func BytesToScriptHash(bs []byte) (UInt160, error) {
	return UInt160FromBytes(crypto.Hash160(bs))
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

func BytesToUInt64(bs []byte) uint64 {
	bs = PadRight(bs, 8)
	return binary.LittleEndian.Uint64(bs)
}

func BytesToUInt32(bs []byte) uint32 {
	bs = PadRight(bs, 4)
	return binary.LittleEndian.Uint32(bs)
}

func PadRight(data []byte, length int) []byte {
	if len(data) >= length {
		return data[:length] // return the most left bytes of length
	}
	newData := data
	for len(newData) < length {
		newData = append(newData, byte(0))
	}
	return newData
}

func BigIntToNeoBytes(data *big.Int) []byte {
	bs := data.Bytes()
	if len(bs) == 0 {
		return []byte{}
	}
	// golang big.Int use big-endian
	bs = ReverseBytes(bs)
	// bs now is little-endian
	if data.Sign() < 0 {
		for i, b := range bs {
			bs[i] = ^b
		}
		for i := 0; i < len(bs); i++ {
			if bs[i] == 255 {
				bs[i] = 0
			} else {
				bs[i] += 1
				break
			}
		}
		if bs[len(bs)-1] < 128 {
			bs = append(bs, 255)
		}
	} else {
		if bs[len(bs)-1] >= 128 {
			bs = append(bs, 0)
		}
	}
	return bs
}

var bigOne = big.NewInt(1)

func BigIntFromNeoBytes(ba []byte) *big.Int {
	res := big.NewInt(0)
	l := len(ba)
	if l == 0 {
		return res
	}

	bytes := make([]byte, 0, l)
	bytes = append(bytes, ba...)
	bytes = ReverseBytes(bytes)

	if bytes[0]>>7 == 1 {
		for i, b := range bytes {
			bytes[i] = ^b
		}

		temp := big.NewInt(0)
		temp.SetBytes(bytes)
		temp.Add(temp, bigOne)
		bytes = temp.Bytes()
		res.SetBytes(bytes)
		return res.Neg(res)
	}

	res.SetBytes(bytes)
	return res
}

//func HashToInt(hash []byte) *big.Int {
//	orderBits := 256
//	orderBytes := (orderBits + 7) / 8
//	if len(hash) > orderBytes {
//		hash = hash[:orderBytes]
//	}
//
//	ret := new(big.Int).SetBytes(hash)
//	excess := len(hash)*8 - orderBits
//	if excess > 0 {
//		ret.Rsh(ret, uint(excess))
//	}
//	return ret
//}
