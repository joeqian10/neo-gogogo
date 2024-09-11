package sc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/joeqian10/neo-gogogo/crypto"
	"github.com/joeqian10/neo-gogogo/helper"
)

type ScriptBuilder struct {
	buff *bytes.Buffer
}

func NewScriptBuilder() ScriptBuilder {
	return ScriptBuilder{buff: new(bytes.Buffer)}
}

func (sb *ScriptBuilder) ToArray() []byte {
	return sb.buff.Bytes()
}

func (sb *ScriptBuilder) MakeInvocationScript(scriptHash []byte, operation string, args []ContractParameter) {
	if len(operation) == 0 { // Neo.VM.Helper.cs: Line 28
		l := len(args)
		for i := l - 1; i >= 0; i-- {
			sb.EmitPushParameter(args[i])
		}
		sb.EmitAppCall(scriptHash, false)
	} else {
		if args != nil { // Neo.VM.Helper.cs: Line 43
			l := len(args)
			for i := l - 1; i >= 0; i-- {
				sb.EmitPushParameter(args[i])
			}
			sb.EmitPushInt(l)
			sb.Emit(PACK)
			sb.EmitPushString(operation)
			sb.EmitAppCall(scriptHash, false)
		} else { // Neo.VM.Helper.cs: Line 35
			sb.EmitPushBool(false)
			sb.EmitPushString(operation)
			sb.EmitAppCall(scriptHash, false)
		}
	}
}

func (sb *ScriptBuilder) Emit(op OpCode, arg ...byte) error {
	err := sb.buff.WriteByte(byte(op))
	if err != nil {
		return err
	}

	if arg != nil {
		_, err = sb.buff.Write(arg)
	}
	return err
}

func (sb *ScriptBuilder) EmitAppCall(scriptHah []byte, useTailCall bool) error {
	if len(scriptHah) != 20 {
		return fmt.Errorf("the length of scripthash should be 20")
	}
	if useTailCall {
		return sb.Emit(TAILCALL, scriptHah...)
	} else {
		return sb.Emit(APPCALL, scriptHah...)
	}
}

func (sb *ScriptBuilder) EmitJump(op OpCode, offset int16) error {
	if op != JMP && op != JMPIF && op != JMPIFNOT && op != CALL {
		return fmt.Errorf("invalid OpCode")
	}
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(offset))
	//v := helper.VarIntFromInt16(offset)
	return sb.Emit(op, b...)
}

func (sb *ScriptBuilder) EmitPushBigInt(number big.Int) error {
	if number.Cmp(big.NewInt(-1)) == 0 {
		return sb.Emit(PUSHM1)
	}
	if number.Cmp(big.NewInt(0)) == 0 {
		return sb.Emit(PUSH0)
	}
	if number.Cmp(big.NewInt(0)) > 0 && number.Cmp(big.NewInt(16)) <= 0 {
		var b = byte(number.Int64())
		return sb.Emit(PUSH1 - 1 + OpCode(b))
	}
	// need little endian
	b := helper.BigIntToNeoBytes(&number)
	return sb.EmitPushBytes(b)
}

func (sb *ScriptBuilder) EmitPushInt(number int) error {
	return sb.EmitPushBigInt(*big.NewInt(int64(number)))
}

func (sb *ScriptBuilder) EmitPushBool(data bool) error {
	if data {
		return sb.Emit(PUSHT)
	} else {
		return sb.Emit(PUSHF)
	}
}

func (sb *ScriptBuilder) EmitPushBytes(data []byte) error {
	if data == nil {
		return fmt.Errorf("data is empty")
	}
	le := len(data)
	v := helper.VarIntFromInt(le)
	var err error
	if le <= int(PUSHBYTES75) {
		sb.buff.WriteByte(byte(le))
		sb.buff.Write(data)
	} else if le < int(0x100) {
		err = sb.Emit(PUSHDATA1)
		sb.buff.WriteByte(byte(le))
		sb.buff.Write(data)
	} else if le < int(0x10000) {
		err = sb.Emit(PUSHDATA2)
		sb.buff.Write(v.Bytes()[1:])
		sb.buff.Write(data)
	} else {
		err = sb.Emit(PUSHDATA4)
		sb.buff.Write(v.Bytes()[1:])
		sb.buff.Write(data)
	}
	if err != nil {
		return err
	}
	return nil
}

// Convert the string to UTF8 format encoded byte array
func (sb *ScriptBuilder) EmitPushString(data string) error {
	return sb.EmitPushBytes([]byte(data))
}

func (sb *ScriptBuilder) EmitPushParameter(data ContractParameter) error {
	var err error
	switch data.Type {
	case Signature:
	case ByteArray:
		err = sb.EmitPushBytes(data.Value.([]byte))
	case Boolean:
		err = sb.EmitPushBool(data.Value.(bool))
	case Integer:
		num := data.Value.(big.Int)
		err = sb.EmitPushBigInt(num)
	case Hash160:
		u, e := helper.UInt160FromBytes(data.Value.([]byte))
		if e != nil {
			return e
		}
		err = sb.EmitPushBytes(u.Bytes())
	case Hash256:
		u, e := helper.UInt256FromBytes(data.Value.([]byte))
		if e != nil {
			return e
		}
		err = sb.EmitPushBytes(u.Bytes())
	case PublicKey:
		err = sb.EmitPushBytes(data.Value.([]byte))
	case String:
		s := string(data.Value.(string))
		err = sb.EmitPushString(s)
	case Array:
		a := data.Value.([]ContractParameter)
		for i := len(a) - 1; i >= 0; i-- {
			e := sb.EmitPushParameter(a[i])
			if e != nil {
				return e
			}
		}
		err = sb.EmitPushInt(len(a))
		if err != nil {
			return err
		}
		err = sb.Emit(PACK)
	}
	if err != nil {
		return err
	}
	return nil
}

func (sb *ScriptBuilder) EmitSysCall(api string, args []ContractParameter) error {
	var err error
	l := len(args)
	for i := l - 1; i >= 0; i-- {
		err = sb.EmitPushParameter(args[i])
	}
	if err != nil {
		return err
	}
	return sb.EmitVmSysCall(api, true)
}

func (sb *ScriptBuilder) EmitVmSysCall(api string, compress bool) error {
	if len(api) == 0 {
		return fmt.Errorf("argument api is empty")
	}
	b := []byte(api)
	if compress {
		b = crypto.Sha256(b)
		b = b[0:4]
	} else {
		if len(b) > 252 {
			return fmt.Errorf("argument api has a too long length")
		}
	}
	arg := append([]byte{byte(len(b))}, b...)
	return sb.Emit(SYSCALL, arg...)
}
