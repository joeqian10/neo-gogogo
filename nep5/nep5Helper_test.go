package nep5

import (
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNep5Helper(t *testing.T) {
	nep5helper := NewNep5Helper("http://seed1.ngd.network:20332")
	if nep5helper == nil {
		t.Fail()
	}
	assert.Equal(t, "seed1.ngd.network:20332", nep5helper.Client.Endpoint.Host)
}

func TestNep5Helper_Name(t *testing.T) {

	name := string(helper.HexTobytes("516c696e6b20546f6b656e"))
	assert.Equal(t, "Qlink Token", name)
}