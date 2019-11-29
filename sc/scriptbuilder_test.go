package sc

import (
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)



func TestScriptBuilder_Emit(t *testing.T) {
	sb := NewScriptBuilder()
	scriptHash, _ := helper.UInt160FromString("14df5d02f9a52d3e92ab8cdcce5fc76c743a9b26")
	sb.Emit(APPCALL, scriptHash.Bytes()...)
	b := sb.ToArray()
	assert.Equal(t, "67269b3a746cc75fcedc8cab923e2da5f9025ddf14", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitAppCall(t *testing.T) {
	sb := NewScriptBuilder()
	scriptHash, _ := helper.UInt160FromString("14df5d02f9a52d3e92ab8cdcce5fc76c743a9b26")
	sb.EmitAppCall(scriptHash.Bytes(), false)
	b := sb.ToArray()
	assert.Equal(t, "67269b3a746cc75fcedc8cab923e2da5f9025ddf14", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitJump(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitJump(JMP, 77)
	b := sb.ToArray()
	assert.Equal(t, "624d00", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBigInt(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushBigInt(*big.NewInt(7777777777))
	b := sb.ToArray()
	assert.Equal(t, "05717897cf01", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBool(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushBool(true)
	b := sb.ToArray()
	assert.Equal(t, "51", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushBytes(t *testing.T) {
	sb := NewScriptBuilder()
	n := *big.NewInt(7777777777)
	bytes := helper.ReverseBytes(n.Bytes())
	sb.EmitPushBytes(bytes)
	b := sb.ToArray()
	assert.Equal(t, "05717897cf01", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushInt(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushInt(-1)
	sb.EmitPushInt(0)
	sb.EmitPushInt(8)
	sb.EmitPushInt(100)
	sb.EmitPushInt(1000)
	sb.EmitPushInt(10000)
	sb.EmitPushInt(0x20000)
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
	sb.EmitPushParameter(cp)
	b := sb.ToArray()
	assert.Equal(t, "209b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc5", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitPushString(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushString("Hello World!")
	b := sb.ToArray()
	assert.Equal(t, "0c48656c6c6f20576f726c6421", helper.BytesToHex(b))
}

func TestScriptBuilder_EmitSysCall(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitSysCall("syscall", false)
	b := sb.ToArray()
	assert.Equal(t, "680773797363616c6c", helper.BytesToHex(b))

	sb = NewScriptBuilder()
	sb.EmitSysCall("syscall", true)
	b = sb.ToArray()
	assert.Equal(t, "680444b1bb13", helper.BytesToHex(b))
}

func TestScriptBuilder_MakeInvocationScript(t *testing.T) {
	sb := NewScriptBuilder()
	scriptHash, _ := helper.UInt160FromString("b9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
	sb.MakeInvocationScript(scriptHash.Bytes(), "name", []ContractParameter{})
	b := sb.ToArray()
	assert.Equal(t, "00c1046e616d656763d26113bac4208254d98a3eebaee66230ead7b9", helper.BytesToHex(b))
}

func TestScriptBuilder_ToArray(t *testing.T) {
	sb := NewScriptBuilder()
	sb.EmitPushString("Hello World!")
	b := sb.ToArray()
	assert.Equal(t, "0c48656c6c6f20576f726c6421", helper.BytesToHex(b))
}
