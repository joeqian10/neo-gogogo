package tx

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joeqian10/neo-gogogo/helper/io"
)

func TestStateTransaction(t *testing.T) {
	// transaction taken from testnet 8abf5ebdb9a8223b12109513647f45bd3c0a6cf1a6346d56684cff71ba308724
	rawTx := "900001482103c089d7122b840a4935234e82e26ae5efd0c2acb627239dc9f207311337b6f2c10a5265676973746572656401010001cb4184f0a96e72656c1fbdd4f75cca567519e909fd43cefcec13d6c6abcb92a1000001e72d286979ee6cb1b7e65dfddfb2e384100b8d148e7758de42e4168b71792c6000b8fb050109000071f9cf7f0ec74ec0b0f28a92b12e1081574c0af00141408780d7b3c0aadc5398153df5e2f1cf159db21b8b0f34d3994d865433f79fafac41683783c48aef510b67660e3157b701b9ca4dd9946a385d578fba7dd26f4849232103c089d7122b840a4935234e82e26ae5efd0c2acb627239dc9f207311337b6f2c1ac"
	stx := &StateTransaction{
		Transaction: NewTransaction(),
	}
	// Deserialize
	stx, err := stx.FromHexString(rawTx)
	assert.Nil(t, err)
	assert.Equal(t, State_Transaction, stx.Type)
	assert.IsType(t, &StateTransaction{}, stx)
	assert.Equal(t, "8abf5ebdb9a8223b12109513647f45bd3c0a6cf1a6346d56684cff71ba308724", stx.HashString())

	assert.Equal(t, 1, len(stx.Inputs))
	input := stx.Inputs[0]
	assert.Equal(t, "a192cbabc6d613ecfcce43fd09e9197556ca5cf7d4bd1f6c65726ea9f08441cb", input.PrevHash.String())
	assert.Equal(t, uint16(0), input.PrevIndex)

	assert.Equal(t, 1, len(stx.Descriptors))
	descriptor := stx.Descriptors[0]
	assert.Equal(t, "03c089d7122b840a4935234e82e26ae5efd0c2acb627239dc9f207311337b6f2c1", hex.EncodeToString(descriptor.Key))
	assert.Equal(t, "52656769737465726564", hex.EncodeToString(descriptor.Value))
	assert.Equal(t, "\x01", descriptor.Field)
	assert.Equal(t, Validator, descriptor.Type)

	// Serialize
	buf := io.NewBufBinaryWriter()
	stx.Serialize(buf.BinaryWriter)

	assert.Equal(t, nil, buf.Err)
	assert.Equal(t, rawTx, hex.EncodeToString(buf.Bytes()))
}
