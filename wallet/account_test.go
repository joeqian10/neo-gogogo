package wallet

import (
	"github.com/joeqian10/neo-gogogo/wallet/keys"
	"testing"
)

func TestNewAccount(t *testing.T) {
	for _, testCase := range keys.KeyCases {
		acc, err := NewAccountFromWIF(testCase.Wif)
		if err != nil {
			t.Fatal(err)
		}
		compareFields(t, testCase, acc)
	}
}

func TestDecryptAccount(t *testing.T) {
	for _, testCase := range keys.KeyCases {
		acc, err := DecryptAccount(testCase.Nep2key, testCase.Passphrase)
		if err != nil {
			t.Fatal(err)
		}
		compareFields(t, testCase, acc)
	}
}

func TestNewFromWif(t *testing.T) {
	for _, testCase := range keys.KeyCases {
		acc, err := NewAccountFromWIF(testCase.Wif)
		if err != nil {
			t.Fatal(err)
		}
		compareFields(t, testCase, acc)
	}
}

func compareFields(t *testing.T, tk keys.Ktype, acc *Account) {
	if want, have := tk.Address, acc.Address; want != have {
		t.Fatalf("expected %s got %s", want, have)
	}
	if want, have := tk.Wif, acc.wif; want != have {
		t.Fatalf("expected %s got %s", want, have)
	}
	if want, have := tk.PublicKey, acc.KeyPair.PublicKey.String(); want != have {
		t.Fatalf("expected %s got %s", want, have)
	}
	if want, have := tk.PrivateKey, acc.KeyPair.String(); want != have {
		t.Fatalf("expected %s got %s", want, have)
	}
}
