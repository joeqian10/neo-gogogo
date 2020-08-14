package tx

import (
	"fmt"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/joeqian10/neo-gogogo/rpc/models"
	"github.com/joeqian10/neo-gogogo/sc"
	"math/big"
	"sort"
)

const NeoTokenId = "c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b"
const GasTokenId = "602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7"

var NeoToken, _ = helper.UInt256FromString(NeoTokenId)
var GasToken, _ = helper.UInt256FromString(GasTokenId)

type TransactionBuilder struct {
	EndPoint string
	Client   rpc.IRpcClient
}

func NewTransactionBuilder(endPoint string) *TransactionBuilder {
	client := rpc.NewClient(endPoint)
	if client == nil {
		return nil
	}
	return &TransactionBuilder{
		EndPoint: endPoint,
		Client:   client,
	}
}

func (tb *TransactionBuilder) MakeContractTransaction(from helper.UInt160, to helper.UInt160, assetId helper.UInt256, amount helper.Fixed8,
	attributes []*TransactionAttribute, changeAddress helper.UInt160, fee helper.Fixed8) (*ContractTransaction, error) {
	if changeAddress.String() == "0000000000000000000000000000000000000000" {
		changeAddress = from
	}
	assetIdString := assetId.String()

	ctx := NewContractTransaction()
	var inputs, gasInputs []*CoinReference
	var outputs []*TransactionOutput
	var totalPay, totalPayGas helper.Fixed8
	var err error
	if attributes != nil {
		ctx.Attributes = attributes
	}

	output := NewTransactionOutput(assetId, amount, to)
	outputs = append(outputs, output)

	// no system fee for contract transaction
	if fee.GreaterThan(helper.Zero) {
		// has network fee
		if assetIdString == GasTokenId { // all are gas
			amount = amount.Add(fee)
			inputs, totalPayGas, err = tb.GetTransactionInputs(from, GasToken, amount)
			if err != nil {
				return nil, err
			}
			if totalPayGas.GreaterThan(amount) {
				outputs = append(outputs, NewTransactionOutput(assetId, totalPayGas.Sub(amount), changeAddress))
			}
		} else { // more than gas
			inputs, totalPay, err = tb.GetTransactionInputs(from, assetId, amount)
			if err != nil {
				return nil, err
			}
			if totalPay.GreaterThan(amount) {
				outputs = append(outputs, NewTransactionOutput(assetId, totalPay.Sub(amount), changeAddress))
			}
			gasInputs, totalPayGas, err = tb.GetTransactionInputs(from, GasToken, fee)
			if err != nil {
				return nil, err
			}
			for _, gasInput := range gasInputs {
				inputs = append(inputs, gasInput)
			}
			if totalPayGas.GreaterThan(fee) {
				outputs = append(outputs, NewTransactionOutput(GasToken, totalPayGas.Sub(fee), changeAddress))
			}
		}
	} else {
		// no network fee
		inputs, totalPay, err = tb.GetTransactionInputs(from, assetId, amount)
		if err != nil {
			return nil, err
		}
		if totalPay.GreaterThan(amount) {
			outputs = append(outputs, NewTransactionOutput(assetId, totalPay.Sub(amount), changeAddress))
		}
	}
	ctx.Inputs = inputs
	ctx.Outputs = outputs
	return ctx, nil // return unsigned contract transaction
}

// get transaction inputs according to the amount, and return UTXOs and their total amount
func (tb *TransactionBuilder) GetTransactionInputs(from helper.UInt160, assetId helper.UInt256, amount helper.Fixed8) ([]*CoinReference, helper.Fixed8, error) {
	if amount.Equal(helper.Zero) {
		return nil, helper.Zero, nil
	}
	unspentBalance, available, err := tb.GetBalance(from, assetId)
	if err != nil {
		return nil, helper.Zero, err
	}
	if available.LessThan(amount) {
		return nil, helper.Zero, fmt.Errorf("not enough balance in address: %s", helper.ScriptHashToAddress(from))
	}
	unspents := unspentBalance.Unspents
	sort.Sort(sort.Reverse(models.UnspentSlice(unspents))) // sort in decreasing order
	var i int = 0
	var a float64 = helper.Fixed8ToFloat64(amount)
	var inputs []*CoinReference = []*CoinReference{}
	var sum helper.Fixed8 = helper.Zero
	for i < len(unspents) && unspents[i].Value <= a {
		a -= unspents[i].Value
		inputs = append(inputs, ToCoinReference(unspents[i]))
		sum = sum.Add(helper.Fixed8FromFloat64(unspents[i].Value))
		i++
	}
	if a == 0 {
		return inputs, sum, nil
	}
	// use the nearest amount
	for i < len(unspents) && unspents[i].Value >= a {
		i++
	}
	inputs = append(inputs, ToCoinReference(unspents[i-1]))
	sum = sum.Add(helper.Fixed8FromFloat64(unspents[i-1].Value))
	return inputs, sum, nil
}

