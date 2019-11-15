package tx

import (
	"fmt"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/joeqian10/neo-gogogo/rpc/models"
	"github.com/joeqian10/neo-gogogo/sc"
	"sort"
)

const NeoTokenId = "c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b"
const GasTokenId = "602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7"

var NeoToken, _ = helper.UInt256FromString(NeoTokenId)
var GasToken, _ = helper.UInt256FromString(GasTokenId)
var LocalEndPoint = "http://localhost:50003" // change to yours when using this SDK
var TestNetEndPoint = "http://seed1.ngd.network:20332"
var client *rpc.RpcClient = rpc.NewClient(LocalEndPoint)

func MakeContractTransaction(from helper.UInt160, to helper.UInt160, assetId helper.UInt256, amount helper.Fixed8,
	attributes []*TransactionAttribute, changeAddress helper.UInt160, fee helper.Fixed8) (*ContractTransaction, error) {
	if len(changeAddress.Bytes()) == 0 {
		changeAddress = from
	}
	assetIdString := assetId.String()

	ctx := NewContractTransaction()
	var inputs, gasInputs []*CoinReference
	var outputs []*TransactionOutput
	var totalPay, totalPayGas helper.Fixed8
	var err error
	if attributes != nil {ctx.Attributes = attributes}

	output := NewTransactionOutput(assetId, amount, to)
	outputs = append(outputs, output)

	// no system fee for contract transaction
	if fee.GreaterThan(helper.Zero) {
		// has network fee
		if assetIdString == GasTokenId { // all are gas
			amount = amount.Add(fee)
			inputs, totalPayGas, err = GetTransactionInputs(from, GasToken, amount)
			if err != nil {return nil, err}
			if totalPayGas.GreaterThan(amount) {
				outputs = append(outputs, NewTransactionOutput(assetId, totalPayGas.Sub(amount), changeAddress))
			}
		} else { // more than gas
			inputs, totalPay, err = GetTransactionInputs(from, assetId, amount)
			if err != nil {return nil, err}
			if totalPay.GreaterThan(amount) {
				outputs = append(outputs, NewTransactionOutput(assetId, totalPay.Sub(amount), changeAddress))
			}
			gasInputs, totalPayGas, err = GetTransactionInputs(from, GasToken, fee)
			if err != nil {return nil, err}
			for _, gasInput := range gasInputs {inputs = append(inputs, gasInput)}
			if totalPayGas.GreaterThan(fee) {
				outputs = append(outputs, NewTransactionOutput(GasToken, totalPayGas.Sub(fee), changeAddress))
			}
		}
	} else {
		// no network fee
		inputs, totalPay, err = GetTransactionInputs(from, assetId, amount)
		if err != nil {return nil, err}
		if totalPay.GreaterThan(amount) {
			outputs = append(outputs, NewTransactionOutput(assetId, totalPay.Sub(amount), changeAddress))
		}
	}
	ctx.Inputs = inputs
	ctx.Outputs = outputs
	return ctx, nil // return unsigned contract transaction
}

// get transaction inputs according to the amount, and return UTXOs and their total amount
func GetTransactionInputs(from helper.UInt160, assetId helper.UInt256, amount helper.Fixed8) ([]*CoinReference, helper.Fixed8, error) {
	response := client.GetUnspents(helper.ScriptHashToAddress(from))
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return nil, helper.Zero, fmt.Errorf(msg)
	}
	balances := response.Result.Balance
	var unspentBalance *models.UnspentBalance = nil
	// check if there is enough balance of this asset in this account
	for _, balance := range balances {
		if balance.AssetHash == assetId.String() {
			if balance.Amount >= helper.Fixed8ToFloat64(amount) {
				unspentBalance = &balance
			} else {
				return nil, helper.Zero, fmt.Errorf("not enough balance in address: %s", helper.ScriptHashToAddress(from))
			}
		}
	}
	if unspentBalance == nil {
		return nil, helper.Zero, fmt.Errorf("asset you want to transfer is not found")
	}
	unspents := unspentBalance.Unspent
	sort.Sort(sort.Reverse(models.UnspentSlice(unspents))) // sort in decreasing order
	var i int = 0
	var a float64 = helper.Fixed8ToFloat64(amount)
	var inputs []*CoinReference = []*CoinReference{}
	var sum helper.Fixed8 = helper.Zero
	for unspents[i].Value <= a && i < len(unspents) {
		a -= unspents[i].Value
		inputs = append(inputs, ToCoinReference(unspents[i]))
		sum = sum.Add(helper.Fixed8FromFloat64(unspents[i].Value))
		i++
	}
	if a == 0 {
		return inputs, sum, nil
	}
	// use the nearest amount
	for unspents[i].Value >= a && i < len(unspents) {i++}
	inputs = append(inputs, ToCoinReference(unspents[i-1]))
	sum = sum.Add(helper.Fixed8FromFloat64(unspents[i-1].Value))
	return inputs, sum, nil
}


