package wallet

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestNewWalletFromFile(t *testing.T) {
	path := "test.json"
	testWallet, err := NewWalletFromFile(path)
	assert.Nil(t, err)
	assert.Equal(t, len(testWallet.Accounts), 5)

	jsonBytes, err := testWallet.JSON()
	assert.Nil(t, err)

	data, err := ioutil.ReadFile(path)
	assert.Nil(t, err)

	jsonString := string(jsonBytes)
	assert.Equal(t, string(data), jsonString)

}

func TestWallet_Save(t *testing.T) {
	path := "test.json"
	testWallet, err := NewWalletFromFile(path)
	assert.Nil(t, err)

	err = testWallet.Save("testWrite.json")
	assert.Nil(t, err)
	testWrite, err := NewWalletFromFile("testWrite.json")

	assert.Nil(t, err)
	assert.Equal(t, testWallet, testWrite)
}
