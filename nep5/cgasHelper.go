package nep5

import (
	"fmt"
	"sort"
	"time"

	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/sc"
	"github.com/joeqian10/neo-gogogo/tx"
	"github.com/joeqian10/neo-gogogo/wallet"
	"github.com/joeqian10/neo-gogogo/wallet/keys"
)

type CgasHelper Nep5Helper

func (c *CgasHelper) MintTokens(from *wallet.Account, amount float64) (string, error) {
	// A mintTokens method for CGAS users, who can transfer GAS to CGAS contract address by constructing InvocationTransaction and convert GAS to CGAS by invoking mintTokens method.
	// Upon successful invocation, CGAS in the equal value of the GAS will be added to the user's asset account.
	a := helper.Fixed8FromFloat64(amount)
	f, err := helper.AddressToScriptHash(from.Address)
	if err != nil {
		return "", err
	}

	// First, to create an InvocationTransaction, need to build the invocation script
	sb := sc.NewScriptBuilder()
	sb.MakeInvocationScript(c.scriptHash.Bytes(), "mintTokens", nil)
	script := sb.ToArray()
	//script := helper.HexToBytes("000a6d696e74546f6b656e7367f3c7a1170d2e9cb33827210daedf257db0c38c2a") // 00  0a6d696e74546f6b656e7367 f3c7a1170d2e9cb33827210daedf257db0c38c2a

	// Second, instantiate an object of InvocationTransaction
	tb := tx.NewTransactionBuilder(c.EndPoint)
	gas, err := tb.GetGasConsumed(script)
	if err != nil {
		return "", err
	}

	// Third, add Transaction inputs to the InvocationTransaction object
	inputs, totalPay, err := tb.GetTransactionInputs(f, tx.GasToken, a.Add(*gas))
	if err != nil {
		return "", err
	}
	if !totalPay.GreaterThan(a) {
		return "", fmt.Errorf("insufficient funds")
	}
	myTx := tx.NewInvocationTransaction(script)
	myTx.Gas = *gas // here gas = 0
	myTx.Inputs = append(myTx.Inputs, inputs...)

	// Fourth, add TransactionOutput to the InvocationTransaction object
	myTx.Outputs = append(myTx.Outputs, tx.NewTransactionOutput(tx.GasToken, a, c.scriptHash))              // send to CGAS contract
	myTx.Outputs = append(myTx.Outputs, tx.NewTransactionOutput(tx.GasToken, totalPay.Sub(a).Sub(*gas), f)) // send to sender account

	// Fifth, sign
	err = tx.AddSignature(myTx, from.KeyPair)
	if err != nil {
		return "", err
	}

	// use RPC to send the tx
	response := tb.Client.SendRawTransaction(myTx.RawTransactionString())
	if response.HasError() {
		return "", fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	return myTx.HashString(), nil
}

// this method calls two sub methods inside, since refund from CGAS to gas needs two steps (transactions)
func (c *CgasHelper) Refund(from *wallet.Account, txHash helper.UInt256, amount float64) (string, error) {
	txId, err := c.Refund1(from, txHash, amount)
	if err != nil {
		return "", err
	}
	time.Sleep(time.Duration(15) * time.Second) // sleep for 15 seconds to confirm the tx
	txHash1, err := helper.UInt256FromString(txId)
	if err != nil {
		return "", err
	}
	return c.Refund2(from, txHash1, amount)
}

// make this method public so developers can call this separately
func (c *CgasHelper) Refund1(from *wallet.Account, txHash helper.UInt256, amount float64) (string, error) {
	// --------------------------------------------------------
	// STEP 1
	// --------------------------------------------------------
	// build inputs
	input := tx.CoinReference{
		PrevHash:  txHash, // tx hash must be the transaction that you want to refund from
		PrevIndex: 0,
	}

	// build outputs
	output0 := tx.TransactionOutput{
		AssetId:    tx.GasToken,                      // must be GAS
		Value:      helper.Fixed8FromFloat64(amount), // if large than the amount you mint, this will fail
		ScriptHash: c.scriptHash,                     // must be the CGAS contract script hash
	}

	// build script
	user, err := helper.AddressToScriptHash(from.Address)
	if err != nil {
		return "", err
	}
	param := sc.ContractParameter{
		Type:  sc.Hash160,
		Value: user.Bytes(),
	}
	sb := sc.NewScriptBuilder()
	sb.MakeInvocationScript(c.scriptHash.Bytes(), "refund", []sc.ContractParameter{param})
	_ = sb.Emit(sc.THROWIFNOT)
	applicationScript := sb.ToArray()

	// build attributes
	attr := tx.TransactionAttribute{
		Usage: tx.Script,
		Data:  user.Bytes(), // add the user's script hash
	}

	// build transaction
	t := tx.NewInvocationTransaction(applicationScript)
	t.Inputs = append(t.Inputs, &input)
	t.Outputs = append(t.Outputs, &output0)
	t.Attributes = append(t.Attributes, &attr)

	// add two witnesses to tx
	// add the user's signature
	additionalSignature, err := from.KeyPair.Sign(t.UnsignedRawTransaction())
	if err != nil {
		return "", err
	}
	sb2 := sc.NewScriptBuilder()
	_ = sb2.EmitPushBytes(additionalSignature)
	additionalVerificationScript := sb2.ToArray()
	additionalWitness, err := tx.CreateWitness(additionalVerificationScript, keys.CreateSignatureRedeemScript(from.KeyPair.PublicKey))
	if err != nil {
		return "", err
	}

	// add CGAS script hash witness
	sb3 := sc.NewScriptBuilder()
	_ = sb3.EmitPushInt(2)
	_ = sb3.EmitPushString("1") // just build a random script matching the signature of CGAS Main method
	verificationScript := sb3.ToArray()
	witness := tx.Witness{
		InvocationScript:   verificationScript,
		VerificationScript: []byte{}, // no need to add, or the tx will become too big
	}
	ws := tx.WitnessSlice{&witness, additionalWitness}
	sort.Sort(ws)
	t.Witnesses = ws

	// use RPC to send the tx
	response := c.Client.SendRawTransaction(t.RawTransactionString())
	if response.HasError() {
		return "", fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	return t.HashString(), nil
}

// make this method public so developers can call this separately
func (c *CgasHelper) Refund2(from *wallet.Account, txHash helper.UInt256, amount float64) (string, error) {
	// --------------------------------------------------------
	// STEP 2
	// --------------------------------------------------------
	// build inputs
	input := tx.CoinReference{
		PrevHash:  txHash,
		PrevIndex: 0,
	}

	// build outputs
	user, _ := helper.AddressToScriptHash(from.Address)
	output0 := tx.TransactionOutput{
		AssetId:    tx.GasToken,                      // must be GAS
		Value:      helper.Fixed8FromFloat64(amount), // if large than the amount you mint, this will fail
		ScriptHash: user,                             // must be the CGAS contract script hash
	}

	sb := sc.NewScriptBuilder()
	_ = sb.EmitPushInt(2)
	_ = sb.EmitPushString("1") // just build a random script matching the signature of CGAS Main method
	verificationScript := sb.ToArray()
	witness := tx.Witness{
		InvocationScript:   verificationScript,
		VerificationScript: []byte{},
	}

	t := tx.NewContractTransaction()
	t.Inputs = []*tx.CoinReference{&input}
	t.Outputs = []*tx.TransactionOutput{&output0}
	t.Witnesses = []*tx.Witness{&witness}

	// use RPC to send the tx
	response := c.Client.SendRawTransaction(t.RawTransactionString())
	if response.HasError() {
		return "", fmt.Errorf(response.ErrorResponse.Error.Message)
	}

	return t.HashString(), nil
}