// this is a general api for invoking smart contract and creating an invocation transaction, including transferring nep-5 assets
func MakeInvocationTransaction(scriptHash []byte, operation string, args []sc.ContractParameter, from helper.UInt160,
	attributes []*TransactionAttribute, changeAddress helper.UInt160, fee helper.Fixed8) (*InvocationTransaction, error) {
	if len(changeAddress.Bytes())==0 {
		changeAddress = from
	}
	// make script
	sb := sc.NewScriptBuilder()
	sb.MakeInvocationScript(scriptHash, operation, args)
	script := sb.ToArray()
	// use rpc to get gas consumed
	gas, err := GetGasConsumed(script)
	if err != nil {
		return nil, err
	}
	fee = fee.Add(*gas)
	itx := NewInvocationTransaction(script)
	itx.Gas = *gas
	if itx.Size() > 1024 {
		fee = fee.Add(helper.Fixed8FromFloat64(0.001))
		fee = fee.Add(helper.Fixed8FromFloat64( float64(itx.Size()) * 0.00001))
	}
	// get transaction inputs
	inputs, totalPayGas, err := GetTransactionInputs(from, GasToken, fee)
	if err != nil {
		return nil, err
	}
	if totalPayGas.GreaterThan(fee) {
		itx.Outputs = append(itx.Outputs, NewTransactionOutput(GasToken, totalPayGas.Sub(fee), changeAddress))
	}
	itx.Inputs = inputs
	return itx, nil
}

func GetGasConsumed(script []byte) (*helper.Fixed8, error) {
	response := client.InvokeScript(helper.BytesToHex(script))
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return nil, fmt.Errorf(msg)
	}
	if response.Result.State == "FAULT" {
		return nil, fmt.Errorf("engine faulted")
	}
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
func MakeClaimTransaction(from helper.UInt160, changeAddress helper.UInt160, attributes []*TransactionAttribute) (*ClaimTransaction, error) {
	// use rpc to get claimable gas from the address
	claims, total, err := GetClaimables(from)
	if err != nil {
		return nil, err
	}
	if claims == nil || len(claims) == 0 {
		return nil, fmt.Errorf("no claim in this address")
	}
	if len(changeAddress.Bytes())==0 {
		changeAddress = from
	}
	ctx := NewClaimTransaction(claims)
	ctx.Claims = claims
	if attributes != nil {ctx.Attributes = attributes}
	var outputs []*TransactionOutput
	gasToken, _ := helper.UInt256FromString(GasTokenId)
	output := NewTransactionOutput(gasToken, *total, changeAddress)
	outputs = append(outputs, output)
	ctx.Outputs = outputs
	return ctx, nil
}

func GetClaimables(from helper.UInt160) ([]*CoinReference, *helper.Fixed8, error) {
	response := client.GetClaimable(helper.ScriptHashToAddress(from))
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return nil, nil, fmt.Errorf(msg)
	}
	var claims []*CoinReference
	claimables := response.Result.Claimables
	var MAX_CLAIMS_AMOUNT = 50 // take no more than 50 claimables
	var total helper.Fixed8
	l:= len(claimables)
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