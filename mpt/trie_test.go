package mpt

import (
	"testing"

	"github.com/joeqian10/neo-gogogo/helper"
)

func TestVerifyProof(t *testing.T) {
	proof := []string{"01020a0c20ee990773274d7eaac0ae706e2ee5adc248339f5ff8a76c3c98481018c74528da", "00200acbf6b69f84623af9293ea80a3c4c4bb287df1423f24d25a55e168cef6d7fb20000000000000000202cd7c543593a9f1bfb73d816c51212e16988792328be33bdb9aa2db27c0b0bf820704e06d687922e7e3ac845053900d85140a087df4086f27d8caac4f9b6479f1d000000000000", "01010120ef1debfaddd605880af0f9cb04d30bb01f57ba4f4fd54b2c2164e33847add3d2", "0302abcd"}
	root := "eb2c5e3c8f16ffcc82d1fb157f496a517187e9812a6a9c62cade3449e8d86824"
	path := "ac01"
	value, err := VerifyProof(helper.HexTobytes(root), helper.HexTobytes(path), proof)
	if err != nil {
		t.Error("verify proof err:", err)
	}
	if helper.BytesToHex(value) != "abcd" {
		t.Error("wrong value")
	}
}
