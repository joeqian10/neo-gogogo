package wallet

import (
	"fmt"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/sc"
	"github.com/joeqian10/neo-gogogo/tx"
	"math/big"
	"strconv"
)

type WalletHelper struct {
	TxBuilder *tx.TransactionBuilder
	Account   *Account
}

func NewWalletHelper(txBuilder *tx.TransactionBuilder, account *Account) *WalletHelper {
	return &WalletHelper{
		TxBuilder: txBuilder,
		Account:   account,
	}
}

// GetBalance is used to transfer neo or gas or other utxo asset, single signature
func (w *WalletHelper) GetBalance(address string) (neoBalance int, gasBalance float64, err error) {
	response := w.TxBuilder.Client.GetAccountState(address)
	if response.HasError() {
		return 0, 0, fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	balances := response.Result.Balances
	for _, balance := range balances {
		assetId, err := helper.UInt256FromString(balance.Asset)
		if err != nil {
			return 0, 0, err
		}
		if assetId == tx.NeoToken {
			neoBalance, err = strconv.Atoi(balance.Value)
			if err != nil {
				return 0, 0, err
			}
		} else if assetId == tx.GasToken {
			gasBalance, err = strconv.ParseFloat(balance.Value, 64)
			if err != nil {
				return 0, 0, err
			}
		}
	}
	return neoBalance, gasBalance, nil
}

// Transfer is used to transfer neo or gas or other utxo asset, single signature, return txid
func (w *WalletHelper) Transfer(assetId helper.UInt256, from string, to string, amount float64) (string, error) {
	f, err := helper.AddressToScriptHash(from)
	if err != nil {
		return "", err
	}
	t, err := helper.AddressToScriptHash(to)
	if err != nil {
		return "", err
	}
	a := helper.Fixed8FromFloat64(amount)
	ctx, err := w.TxBuilder.MakeContractTransaction(f, t, assetId, a, nil, helper.UInt160{}, helper.Zero)
	if err != nil {
		return "", err
	}
	// sign
	err = tx.AddSignature(ctx, w.Account.KeyPair)
	if err != nil {
		return "", err
	}
	// use RPC to send the tx
	response := w.TxBuilder.Client.SendRawTransaction(ctx.RawTransactionString())
	if response.HasError() {
		return "", fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	return ctx.HashString(), nil
}

// ClaimGas, return txid
func (w *WalletHelper) ClaimGas(from string) (string, error) {
	f, err := helper.AddressToScriptHash(from)
	if err != nil {
		return "", err
	}
	ctx, err := w.TxBuilder.MakeClaimTransaction(f, helper.UInt160{}, nil)
	if err != nil {
		return "", err
	}
	// sign
	err = tx.AddSignature(ctx, w.Account.KeyPair)
	if err != nil {
		return "", err
	}
	// use RPC to send the tx
	response := w.TxBuilder.Client.SendRawTransaction(ctx.RawTransactionString())
	if response.HasError() {
		return "", fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	return ctx.HashString(), nil
}

func (w *WalletHelper) TransferNep5(assetId helper.UInt160, from string, to string, amount float64) (string, error) {
	f, err := helper.AddressToScriptHash(from)
	if err != nil {
		return "", err
	}
	t, err := helper.AddressToScriptHash(to)
	if err != nil {
		return "", err
	}
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
		Value: *big.NewInt(a.Value),
	}
	sb.MakeInvocationScript(assetId.Bytes(), "transfer", []sc.ContractParameter{cp1, cp2, cp3})
	script := sb.ToArray()
	itx, err := w.TxBuilder.MakeInvocationTransaction(script, f, nil, helper.UInt160{}, helper.Zero)
	if err != nil {
		return "", err
	}
	// sign
	err = tx.AddSignature(itx, w.Account.KeyPair)
	if err != nil {
		return "", err
	}
	// use RPC to send the tx
	response := w.TxBuilder.Client.SendRawTransaction(itx.RawTransactionString())
	if response.HasError() {
		return "", fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	return itx.HashString(), nil
}

func (w *WalletHelper) DeployContract(script []byte,
	paramTypes string, returnTypeHexString string,
	hasStorage bool, hasDynamicInvoke bool, isPayable bool,
	contractName string, contractVersion string, contractAuthor string, contractEmail string, contractDescription string) (*helper.UInt160, error) {
	scriptHash, err := helper.BytesToScriptHash(script)
	if err != nil {
		return nil, err
	}
	parameterList := helper.HexToBytes(paramTypes)                                    // 0710
	returnType := sc.ContractParameterType(helper.HexToBytes(returnTypeHexString)[0]) // 05
	property := sc.NoProperty
	if hasStorage {
		property |= sc.HasStorage
	}
	if hasDynamicInvoke {
		property |= sc.HasDynamicInvoke
	}
	if isPayable {
		property |= sc.Payable
	}
	p1 := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: script,
	}
	p2 := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: parameterList,
	}
	p3 := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(returnType)),
	}
	p4 := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(property)),
	}
	p5 := sc.ContractParameter{
		Type:  sc.String,
		Value: contractName,
	}
	p6 := sc.ContractParameter{
		Type:  sc.String,
		Value: contractVersion,
	}
	p7 := sc.ContractParameter{
		Type:  sc.String,
		Value: contractAuthor,
	}
	p8 := sc.ContractParameter{
		Type:  sc.String,
		Value: contractEmail,
	}
	p9 := sc.ContractParameter{
		Type:  sc.String,
		Value: contractDescription,
	}
	sb := sc.NewScriptBuilder()
	sb.EmitSysCall("Neo.Contract.Create", []sc.ContractParameter{p1, p2, p3, p4, p5, p6, p7, p8, p9})
	newScript := sb.ToArray()

	from, err := helper.AddressToScriptHash(w.Account.Address)
	if err != nil {
		return nil, err
	}
	itx, err := w.TxBuilder.MakeInvocationTransaction(newScript, from, nil, helper.UInt160{}, helper.Zero)
	if err != nil {
		return nil, err
	}
	err = tx.AddSignature(itx, w.Account.KeyPair)
	if err != nil {
		return nil, err
	}
	response := w.TxBuilder.Client.SendRawTransaction(itx.RawTransactionString())
	if response.HasError() {
		return nil, fmt.Errorf(response.ErrorResponse.Error.Message)
	}

	return &scriptHash, nil
}

func (w *WalletHelper) InvokeContract(scriptHash helper.UInt160, method string, args []sc.ContractParameter) (*helper.UInt256, error) {
	sb := sc.NewScriptBuilder()
	sb.MakeInvocationScript(scriptHash.Bytes(), method, args)
	script := sb.ToArray()

	from, err := helper.AddressToScriptHash(w.Account.Address)
	if err != nil {
		return nil, err
	}
	itx, err := w.TxBuilder.MakeInvocationTransaction(script, from, nil, helper.UInt160{}, helper.Zero)
	if err != nil {
		return nil, err
	}
	response := w.TxBuilder.Client.SendRawTransaction(itx.RawTransactionString())
	if response.HasError() {
		return nil, fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	return &itx.Hash, nil
}
