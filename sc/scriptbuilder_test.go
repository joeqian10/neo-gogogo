package sc

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScriptBuilder_EmitPushInt(t *testing.T) {
	builder := NewScriptBuilder()
	builder.EmitPushInt(-1)
	builder.EmitPushInt(0)
	builder.EmitPushInt(8)
	builder.EmitPushInt(100)
	builder.EmitPushInt(1000)
	builder.EmitPushInt(10000)
	builder.EmitPushInt(0x20000)
	bytes := builder.ToArray()
	assert.Equal(t, "4f0058016402e80302102703000002", hex.EncodeToString(bytes))
}
