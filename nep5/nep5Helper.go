package nep5

import (
	"encoding/binary"
	"fmt"
	"strconv"

	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/joeqian10/neo-gogogo/sc"
)

// nep5 wrapper class, api reference: https://github.com/neo-project/proposals/blob/master/nep-5.mediawiki#name
type Nep5Helper struct {
	scriptHash helper.UInt160 // the script hash of the nep5 token
	EndPoint   string
	Client     rpc.IRpcClient
}

func NewNep5Helper(scriptHash helper.UInt160, endPoint string) *Nep5Helper {
	client := rpc.NewClient(endPoint)
	if client == nil {
		return nil
	}
	return &Nep5Helper{
		scriptHash: scriptHash,
		EndPoint:   endPoint,
		Client:     client,
	}
}

func (n *Nep5Helper) TotalSupply() (uint64, error) {
	sb := sc.NewScriptBuilder()
	sb.MakeInvocationScript(n.scriptHash.Bytes(), "totalSupply", []sc.ContractParameter{})
	script := sb.ToArray()
	response := n.Client.InvokeScript(helper.BytesToHex(script), helper.ZeroScriptHashString)
	if response.HasError() {
		return 0, fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	if response.Result.State == "FAULT" {
		return 0, fmt.Errorf("engine faulted")
	}
	if len(response.Result.Stack) == 0 {
		return 0, fmt.Errorf("no stack result returned")
	}
	stack := response.Result.Stack[0]
	bytes := helper.HexToBytes(stack.Value.(string))
	for len(bytes) < 8 {
		bytes = append(bytes, byte(0x00))
	}
	ts := binary.LittleEndian.Uint64(bytes)
	return ts, nil
}

func (n *Nep5Helper) Name() (string, error) {
	sb := sc.NewScriptBuilder()
	sb.MakeInvocationScript(n.scriptHash.Bytes(), "name", []sc.ContractParameter{})
	script := sb.ToArray()
	response := n.Client.InvokeScript(helper.BytesToHex(script), helper.ZeroScriptHashString)
	if response.HasError() {
		return "", fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	if response.Result.State == "FAULT" {
		return "", fmt.Errorf("engine faulted")
	}
	if len(response.Result.Stack) == 0 {
		return "", fmt.Errorf("no stack result returned")
	}
	stack := response.Result.Stack[0]
	name := string(helper.HexToBytes(stack.Value.(string)))
	return name, nil
}

func (n *Nep5Helper) Symbol() (string, error) {
	sb := sc.NewScriptBuilder()
	sb.MakeInvocationScript(n.scriptHash.Bytes(), "symbol", []sc.ContractParameter{})
	script := sb.ToArray()
	response := n.Client.InvokeScript(helper.BytesToHex(script), helper.ZeroScriptHashString)
	if response.HasError() {
		return "", fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	if response.Result.State == "FAULT" {
		return "", fmt.Errorf("engine faulted")
	}
	if len(response.Result.Stack) == 0 {
		return "", fmt.Errorf("no stack result returned")
	}
	stack := response.Result.Stack[0]
	symbol := string(helper.HexToBytes(stack.Value.(string)))
	return symbol, nil
}

func (n *Nep5Helper) Decimals() (uint8, error) {
	sb := sc.NewScriptBuilder()
	sb.MakeInvocationScript(n.scriptHash.Bytes(), "decimals", []sc.ContractParameter{})
	script := sb.ToArray()
	response := n.Client.InvokeScript(helper.BytesToHex(script), helper.ZeroScriptHashString)
	if response.HasError() {
		return 0, fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	if response.Result.State == "FAULT" {
		return 0, fmt.Errorf("engine faulted")
	}
	if len(response.Result.Stack) == 0 {
		return 0, fmt.Errorf("no stack result returned")
	}
	stack := response.Result.Stack[0]
	decimals, err := strconv.ParseUint(stack.Value.(string), 10, 8)
	if err != nil {
		return 0, fmt.Errorf("conversion failed")
	}
	return uint8(decimals), nil
}

func (n *Nep5Helper) BalanceOf(address helper.UInt160) (uint64, error) {
	sb := sc.NewScriptBuilder()
	cp := sc.ContractParameter{
		Type:  sc.Hash160,
		Value: address.Bytes(),
	}
	sb.MakeInvocationScript(n.scriptHash.Bytes(), "balanceOf", []sc.ContractParameter{cp})
	script := sb.ToArray()
	response := n.Client.InvokeScript(helper.BytesToHex(script), helper.ZeroScriptHashString)
	if response.HasError() {
		return 0, fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	if response.Result.State == "FAULT" {
		return 0, fmt.Errorf("engine faulted")
	}
	if len(response.Result.Stack) == 0 {
		return 0, fmt.Errorf("no stack result returned")
	}
	stack := response.Result.Stack[0]
	bytes := helper.HexToBytes(stack.Value.(string))
	for len(bytes) < 8 {
		bytes = append(bytes, byte(0x00))
	}
	balance := binary.LittleEndian.Uint64(bytes)
	return balance, nil
}

// This method is deprecated
// Transfer is only testing the transfer script, please use WalletHelper to truly transfer nep5 token
//func (n *Nep5Helper) Transfer(from helper.UInt160, to helper.UInt160, amount helper.Fixed8) (bool, []byte, error) {
//	sb := sc.NewScriptBuilder()
//	cp1 := sc.ContractParameter{
//		Type:  sc.Hash160,
//		Value: from.Bytes(),
//	}
//	cp2 := sc.ContractParameter{
//		Type:  sc.Hash160,
//		Value: to.Bytes(),
//	}
//	cp3 := sc.ContractParameter{
//		Type:  sc.Integer,
//		Value: amount.Value,
//	}
//	sb.MakeInvocationScript(n.scriptHash.Bytes(), "transfer", []sc.ContractParameter{cp1, cp2, cp3})
//	script := sb.ToArray()
//	response := n.Client.InvokeScript(helper.BytesToHex(script))
//	if response.HasError() {
//		return false, []byte{}, fmt.Errorf(response.ErrorResponse.Error.Message)
//	}
//	if response.Result.State == "FAULT" {
//		return false, []byte{}, fmt.Errorf("engine faulted")
//	}
//	if len(response.Result.Stack) == 0 {
//		return false, []byte{}, fmt.Errorf("no stack result returned")
//	}
//	stack := response.Result.Stack[0]
//	b, err := strconv.ParseBool(stack.Value)
//	if err != nil {
//		return false, []byte{}, fmt.Errorf("conversion failed")
//	}
//	return b, script, nil
//}
