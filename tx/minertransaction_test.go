package tx

import (
	"encoding/hex"
	"github.com/joeqian10/neo-gogogo/helper/io"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinerTransaction(t *testing.T) {
	rawTx := "0000fcd30e22000001e72d286979ee6cb1b7e65dfddfb2e384100b8d148e7758de42e4168b71792c60c8000000000000001f72e68b4e39602912106d53b229378a082784b200"
	mtx := &MinerTransaction{
		Transaction:NewTransaction(),
	}
	// Deserialize
	mtx, err := mtx.FromHexString(rawTx)
	assert.Nil(t, err)
	assert.Equal(t, Miner_Transaction, mtx.Type)
	assert.IsType(t, &MinerTransaction{}, mtx)
	assert.Equal(t, 0, int(mtx.Version))
	assert.Equal(t, uint32(571397116), mtx.Nonce)
	assert.Equal(t, "a1f219dc6be4c35eca172e65e02d4591045220221b1543f1a4b67b9e9442c264", mtx.HashString())

	// Serialize
	buf := io.NewBufBinaryWriter()
	mtx.Serialize(buf.BinaryWriter)
	assert.Equal(t, nil, buf.Err)
	assert.Equal(t, rawTx, hex.EncodeToString(buf.Bytes()))
}