// GetBalance is used to get balance of neo or gas or other utxo asset
func (tb *TransactionBuilder) GetBalance(account helper.UInt160, assetId helper.UInt256) (*models.UnspentBalance, helper.Fixed8, error) {
	response := tb.Client.GetUnspents(helper.ScriptHashToAddress(account))
	if response.HasError() {
		return nil, helper.Zero, fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	balances := response.Result.Balances
	// check if there is enough balance of this asset in this account
	for _, balance := range balances {
		if balance.AssetHash == assetId.String() {
			return &balance, helper.Fixed8FromFloat64(balance.Amount), nil
		}
	}
	return nil, helper.Zero, fmt.Errorf("asset not found")
}

// this is a general api for invoking smart contract and creating an invocation transaction, including transferring nep-5 assets
func (tb *TransactionBuilder) MakeInvocationTransaction(script []byte, from helper.UInt160, attributes []*TransactionAttribute, changeAddress helper.UInt160, fee helper.Fixed8) (*InvocationTransaction, error) {
	if changeAddress.String() == "0000000000000000000000000000000000000000" {
		changeAddress = from
	}
	// use rpc to get gas consumed
	gas, err := tb.GetGasConsumed(script)
	if err != nil {
		return nil, err
	}
	newGas := *gas
	//newGas = newGas.Add(helper.Fixed8FromInt64(1))
	fee = fee.Add(newGas)
	//fee = fee.Add(helper.Fixed8FromInt64(1))
	itx := NewInvocationTransaction(script)
	if attributes != nil {
		itx.Attributes = attributes
	}
	itx.Gas = newGas
	if itx.Size() > 1024 {
		fee = fee.Add(helper.Fixed8FromFloat64(0.001))
		fee = fee.Add(helper.Fixed8FromFloat64(float64(itx.Size()) * 0.00001))
	}
	// get transaction inputs
	inputs, totalPayGas, err := tb.GetTransactionInputs(from, GasToken, fee)
	if err != nil {
		return nil, err
	}
	if totalPayGas.GreaterThan(fee) {
		itx.Outputs = append(itx.Outputs, NewTransactionOutput(GasToken, totalPayGas.Sub(fee), changeAddress))
	}
	itx.Inputs = inputs
	return itx, nil
}

func (tb *TransactionBuilder) GetGasConsumed(script []byte) (*helper.Fixed8, error) {
	response := tb.Client.InvokeScript(helper.BytesToHex(script))
	if response.HasError() {
		return nil, fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	// transfer script will return "FAULT" when checking witness, so comment error for this issue https://github.com/neo-project/neo/pull/335
	//if response.Result.State == "FAULT" {
	//	return nil, fmt.Errorf("engine faulted")
	//}
	gasConsumed, err := helper.Fixed8FromString(response.Result.GasConsumed)
	if err != nil {
		return nil, err
	}
	gas := gasConsumed.Sub(helper.Fixed8FromInt64(10))
	if gas.LessThan(helper.Zero) || gas.Equal(helper.Zero) {
		return &helper.Zero, nil
	} else {
		g := gas.Ceiling()
		return &g, nil
	}
}

//
func (tb *TransactionBuilder) MakeClaimTransaction(from helper.UInt160, changeAddress helper.UInt160, attributes []*TransactionAttribute) (*ClaimTransaction, error) {
	// use rpc to get claimable gas from the address
	claims, total, err := tb.GetClaimables(from)
	if err != nil {
		return nil, err
	}
	if claims == nil || len(claims) == 0 {
		return nil, fmt.Errorf("no claim in this address")
	}
	if changeAddress.String() == "0000000000000000000000000000000000000000" {
		changeAddress = from
	}
	ctx := NewClaimTransaction(claims)
	ctx.Claims = claims
	if attributes != nil {
		ctx.Attributes = attributes
	}
	var outputs []*TransactionOutput
	gasToken, _ := helper.UInt256FromString(GasTokenId)
	output := NewTransactionOutput(gasToken, *total, changeAddress)
	outputs = append(outputs, output)
	ctx.Outputs = outputs
	return ctx, nil
}

func (tb *TransactionBuilder) GetClaimables(from helper.UInt160) ([]*CoinReference, *helper.Fixed8, error) {
	response := tb.Client.GetClaimable(helper.ScriptHashToAddress(from))
	if response.HasError() {
		return nil, nil, fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	var claims []*CoinReference
	claimables := response.Result.Claimables
	var MAX_CLAIMS_AMOUNT = 50 // take no more than 50 claimables
	var total helper.Fixed8
	l := len(claimables)
	for i := 0; i <= l-1 && i <= MAX_CLAIMS_AMOUNT-1; i++ {
		h, err := helper.UInt256FromString(claimables[i].TxId)
		if err != nil {
			return nil, nil, err
		}
		claim := &CoinReference{
			PrevHash:  h,
			PrevIndex: uint16(claimables[i].N),
		}
		claims = append(claims, claim)
		total = total.Add(helper.Fixed8FromFloat64(claimables[i].Unclaimed))
	}
	return claims, &total, nil
}

func (tb *TransactionBuilder) LoadScriptTransaction(script []byte,
	paramTypes string, returnTypeHexString string,
	hasStorage bool, hasDynamicInvoke bool, isPayable bool,
	contractName string, contractVersion string, contractAuthor string, contractEmail string, contractDescription string) (tx *InvocationTransaction, scriptHash helper.UInt160, err error) {
	scriptHash, err = helper.BytesToScriptHash(script)
	if err != nil {
		return nil, scriptHash, err
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
	tx = NewInvocationTransaction(newScript)
	return tx, scriptHash, nil
}
