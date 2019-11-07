package wallet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWallet(t *testing.T) {
	path := "test.json"
	testWallet, err := NewWalletFromFile(path)

	assert.Nil(t, err)
	assert.Equal(t, len(testWallet.Accounts), 5)
}

func TestWallet_AddAccount(t *testing.T) {

}

func TestWallet_Close(t *testing.T) {

}

func TestWallet_CreateAccount(t *testing.T) {

}

func TestWallet_JSON(t *testing.T) {

}

func TestWallet_Path(t *testing.T) {

}

func TestWallet_Save(t *testing.T) {

}

func Test_newWallet(t *testing.T) {

}
