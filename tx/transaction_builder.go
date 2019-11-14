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

var LocalEndPoint = "http://localhost:50003" // change to yours when using this SDK
var TestNetEndPoint = "http://seed1.ngd.network:20332"
var client *rpc.RpcClient = rpc.NewClient(LocalEndPoint)

func MakeContractTransaction(fromString string, toString string, assetIdString string, amount helper.Fixed8,
	attributes []*TransactionAttribute, changeAddressString string, fee helper.Fixed8) (*ContractTransaction, error) {
	if len(assetIdString) != 64 {
		return nil, fmt.Errorf("only global asset is allowed")
	}
	from, err := helper.UInt160FromString(fromString)
	if err != nil {
		return nil, err
	}
	to, err := helper.UInt160FromString(toString)
	if err != nil {
		return nil, err
	}
	var changeAddress helper.UInt160
	if len(changeAddressString)!=0 {
		changeAddress, err = helper.UInt160FromString(changeAddressString)
		if err != nil {
			return nil, err
		}
	} else {
		changeAddress = from
	}
	assetId, err := helper.UInt256FromString(assetIdString)
	GasToken, _ := helper.UInt256FromString(GasTokenId)
	if err != nil {
		return nil, err
	}
	ctx := NewContractTransaction()
	var inputs, gasInputs []*CoinReference
	var outputs []*TransactionOutput
	var totalPay, totalPayGas helper.Fixed8
	if attributes != nil {ctx.Attributes = attributes}

	output := NewTransactionOutput(assetId, amount, to)
	outputs = append(outputs, output)

	// no system fee for contract transaction
	if fee.GreaterThan(helper.Zero) {
		// has network fee
		if assetIdString == GasTokenId { // all are gas
			amount = amount.Add(fee)
			inputs, totalPayGas, err = GetTransactionInputs(fromString, GasTokenId, amount)
			if err != nil {return nil, err}
			if totalPayGas.GreaterThan(amount) {
				outputs = append(outputs, NewTransactionOutput(assetId, totalPayGas.Sub(amount), changeAddress))
			}
		} else { // more than gas
			inputs, totalPay, err = GetTransactionInputs(fromString, assetIdString, amount)
			if err != nil {return nil, err}
			if totalPay.GreaterThan(amount) {
				outputs = append(outputs, NewTransactionOutput(assetId, totalPay.Sub(amount), changeAddress))
			}
			gasInputs, totalPayGas, err = GetTransactionInputs(fromString, GasTokenId, fee)
			if err != nil {return nil, err}
			for _, gasInput := range gasInputs {inputs = append(inputs, gasInput)}
			if totalPayGas.GreaterThan(fee) {
				outputs = append(outputs, NewTransactionOutput(GasToken, totalPayGas.Sub(fee), changeAddress))
			}
		}
	} else {
		// no network fee
		inputs, totalPay, err = GetTransactionInputs(fromString, assetIdString, amount)
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
func GetTransactionInputs(fromString string, assetIdString string, amount helper.Fixed8) ([]*CoinReference, helper.Fixed8, error) {
	response := client.GetUnspents(fromString)
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return nil, helper.Zero, fmt.Errorf(msg)
	}
	balances := response.Result.Balance
	var unspentBalance *models.UnspentBalance = nil
	// check if there is enough balance of this asset in this account
	for _, balance := range balances {
		if balance.AssetHash == assetIdString {
			if balance.Amount >= helper.Fixed8ToFloat64(amount) {
				unspentBalance = &balance
			} else {
				return nil, helper.Zero, fmt.Errorf("not enough balance in address: %s", fromString)
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
	for unspents[i].Value <= a {
		a -= unspents[i].Value
		inputs = append(inputs, unspents[i].ToCoinReference())
		sum = sum.Add(helper.Fixed8FromFloat64(unspents[i].Value))
		i++
	}
	if a == 0 {
		return inputs, sum, nil
	}
	// use the nearest amount
	for unspents[i].Value >= a {i++}
	inputs = append(inputs, unspents[i-1].ToCoinReference())
	sum = sum.Add(helper.Fixed8FromFloat64(unspents[i-1].Value))
	return inputs, sum, nil
}


// this is a general api for invoking smart contract and creating an invocation transaction, including transferring nep-5 assets
func MakeInvocationTransaction(scriptHash []byte, operation string, args []sc.ContractParameter, fromString string,
	attributes []*TransactionAttribute, changeAddressString string, fee helper.Fixed8) (*InvocationTransaction, error) {
	from, err := helper.UInt160FromString(fromString)
	if err != nil {
		return nil, err
	}
	var changeAddress helper.UInt160
	if len(changeAddressString)!=0 {
		changeAddress, err = helper.UInt160FromString(changeAddressString)
		if err != nil {
			return nil, err
		}
	} else {
		changeAddress = from
	}
	GasToken, _ := helper.UInt256FromString(GasTokenId)
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
	inputs, totalPayGas, err := GetTransactionInputs(fromString, GasTokenId, fee)
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

// TODO
//func MakeClaimTransaction() *ClaimTransaction {
//
//}