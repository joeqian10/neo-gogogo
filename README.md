# neo-gogogo

## Overview
This is a light-weight golang SDK for neo 2.x created by NGD Shanghai.

## Getting Started
This SDK has seven modules, features and usages of each module will be introduced below.

### "crypto" module
This module offers methods used for cryptography purposes, such as AES encryption/decryption, Base58 encoding/decoding, Hash160/Hash256 hashing functions. For more information about the crypto algorithms used in neo, refer to [Cryptography](https://docs.neo.org/docs/en-us/tooldev/concept/cryptography/encode_algorithm.html).

Typical usage:

```golang
package sample

import "encoding/hex"
import "github.com/joeqian10/neo-gogogo/crypto"

func SampleMethod() {
	var b58CheckEncoded = "KxhEDBQyyEFymvfJD96q8stMbJMbZUb6D1PmXqBWZDU2WvbvVs9o"
	var b58CheckDecodedHex = "802bfe58ab6d9fd575bdc3a624e4825dd2b375d64ac033fbc46ea79dbab4f69a3e01"

	b58CheckDecoded, _ := hex.DecodeString(b58CheckDecodedHex)
	encoded := crypto.Base58CheckEncode(b58CheckDecoded)
    decoded, err := crypto.Base58CheckDecode(b58CheckEncoded)
    
    ...
}
```

### "helper" module
As its name indicated, this module acts as a helper and provides some standard data types used in neo, such as `Fixed8`, `UInt160`, `UInt256`, and some auxiliary methods with basic functionalities including conversion between a hex string and a byte array, conversion between a script hash and a standard neo address, concatenating/reversing byte arrays and so on.

Typical usage:

```golang
package sample

import "encoding/hex"
import "github.com/joeqian10/neo-gogogo/helper"

func SampleMethod() {
    // Fixed8
    f1 := helper.NewFixed8(1234567800000000)
    f2 := helper.Fixed8FromInt64(12345678)
    f3 := helper.Fixed8FromFloat64(12345678.0)
    // f1, f2, f3 are all equal

    // UInt160
    hexStr := "2d3b96ae1bcc5a585e075e3b81920210dec16302"
    v1, err := helper.UInt160FromString(hexStr)
	b1, err := hex.DecodeString(hexStr)
    v2, err := helper.UInt160FromBytes(ReverseBytes(b))
    // v1 and v2 are equal

    // UInt256
    str := "f037308fa0ab18155bccfc08485468c112409ea5064595699e98c545f245f32d"
    u1, err := helper.UInt256FromString(str)
	b2, err := hex.DecodeString(hexStr)
    u2, err := helper.UInt256FromBytes(ReverseBytes(b))
    // u1 and u2 are equal
    
    // reverse bytes
    b3 := []byte{1, 2, 3}
    r := helper.ReverseBytes(b3)
    
    // concatenate bytes
    b4 := []byte{4, 5, 6}
    c := helper.ConcatBytes(b3, b4)

    // convert byte array to hex string
    s := helper.BytesToHex(b3)

    // convert hex string to byte array
    b5 := helper.HexToBytes(s)

    // convert ScriptHash to address string
    a := helper.ScriptHashToAddress(v1)

    // convert address string to ScriptHash
    v3 := helper.AddressToScriptHash

    ...
}
```

### "rpc" module
This module provides structs and methods which can be used to send RPC requests to and receive RPC responses from a neo node. For more information about neo RPC API, refer to [API Reference](https://docs.neo.org/docs/en-us/reference/rpc/latest-version/api.html).

Typical usage:

```golang

package sample

import "github.com/joeqian10/neo-gogogo/rpc"

func SampleMethod() {
    // create a rpc client
    var TestNetEndPoint = "http://seed1.ngd.network:20332"
    client := rpc.NewClient(TestNetEndPoint)

    // get block count
    r1 := client.GetBlockCount()
	height := r1.Result

    // get raw mempool, get all the transactions' id in this node's mempool
    r2 := client.GetRawMemPool()
    transactions := r2.Result

    // get transaction detail by its id
    r3 := client.GetRawTransaction("your transaction id string")
    tx := r3.Result

    // send raw transaction
    r4 := client.SendRawTransaction("raw transaction hex string")

    ...
}

```

### "sc" module
This module is mainly used to build smart contract scripts which can be run in a neo virtual machine. For more information about neo smart contract and virtual machine, refer to [NeoContract](https://docs.neo.org/docs/en-us/basic/technology/neocontract.html) and [NeoVM](https://docs.neo.org/docs/en-us/basic/technology/neovm.html).

Typical usage:

```golang

package sample

import "github.com/joeqian10/neo-gogogo/sc"

func SampleMethod() {
    // create a script builder
    sb := sc.NewScriptBuilder()

    // make invocation script, call a specific method from a specific contract
    scriptHash, _ := helper.UInt160FromString("b9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
    sb.MakeInvocationScript(scriptHash.Bytes(), "name", []ContractParameter{})
    bytes := sb.ToArray()

    ...
}

```

### "tx" module
This module defines different types of transactions in the neo network and also provides structs and methods for building transactions from scratch. For more information about neo transactions, refer to [Transaction](https://docs.neo.org/docs/en-us/tooldev/transaction/transaction.html).

Typical usage:

```golang

package sample

import "github.com/joeqian10/neo-gogogo/tx"

func SampleMethod() {
    // create a transaction builder
    var TestNetEndPoint = "http://seed1.ngd.network:20332"
    tb := tx.NewTransactionBuilder(TestNetEndPoint)

    // build a contract transaction
    from, _ := helper.AddressToScriptHash("APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR")
	to, _ := helper.AddressToScriptHash("AdQk428wVzpkHTxc4MP5UMdsgNdrm36dyV")
	assetId := NeoToken
	amount := helper.Fixed8FromInt64(50000000)
	ctx, _ := tb.MakeContractTransaction(from, to, assetId, amount, nil, helper.UInt160{}, helper.Fixed8FromInt64(0))
    // get the raw byte array of this transaction
    unsignedRaw := ctx.UnsignedRawTransaction()

    ...
}

```

### "wallet" module
This module defines the account and wallet in the neo network, and methods for creating an account or a wallet, signing a message/verifying signature with private/public key pair are also provided. For more information about the neo wallet, refer to [Wallet](https://docs.neo.org/docs/en-us/tooldev/wallets.html).

Typical usage:

```golang

package sample

import "github.com/joeqian10/neo-gogogo/tx"
import "github.com/joeqian10/neo-gogogo/wallet"

func SampleMethod() {
    // create an account with a random generated private key
    a1, err := wallet.NewAccount()
    // or create an account with your own private key in WIF format
    a2, err := wallet.NewAccountFromWIF("your private key in WIF format")
    // or create an account with a private key encrypted in NEP-2 standard and a passphrase
    a3, err := wallet.NewAccountFromNep2("your private key encrypted in NEP-2 standard", "your passphrase")

    // create a new wallet
    w := wallet.NewWallet()
    // add a new account into the wallet
    w.AddNewAccount()
    // or import an account from a WIF key
    w.ImportFromWIF("your account private key in WIF format")
    // or import an account from a private key encrypted in NEP-2 standard and a passphrase
    w.ImportFromNep2Key("your account private key encrypted in NEP-2 standard", "your account passphrase")
    // or simply add an existing account
    w.AddAccount(a1)

    // create a WalletHelper
    var TestNetEndPoint = "http://seed1.ngd.network:20332"
    tb := tx.NewTransactionBuilder(TestNetEndPoint)
    wh := wallet.NewWalletHelper(tb, a2)
    // transfer some neo
    wh.Transfer(tx.NeoToken, a2.Address, a3.Address, 80000)
    // claim gas
    wh.ClaimGas(a2.Address)

    ...
}

```

### "nep5" module
This module is to make life easier when dealing with NEP-5 tokens. Methods for querying basic information of a NEP-5 token, such as name, total supply, are provided. Also, it offers the ability to test run the scripts to transfer and get the balance of a NEP-5 token. For more information about NEP-5, refer to [NEP-5](https://github.com/neo-project/proposals/blob/master/nep-5.mediawiki).

Typical usage:

```golang

package sample

import "github.com/joeqian10/neo-gogogo/nep5"
import "github.com/joeqian10/neo-gogogo/wallet"

func SampleMethod() {
    // create a Nep5Helper
    var TestNetEndPoint = "http://seed1.ngd.network:20332"
    nh := nep5.NewNep5Helper(TestNetEndPoint)
    
    // get the name of a NEP-5 token
    scriptHash, _ := helper.UInt160FromString("0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
    name, err := nh.Name(scriptHash)
    
    // get the total supply of a NEP-5 token
    scriptHash, _ := helper.UInt160FromString("0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
	s, e := nh.TotalSupply(scriptHash)

    // get the balance of a NEP-5 token of an address
    scriptHash, _ := helper.UInt160FromString("0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
	address, _ := helper.AddressToScriptHash("AUrE5r4NHznrgvqoFAGhoUbu96PE5YeDZY")
	u, e := nh.BalanceOf(scriptHash, address)

    // test run the script for transfer a NEP-5 token
    scriptHash, _ := helper.UInt160FromString("0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263")
    address1, _ := helper.AddressToScriptHash("AUrE5r4NHznrgvqoFAGhoUbu96PE5YeDZY")
    address2, _ := helper.AddressToScriptHash("AdQk428wVzpkHTxc4MP5UMdsgNdrm36dyV")
	b, e := nh.Transfer(scriptHash, address1, address2, 1) 

    ...
}

```
