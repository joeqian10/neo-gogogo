package nep5
//
//import (
//	"github.com/joeqian10/neo-gogogo/helper"
//	"github.com/joeqian10/neo-gogogo/sc"
//	"github.com/joeqian10/neo-gogogo/tx"
//	"github.com/joeqian10/neo-gogogo/wallet"
//	"github.com/stretchr/testify/assert"
//	"log"
//	"testing"
//	"time"
//)
//
//func TestTransferCGas(t *testing.T) {
//	cgasHash, _ := helper.UInt160FromString("0x74f2dc36a68fdc4682034178eb2220729231db76")
//	cgas := NewNep5Helper(cgasHash,"http://seed1.ngd.network:20332")
//	addr1, _ := helper.AddressToScriptHash("AQJgMwLnhJWj6RqvESSXyXwDsM3hxxfSEr")
//	addr2, _ := helper.AddressToScriptHash("AaCkxSP1gkukyU5ZJPE4RiZC4mMtsjexnx")
//	bal1, err := cgas.BalanceOf(addr1)
//	if err != nil {log.Fatalln(err)}
//	bal2, err := cgas.BalanceOf(addr2)
//	if err != nil {log.Fatalln(err)}
//	log.Print(bal1,bal2)
//	ok, script, err := cgas.Transfer(addr1, addr2,helper.Fixed8FromFloat64(0.01))
//	if err != nil {log.Fatalln(err)}
//	if !ok {log.Fatalln("Failed")}
//	log.Println(script)
//}
//
//func Test_TransferCGas(t *testing.T)  {
//	// first mint some CGAS to "AQzRMe3zyGS8W177xLJfewRRQZY2kddMun"
//	cgasHash, _ := helper.UInt160FromString("0x2a8cc3b07d25dfae0d212738b39c2e0d17a1c7f3")
//	endPoint := "http://localhost:20002" // local
//	cgasHelper := NewCgasHelperFromNep5Helper(NewNep5Helper(cgasHash, endPoint))
//	fromAccount, err := wallet.NewAccountFromWIF("L3Hab7wL43SbWLnkfnVCp6dT99xzfB4qLZxeR9dFgcWWPirwKyXp")
//	assert.Nil(t, err)
//	txHash, err := cgasHelper.MintTokens(fromAccount, 1000)
//
//	log.Println(txHash)
//
//	time.Sleep(time.Duration(15) * time.Second) // wait for the
//
//	from, _ := helper.AddressToScriptHash("AQzRMe3zyGS8W177xLJfewRRQZY2kddMun")
//	to, _ := helper.AddressToScriptHash("APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR")
//	amount := helper.Fixed8FromFloat64(10)
//	// call another contract to transfer CGAS
//	scriptHash, _ := helper.UInt160FromString("0xff4d6be1716acfd9ece0b957ed1e00722a1a3745")
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
//	sb := sc.NewScriptBuilder()
//	sb.MakeInvocationScript(scriptHash.Bytes(), "transferCGAS", []sc.ContractParameter{cp1,cp2,cp3})
//	script := sb.ToArray()
//	tb := tx.NewTransactionBuilder(endPoint)
//	// create an InvocationTransaction
//	itx, err := tb.MakeInvocationTransaction(script, from, nil, from, helper.Zero)
//	assert.Nil(t, err)
//	// sign transaction
//	err = tx.AddSignature(itx, fromAccount.KeyPair)
//	assert.Nil(t, err)
//
//	rawTxString := itx.RawTransactionString()
//	log.Println("RawTransactionString : ", rawTxString)
//
//	response := tb.Client.SendRawTransaction(rawTxString)
//	if response.HasError() {
//		log.Println("SendRawTransaction error: ", response.ErrorResponse.Error.Message)
//	}
//	log.Println("txHash is: ", itx.HashString())
//}
//
//func Test_TransferCGas2(t *testing.T)  {
//	//cgasHash, _ := helper.UInt160FromString("0x74f2dc36a68fdc4682034178eb2220729231db76")
//	endPoint := "http://seed1.ngd.network:20332"
//	//cgasHelper := (*CgasHelper)(NewNep5Helper(cgasHash,endPoint))
//	fromAccount, err := wallet.NewAccountFromWIF("KwERwXSv5ctu1ne5yJxLDcAsCjiVgp3snX8pn8nBZUN3jDtFeaxH")
//	//if err != nil {log.Fatalln(err)}
//	amountF := 1.0
//	//txHash, err := cgasHelper.MintTokens(fromAccount, amountF)
//	//if err != nil {log.Fatalln(err)}
//	//log.Println(txHash)
//	//time.Sleep(time.Duration(15) * time.Second) // wait for the
//
//	from, _ := helper.AddressToScriptHash(fromAccount.Address) // AQJgMwLnhJWj6RqvESSXyXwDsM3hxxfSEr
//	to, _ := helper.AddressToScriptHash("AaCkxSP1gkukyU5ZJPE4RiZC4mMtsjexnx")
//	amount := helper.Fixed8FromFloat64(amountF)
//	// call another contract to transfer CGAS
//	scriptHash, _ := helper.UInt160FromString("0x40341b1d0123890cddf99501a3e0cfdcd90cba1a")
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
//	sb := sc.NewScriptBuilder()
//	sb.MakeInvocationScript(scriptHash.Bytes(), "transferCGAS", []sc.ContractParameter{cp1,cp2,cp3})
//	script := sb.ToArray()
//	tb := tx.NewTransactionBuilder(endPoint)
//	// create an InvocationTransaction
//	itx, err := tb.MakeInvocationTransaction(script, from, nil, from, helper.Zero)
//	if err != nil {log.Fatalln(err)}
//	// sign transaction
//	err = tx.AddSignature(itx, fromAccount.KeyPair)
//	if err != nil {log.Fatalln(err)}
//
//	rawTxString := itx.RawTransactionString()
//	log.Println("RawTransactionString : ", rawTxString)
//
//	response := tb.Client.SendRawTransaction(rawTxString)
//	if response.HasError() {
//		log.Println("SendRawTransaction error: ", response.ErrorResponse.Error.Message)
//	}
//	log.Println("txHash is: ", itx.HashString())
//}
