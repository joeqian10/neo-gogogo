package nep5

import (
	"fmt"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/joeqian10/neo-gogogo/sc"
	"strconv"
)

// nep5 wrapper class, api reference: https://github.com/neo-project/proposals/blob/master/nep-5.mediawiki#name
type Nep5Helper struct {
	Client *rpc.RpcClient
}

func NewNep5Helper(endPoint string) *Nep5Helper {
	client := rpc.NewClient(endPoint)
	if client == nil {
		return nil
	}
	return &Nep5Helper{
		Client:   client,
	}
}

func (n *Nep5Helper) TotalSupply(scriptHash helper.UInt160) (uint64, error) {
	sb := sc.NewScriptBuilder()
	sb.MakeInvocationScript(scriptHash.Bytes(), "totalSupply", []sc.ContractParameter{})
	script := sb.ToArray()
	response := n.Client.InvokeScript(helper.BytesToHex(script))
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return 0, fmt.Errorf(msg)
	}
	if response.Result.State == "FAULT" {
		return 0, fmt.Errorf("engine faulted")
	}
	if len(response.Result.Stack) == 0 {
		return 0, fmt.Errorf("no stack result returned")
	}
	stack := response.Result.Stack[0]
	ts, err := helper.ParseVarInt(helper.HexTobytes(stack.Value))
	if err != nil {
		return 0, fmt.Errorf("conversion failed")
	}
	return ts.Value, nil
}

func (n *Nep5Helper) Name(scriptHash helper.UInt160) (string, error) {
	sb := sc.NewScriptBuilder()
	sb.MakeInvocationScript(scriptHash.Bytes(), "name", []sc.ContractParameter{})
	script := sb.ToArray()
	response := n.Client.InvokeScript(helper.BytesToHex(script))
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return "", fmt.Errorf(msg)
	}
	if response.Result.State == "FAULT" {
		return "", fmt.Errorf("engine faulted")
	}
	if len(response.Result.Stack) == 0 {
		return "", fmt.Errorf("no stack result returned")
	}
	stack := response.Result.Stack[0]
	name := string(helper.HexTobytes(stack.Value))
	return name, nil
}

func (n *Nep5Helper) Symbol(scriptHash helper.UInt160) (string, error) {
	sb := sc.NewScriptBuilder()
	sb.MakeInvocationScript(scriptHash.Bytes(), "symbol", []sc.ContractParameter{})
	script := sb.ToArray()
	response := n.Client.InvokeScript(helper.BytesToHex(script))
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return "", fmt.Errorf(msg)
	}
	if response.Result.State == "FAULT" {
		return "", fmt.Errorf("engine faulted")
	}
	if len(response.Result.Stack) == 0 {
		return "", fmt.Errorf("no stack result returned")
	}
	stack := response.Result.Stack[0]
	symbol := string(helper.HexTobytes(stack.Value))
	return symbol, nil
}

func (n *Nep5Helper) Decimals(scriptHash helper.UInt160) (uint8, error) {
	sb := sc.NewScriptBuilder()
	sb.MakeInvocationScript(scriptHash.Bytes(), "decimals", []sc.ContractParameter{})
	script := sb.ToArray()
	response := n.Client.InvokeScript(helper.BytesToHex(script))
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return 0, fmt.Errorf(msg)
	}
	if response.Result.State == "FAULT" {
		return 0, fmt.Errorf("engine faulted")
	}
	if len(response.Result.Stack) == 0 {
		return 0, fmt.Errorf("no stack result returned")
	}
	stack := response.Result.Stack[0]
	decimals, err := strconv.ParseUint(stack.Value, 10, 8)
	if err != nil {
		return 0, fmt.Errorf("conversion failed")
	}
	return uint8(decimals), nil
}

func (n *Nep5Helper) BalanceOf(scriptHash helper.UInt160, address helper.UInt160) (uint64, error) {
	sb := sc.NewScriptBuilder()
	cp := sc.ContractParameter{
		Type:  sc.Hash160,
		Value: address.Bytes(),
	}
	sb.MakeInvocationScript(scriptHash.Bytes(), "totalSupply", []sc.ContractParameter{cp})
	script := sb.ToArray()
	response := n.Client.InvokeScript(helper.BytesToHex(script))
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return 0, fmt.Errorf(msg)
	}
	if response.Result.State == "FAULT" {
		return 0, fmt.Errorf("engine faulted")
	}
	if len(response.Result.Stack) == 0 {
		return 0, fmt.Errorf("no stack result returned")
	}
	stack := response.Result.Stack[0]
	balance, err := helper.ParseVarInt(helper.HexTobytes(stack.Value))
	if err != nil {
		return 0, fmt.Errorf("conversion failed")
	}
	return balance.Value, nil
}

func (n *Nep5Helper) Transfer(scriptHash helper.UInt160, from helper.UInt160, to helper.UInt160, amount uint64) ([]byte, error) {
	sb := sc.NewScriptBuilder()
	cp1 := sc.ContractParameter{
		Type:  sc.Hash160,
		Value: from.Bytes(),
	}
	cp2 := sc.ContractParameter{
		Type:  sc.Hash160,
		Value: to.Bytes(),
	}
	cp3 := sc.ContractParameter{
		Type:  sc.Integer,
		Value: amount,
	}
	sb.MakeInvocationScript(scriptHash.Bytes(), "totalSupply", []sc.ContractParameter{cp1, cp2, cp3})
	script := sb.ToArray()
	response := n.Client.InvokeScript(helper.BytesToHex(script))
	msg := response.ErrorResponse.Error.Message
	if len(msg) != 0 {
		return nil, fmt.Errorf(msg)
	}
	if response.Result.State == "FAULT" {
		return nil, fmt.Errorf("engine faulted")
	}
	if len(response.Result.Stack) == 0 {
		return nil, fmt.Errorf("no stack result returned")
	}
	stack := response.Result.Stack[0]
	b, err := strconv.ParseBool(stack.Value)
	if err != nil {
		return nil, fmt.Errorf("conversion failed")
	}
	if b == true {
		return script, nil
	}
	return nil, fmt.Errorf("false result")
}