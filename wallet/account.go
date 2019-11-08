package wallet

import (
	"github.com/joeqian10/neo-gogogo/wallet/keys"
)

// Account represents a NEO account. It holds the private and public key
// along with some metadata.
type Account struct {
	// NEO  KeyPair.
	KeyPair *keys.KeyPair `json:"-"`

	// Account import file.
	wif string

	// NEO public address.
	Address string `json:"address"`

	// Label is a label the user had made for this account. string
	Label interface{} `json:"label"`

	// Indicates whether the account is the default change account.
	Default bool `json:"isDefault"`

	// Indicates whether the account is locked by the user.
	// the client shouldn't spend the funds in a locked account.
	Locked bool `json:"lock"`

	// Encrypted WIF of the account also known as the key. string
	EncryptedWIF interface{} `json:"key"`

	// contract is a Contract object which describes the details of the contract.
	// This field can be null (for watch-only address).
	Contract *Contract `json:"contract"`

	// This field can be empty.
	Extra interface{} `json:"extra"`
}

// Contract represents a subset of the smartcontract to embed in the
// Account so it's NEP-6 compliant.
type Contract struct {
	// Script hash of the contract deployed on the blockchain.
	Script string `json:"script"`

	// A list of parameters used deploying this contract.
	Parameters []interface{} `json:"parameters"`

	// Indicates whether the contract has been deployed to the blockchain.
	Deployed bool `json:"deployed"`
}

// NewAccount creates a new Account with a random generated PrivateKey.
func NewAccount() (*Account, error) {
	priv, err := keys.GenerateKeyPair()
	if err != nil {
		return nil, err
	}
	return NewAccountFromKeyPair(priv), nil
}

// DecryptAccount decrypt the encryptedWIF with the given passphrase and
// return the decrypted Account.
func DecryptAccount(encryptedWIF, passphrase string) (*Account, error) {
	wif, err := keys.NEP2Decrypt(encryptedWIF, passphrase)
	if err != nil {
		return nil, err
	}
	return NewAccountFromKeyPair(wif), nil
}

// Encrypt encrypts the wallet's PrivateKey with the given passphrase
// under the NEP-2 standard.
func (a *Account) Encrypt(passphrase string) error {
	wif, err := keys.NEP2Encrypt(a.KeyPair, passphrase)
	if err != nil {
		return err
	}
	a.EncryptedWIF = wif
	return nil
}

// NewAccountFromWIF creates a new Account from the given WIF.
func NewAccountFromWIF(wif string) (*Account, error) {
	privKey, err := keys.NewKeyPairFromWIF(wif)
	if err != nil {
		return nil, err
	}
	return NewAccountFromKeyPair(privKey), nil
}

// newAccountFromPrivateKey created a wallet from the given PrivateKey.
func NewAccountFromKeyPair(p *keys.KeyPair) *Account {
	pubAddr := p.PublicKey.Address()
	wif := p.ExportWIF()

	a := &Account{
		KeyPair: p,
		Address: pubAddr,
		wif:     wif,
	}

	return a
}
