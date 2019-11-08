package wallet

import (
	"encoding/json"
	"github.com/joeqian10/neo-gogogo/wallet/keys"
	"os"
)

const (
	// The current version of neo-go wallet implementations.
	walletVersion = "1.0"
)

// Wallet represents a NEO (NEP-2, NEP-6) compliant wallet.
type Wallet struct {
	// string type
	Name interface{} `json:"name"`

	// Version of the wallet, used for later upgrades.
	Version string `json:"version"`

	Scrypt ScryptParams `json:"scrypt"`

	// A list of accounts which describes the details of each account
	// in the wallet.
	Accounts []*Account `json:"accounts"`

	// Extra metadata can be used for storing arbitrary data.
	// This field can be empty.
	Extra interface{} `json:"extra"`
}

// ScryptParams is a json-serializable container for scrypt KDF parameters.
type ScryptParams struct {
	N int `json:"n"`
	R int `json:"r"`
	P int `json:"p"`
}

// NewWallet creates a NEO wallet.
func NewWallet() *Wallet {
	return &Wallet{
		Version:  walletVersion,
		Accounts: []*Account{},
		Scrypt:   ScryptParams{keys.N, keys.R, keys.P},
	}
}

// NewWalletFromFile creates a Wallet from the given wallet file path
func NewWalletFromFile(path string) (*Wallet, error) {
	file, err := os.OpenFile(path, os.O_RDWR, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	wall := &Wallet{}
	if err := json.NewDecoder(file).Decode(wall); err != nil {
		return nil, err
	}
	return wall, nil
}

// CreateAccount generates a new account for the end user and encrypts
// the private key with the given passphrase.
func (w *Wallet) CreateNewAccount(name, passphrase string) error {
	acc, err := NewAccount()
	if err != nil {
		return err
	}
	acc.Label = name
	if err := acc.Encrypt(passphrase); err != nil {
		return err
	}
	w.AddAccount(acc)
	return nil
}

// AddAccount adds an existing Account to the wallet.
func (w *Wallet) AddAccount(acc *Account) {
	w.Accounts = append(w.Accounts, acc)
}

// Save saves the wallet data. It's the internal io.ReadWriter
// that is responsible for saving the data. This can
// be a buffer, file, etc..
func (w *Wallet) Save(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	err = json.NewEncoder(file).Encode(w)
	if err != nil {
		return err
	}
	return file.Close()
}

// JSON outputs a pretty JSON representation of the wallet.
func (w *Wallet) JSON() ([]byte, error) {
	return json.Marshal(w)
}
