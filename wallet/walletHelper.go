package wallet

import (
	"fmt"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/sc"
	"github.com/joeqian10/neo-gogogo/tx"
)

type WalletHelper struct{
	TxBuilder *tx.TransactionBuilder
	Account *Account
}

func NewWalletHelper(txBuilder *tx.TransactionBuilder, account *Account) *WalletHelper {
	return &WalletHelper{
		TxBuilder: txBuilder,
		Account:   account,
	}
}

// Transfer is used to transfer neo or gas or other utxo asset, single signature
func (w *WalletHelper) Transfer(assetId helper.UInt256, from string, to string, amount float64) (bool, error) {
	f, err := helper.AddressToScriptHash(from)
	if err != nil {return false, err}
	t, err := helper.AddressToScriptHash(to)
	if err != nil {return false, err}
	a := helper.Fixed8FromFloat64(amount)
	ctx, err := w.TxBuilder.MakeContractTransaction(f, t, assetId, a, nil, helper.UInt160{}, helper.Zero)
	if err != nil {return false, err}
	// sign
	err = tx.AddSignature(ctx, w.Account.KeyPair)
	if err != nil {return false, err}
	// use RPC to send the tx
	response := w.TxBuilder.Client.SendRawTransaction(ctx.RawTransactionString())
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {return false, fmt.Errorf(msg)}
	return response.Result, nil
}

// ClaimGas
func (w *WalletHelper) ClaimGas(from string) (bool, error) {
	f, err := helper.AddressToScriptHash(from)
	if err != nil {return false, err}
	ctx, err := w.TxBuilder.MakeClaimTransaction(f, helper.UInt160{}, nil)
	if err != nil {return false, err}
	// sign
	err = tx.AddSignature(ctx, w.Account.KeyPair)
	if err != nil {return false, err}
	// use RPC to send the tx
	response := w.TxBuilder.Client.SendRawTransaction(ctx.RawTransactionString())
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {return false, fmt.Errorf(msg)}
	return response.Result, nil
}

func (w *WalletHelper) TransferNep5(assetId helper.UInt160, from string, to string, amount float64) (bool, error) {
	f, err := helper.AddressToScriptHash(from)
	if err != nil {return false, err}
	t, err := helper.AddressToScriptHash(to)
	if err != nil {return false, err}
	a := helper.Fixed8FromFloat64(amount)
	sb := sc.NewScriptBuilder()
	cp1 := sc.ContractParameter{
		Type:  sc.Hash160,
		Value: f.Bytes(),
	}
	cp2 := sc.ContractParameter{
		Type:  sc.Hash160,
		Value: t.Bytes(),
	}
	cp3 := sc.ContractParameter{
		Type:  sc.Integer,
		Value: a.Value,
	}
	sb.MakeInvocationScript(assetId.Bytes(), "transfer", []sc.ContractParameter{cp1, cp2, cp3})
	script := sb.ToArray()
	itx, err := w.TxBuilder.MakeInvocationTransaction(script, f, nil, helper.UInt160{}, helper.Zero)
	if err != nil {return false, err}
	// sign
	err = tx.AddSignature(itx, w.Account.KeyPair)
	if err != nil {return false, err}
	// use RPC to send the tx
	response := w.TxBuilder.Client.SendRawTransaction(itx.RawTransactionString())
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {return false, fmt.Errorf(msg)}
	return response.Result, nil
}