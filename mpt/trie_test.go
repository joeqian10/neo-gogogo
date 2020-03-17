package mpt

import (
	"testing"

	"github.com/joeqian10/neo-gogogo/blockchain"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/helper/io"
)

func TestVerifyProof1(t *testing.T) {
	proofStr := []string{"01020a0c20ee990773274d7eaac0ae706e2ee5adc248339f5ff8a76c3c98481018c74528da", "00200acbf6b69f84623af9293ea80a3c4c4bb287df1423f24d25a55e168cef6d7fb20000000000000000202cd7c543593a9f1bfb73d816c51212e16988792328be33bdb9aa2db27c0b0bf820704e06d687922e7e3ac845053900d85140a087df4086f27d8caac4f9b6479f1d000000000000", "01010120ef1debfaddd605880af0f9cb04d30bb01f57ba4f4fd54b2c2164e33847add3d2", "0302abcd"}
	proof := make([][]byte, len(proofStr))
	for i, v := range proofStr {
		proof[i] = helper.HexToBytes(v)
	}
	root := "eb2c5e3c8f16ffcc82d1fb157f496a517187e9812a6a9c62cade3449e8d86824"
	path := "ac01"
	value, err := VerifyProof(helper.HexToBytes(root), helper.HexToBytes(path), proof)
	if err != nil {
		t.Error("verify proof err:", err)
	}
	if helper.BytesToHex(value) != "abcd" {
		t.Error("wrong value")
	}
}

func TestVerifyProof2(t *testing.T) {
	proofStr := []string{
		"000000000000000000000020dbe90d0674546fd0e6dc879013ab743b5f171cc1f4e97caab98be714bb7f125400208837fe543b1bfcd58d480ea9988cb7acf58c5f4cf0518a987f70924320ceab390020ff409ca133d8a3a5ea8b72d5b82214d6190d3dd44d19d67062681af2b4b830200000",
		"0128050f090b040f07000c0804010105040006000b0103020f0f010e0205030e0e0c0a0a070d0b080a06202da36d644929144a9869d0ffb581c645f4d29b7b38a068ca1cc3bbd34b0771e2",
		"000020a41c8cc1e18bcbe0f237fa00544455ab92d5a86bc8fa6cd0f0ba294e9f02202d002002ad9e7adebf3057801f6918a97efd49741949f8a538b58b4fae31700b976e7d00000000000000000000000000",
		"010a0703070306050704000020afb2d68b4ae87095c68d113f42234c3e36e15597ec1c11ecce4261f3a9169b0a",
		"00207c00cb1580a95b55ecf7e82829997ebf4176cf95df65712a0b14b37eb13c1c5f0000207e9c938de1d3f1f95753c0867e3326bdb16b2caccefb6f490810bf13ccc7440e00000000000000000000000000",
		"0137040c010e090f070401030e080a050e0b030c0500000f0f0d00060a0c0d04010604040d050d050a090600000000000000000000000000062046a12d0bfc2f3f1d9e18a51f55d32ba4855578c1954d277180addb5b69210c97",
		"03070004c071b50400",
	}

	proof := make([][]byte, len(proofStr))
	for i, v := range proofStr {
		proof[i] = helper.HexToBytes(v)
	}

	root, _ := helper.UInt256FromString("34db5a993a95e0db79efe8220bf142e5952056bb59834fe3b91fc1611ed4385e")

	scriptHash, _ := helper.UInt160FromString("8adba7caee53e2f12f130b065411840cf7b4f9a5")
	skey := blockchain.Storagekey{
		ScriptHash: scriptHash,
		Key:        helper.HexToBytes("61737365740034c1e9f7413e8a5eb3c5ffd06acd41644d5d5a96"),
	}
	skeyBytes, _ := io.ToArray(&skey)

	value, err := VerifyProof(root.Bytes(), skeyBytes, proof)
	if err != nil {
		t.Error("verify proof err:", err)
	}

	sItem := blockchain.StorageItem{}
	io.AsSerializable(&sItem, value)
	if helper.BytesToHex(sItem.Value) != "c071b504" {
		t.Error("wrong value")
	}
}
