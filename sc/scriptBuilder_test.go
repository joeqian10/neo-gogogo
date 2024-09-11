package sc

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joeqian10/neo-gogogo/helper"
)

func TestScriptBuilder_Emit(t *testing.T) {
	sb := NewScriptBuilder()
	scriptHash, _ := helper.UInt160FromString("14df5d02f9a52d3e92ab8cdcce5fc76c743a9b26")
	_ = sb.Emit(APPCALL, scriptHash.Bytes()...)
	b := sb.ToArray()
	assert.Equal(t, "67269b3a746cc75fcedc8cab923e2da5f9025ddf14", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitAppCall(t *testing.T) {
	sb := NewScriptBuilder()
	scriptHash, _ := helper.UInt160FromString("14df5d02f9a52d3e92ab8cdcce5fc76c743a9b26")
	_ = sb.EmitAppCall(scriptHash.Bytes(), false)
	b := sb.ToArray()
	assert.Equal(t, "67269b3a746cc75fcedc8cab923e2da5f9025ddf14", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitJump(t *testing.T) {
	sb := NewScriptBuilder()
	_ = sb.EmitJump(JMP, 77)
	b := sb.ToArray()
	assert.Equal(t, "624d00", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBigInt(t *testing.T) {
	sb := NewScriptBuilder()
	_ = sb.EmitPushBigInt(*big.NewInt(221))
	b := sb.ToArray()
	assert.Equal(t, "02dd00", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBigInt2(t *testing.T) {
	sb := NewScriptBuilder()
	_ = sb.EmitPushBigInt(*big.NewInt(345))
	b := sb.ToArray()
	assert.Equal(t, "025901", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBool(t *testing.T) {
	sb := NewScriptBuilder()
	_ = sb.EmitPushBool(true)
	b := sb.ToArray()
	assert.Equal(t, "51", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBytes(t *testing.T) {
	sb := NewScriptBuilder()
	n := *big.NewInt(7777777777)
	bytes := helper.ReverseBytes(n.Bytes())
	_ = sb.EmitPushBytes(bytes)
	b := sb.ToArray()
	assert.Equal(t, "05717897cf01", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushInt(t *testing.T) {
	sb := NewScriptBuilder()
	_ = sb.EmitPushInt(-1)
	_ = sb.EmitPushInt(0)
	_ = sb.EmitPushInt(8)
	_ = sb.EmitPushInt(100)
	_ = sb.EmitPushInt(1000)
	_ = sb.EmitPushInt(10000)
	_ = sb.EmitPushInt(0x20000)
	bytes := sb.ToArray()
	assert.Equal(t, "4f0058016402e80302102703000002", helper.BytesToHex(bytes))
}

func TestScriptBuilder_EmitPushParameter(t *testing.T) {
	u, _ := helper.UInt256FromString("c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b")
	cp := ContractParameter{
		Type:  Hash256,
		Value: u.Bytes(),
	}
	sb := NewScriptBuilder()
	_ = sb.EmitPushParameter(cp)
	b := sb.ToArray()
	assert.Equal(t, "209b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc5", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushString(t *testing.T) {
	sb := NewScriptBuilder()
	_ = sb.EmitPushString("Hello World!")
	b := sb.ToArray()
	assert.Equal(t, "0c48656c6c6f20576f726c6421", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitVmSysCall(t *testing.T) {
	sb := NewScriptBuilder()
	_ = sb.EmitVmSysCall("syscall", false)
	b := sb.ToArray()
	assert.Equal(t, "680773797363616c6c", helper.BytesToHex(b))

	sb = NewScriptBuilder()
	_ = sb.EmitVmSysCall("syscall", true)
	b = sb.ToArray()
	assert.Equal(t, "680444b1bb13", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitSysCall(t *testing.T) {
	script := []byte{0x01, 0x02, 0x03, 0x04}
	paramTypes := "0710"
	returnTypeHexString := "05"
	hasStorage := true
	hasDynamicInvoke := true
	isPayable := true
	var contractName = "test"
	var contractVersion = "1.0"
	var contractAuthor = "ngd"
	var contractEmail = "test@ngd.neo.org"
	var contractDescription = "cd"

	parameterList := helper.HexToBytes(paramTypes)                                 // 0710
	returnType := ContractParameterType(helper.HexToBytes(returnTypeHexString)[0]) // 05
	property := NoProperty
	if hasStorage {
		property |= HasStorage
	}
	if hasDynamicInvoke {
		property |= HasDynamicInvoke
	}
	if isPayable {
		property |= Payable
	}
	p1 := ContractParameter{
		Type:  ByteArray,
		Value: script,
	}
	p2 := ContractParameter{
		Type:  ByteArray,
		Value: parameterList,
	}
	p3 := ContractParameter{
		Type:  Integer,
		Value: *big.NewInt(int64(returnType)),
	}
	p4 := ContractParameter{
		Type:  Integer,
		Value: *big.NewInt(int64(property)),
	}
	p5 := ContractParameter{
		Type:  String,
		Value: contractName,
	}
	p6 := ContractParameter{
		Type:  String,
		Value: contractVersion,
	}
	p7 := ContractParameter{
		Type:  String,
		Value: contractAuthor,
	}
	p8 := ContractParameter{
		Type:  String,
		Value: contractEmail,
	}
	p9 := ContractParameter{
		Type:  String,
		Value: contractDescription,
	}
	sb := NewScriptBuilder()
	sb.EmitSysCall("Neo.Contract.Create", []ContractParameter{p1, p2, p3, p4, p5, p6, p7, p8, p9})
	newScript := sb.ToArray()
	assert.Equal(t, "0263641074657374406e67642e6e656f2e6f7267036e676403312e300474657374575502071004010203046804f66ca56e", helper.BytesToHex(newScript))
}

func TestScriptBuilder_MakeInvocationScript(t *testing.T) {
	sb := NewScriptBuilder()
	scriptHash, _ := helper.UInt160FromString("b9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
	sb.MakeInvocationScript(scriptHash.Bytes(), "name", []ContractParameter{})
	b := sb.ToArray()
	assert.Equal(t, "00c1046e616d656763d26113bac4208254d98a3eebaee66230ead7b9", helper.BytesToHex(b))
}

func TestScriptBuilder_MakeInvocationScript2(t *testing.T) {
	sb := NewScriptBuilder()
	scriptHash, _ := helper.UInt160FromString("0x43bb08d7c03ac66582079b57059108565f91ece5")
	f, _ := helper.AddressToScriptHash("AKeLhhHm4hEUfLWVBCYRNjio9xhGJAom5G")
	to, _ := helper.AddressToScriptHash("AdmyedL3jdw2TLvBzoUD2yU443NeKrP5t5")

	cp1 := ContractParameter{
		Type:  Hash160,
		Value: f.Bytes(),
	}
	cp2 := ContractParameter{
		Type:  Hash160,
		Value: to.Bytes(),
	}
	cp3 := ContractParameter{
		Type:  Integer,
		Value: *big.NewInt(int64(20000000000)),
	}

	sb.MakeInvocationScript(scriptHash.Bytes(), "transfer", []ContractParameter{cp1, cp2, cp3})
	b := sb.ToArray()
	assert.Equal(t, "0500c817a80414f157c713c1972ba426ceb4c2b10826e54047d522142a73c28a1e57d7bbd212d598715194690e29d8bc53c1087472616e7366657267e5ec915f56089105579b078265c63ac0d708bb43", helper.BytesToHex(b))
}

func TestScriptBuilder_ToArray(t *testing.T) {
	sb := NewScriptBuilder()
	_ = sb.EmitPushString("Hello World!")
	b := sb.ToArray()
	assert.Equal(t, "0c48656c6c6f20576f726c6421", helper.BytesToHex(b))
}
