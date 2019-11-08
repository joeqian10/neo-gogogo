package rpc_test

import (
	"github.com/joeqian10/neo-gogogo/rpc"
	"log"
	"testing"
)

var LocalEndPoint = "http://localhost:50003" // if you want to test, you need to change this endpoint to yours
var TestNetEndPoint = "http://seed1.ngd.network:20332"

var WalletAddress = "APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR"
var AssetIdNeo = "c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b"
var Nep5ScriptHash = "14df5d02f9a52d3e92ab8cdcce5fc76c743a9b26"

func TestNewClient(t *testing.T) {
	client := rpc.NewClient(TestNetEndPoint)
	if client == nil {
		t.Fail()
	}
	log.Printf("%v", client)
}

// RpcWallet plugin required
func TestRpcClient_ClaimGas(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint) // need to open wallet
	if client == nil {
		t.Fail()
	}

	result := client.ClaimGas(WalletAddress)
	log.Printf("%+v", result)
}

func TestRpcClient_GetAccountState(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetAccountState(WalletAddress)
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetAccountState
//2019/11/04 15:41:01 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Version:0 ScriptHash:0x758ec2715fbcaeadf1b2179b11a7d980f8eb9253 Frozen:false Votes:[] Balances:[{Asset:0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b Value:100000000} {Asset:0x602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7 Value:14959.97877}]}}
//--- PASS: TestRpcClient_GetAccountState (0.33s)
//PASS

// ApplicationLogs plugin required
func TestRpcClient_GetApplicationLog(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetApplicationLog("0x0f3009286f9e59ab7cd7a232c9580a145a02cccd155a90c805e90eb9bfed3e40")
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetApplicationLog
//2019/11/04 15:40:04 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{TxId:0x0f3009286f9e59ab7cd7a232c9580a145a02cccd155a90c805e90eb9bfed3e40 Executions:[{Trigger:Application Contract:0x8f056218471dd6d253077eee22d51c9d018f90db VMState:HALT GasConsumed:3.005 Stack:[] Notifications:[{Contract:0x14df5d02f9a52d3e92ab8cdcce5fc76c743a9b26 State:{Type:Array Value:[map[type:ByteArray value:7472616e73666572726564] map[type:ByteArray value:5392ebf880d9a7119b17b2f1adaebc5f71c28e75] map[type:ByteArray value:ed53b14908f15918abb6ba8790a80387f8a5fd89] map[type:ByteArray value:80969800]]}}]}]}}
//--- PASS: TestRpcClient_GetApplicationLog (0.55s)
//PASS

func TestRpcClient_GetAssetState(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetAssetState(AssetIdNeo)
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetAssetState
//2019/11/04 15:43:06 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Version:0 Id:0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b Type:GoverningToken Name:[{Lang:zh-CN Name:小蚁股} {Lang:en Name:AntShare}] Amount:100000000 Available:100000000 Precision:0 Owner:00 Admin:Abf2qMs1pzQb8kYk9RuxtUb9jtRKJVuBJt Issuer:Abf2qMs1pzQb8kYk9RuxtUb9jtRKJVuBJt Expiration:4000000 Frozen:false}}
//--- PASS: TestRpcClient_GetAssetState (0.35s)
//PASS

// RpcWallet plugin required
func TestRpcClient_GetBalance(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetBalance(AssetIdNeo)
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetBalance
//2019/11/04 15:49:03 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Balance:100000000 Confirmed:100000000}}
//--- PASS: TestRpcClient_GetBalance (0.36s)
//PASS

func TestRpcClient_GetBestBlockHash(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetBestBlockHash()
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetBestBlockHash
//2019/11/04 15:49:53 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:0x2727ab449e02150fa66943cf6d8fcdf4af349480a558a0bdbb4eea550ffeb01f}
//--- PASS: TestRpcClient_GetBestBlockHash (0.35s)
//PASS

func TestRpcClient_GetBlockByHash(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetBlockByHash("2727ab449e02150fa66943cf6d8fcdf4af349480a558a0bdbb4eea550ffeb01f")
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetBlockByHash
//2019/11/04 15:51:15 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{RpcBlockHeader:{Hash:0x2727ab449e02150fa66943cf6d8fcdf4af349480a558a0bdbb4eea550ffeb01f Size:452 Version:0 Previousblockhash:0xbfe458b3bb7190070f76e4d09b9b677c62490970ad4673a8a6ae57350f99503b Merkleroot:0x30e1ad309fabb7dc6390074dfefe3fb375d2bdf9eb78ee4a9e7648dbd48cf9d4 Time:1572853450 Index:2022 Nonce:b1a375ca27a78e55 Nextconsensus:AKcMgU2hd9DVm7YTqbKHS1SprWS4ufMhKr CrossStatesRoot: ChainID: Script:{InvocationScript:404afc96d7f59d7c3459c6d5680eed9ff9f891ee28245219e38404318632bf08d8b5b2c13b1d8cfd92c444ec59609967cd5c0cc413261b4f608e125c083a5b214c40af65feb1e3450d4391b6bef3b7c3c9223e20183dc10529152253f0d773de0b11327868849b647f8083ab14e4b6cc2d3b11bfc645ef400e81a78d9b4ff35d311d4029dec98b568b14011f62c68d19d951645cececc5b2cee6c1e1e813d0714ce38217fbb673bc62b514a49ff0c96eda76065d1b5c4bb83d016d64f19951ffd2db48 VerificationScript:53210301f2f3c3122526f8e2330693c6256c4b325a0068bc9955713567e7a3c09498a52103460a483d65cb6139ae5b55d8a2e9f8bbf0bbfa10a407224e93b99481792d4db2210286b70fb44980dd13a54ece4b5f88ef491e501fef7fead5804a25e7e1bc5dcf732103dc31257e8f394989325d4f86f47587e405c52d1d0be9898f81991137d5a3e04b54ae} Confirmations:1 NextBlockHash:0x587ac1a5b048ed84535ebe781f5fc614f18853bcbc1a1d8da5bbb7ad51d41ab8} Tx:[{Txid:0x30e1ad309fabb7dc6390074dfefe3fb375d2bdf9eb78ee4a9e7648dbd48cf9d4 Size:10 Type:MinerTransaction Version:0 Attributes:[] Vin:[] Vout:[] SysFee:0 NetFee:0 Scripts:[] Nonce:665292373 BlockHash: Confirmations:0 Blocktime:0 Script: Gas: Claims:[]}]}}
//--- PASS: TestRpcClient_GetBlockByHash (0.32s)
//PASS

func TestRpcClient_GetBlockByIndex(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetBlockByIndex(1000)
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetBlockByIndex
//2019/11/04 15:53:04 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{RpcBlockHeader:{Hash:0x31c7521f4ac295deffb691689ecfffce687c77ac49704defa626024126b86de9 Size:452 Version:0 Previousblockhash:0xd1b1a123eb2b36161193955e5c84787917bdef8bceeae1ee77a719acbede38b9 Merkleroot:0x22b82e30630f480c7287193162d88e448aabab93eddfc42fee4dcc6b3ce9da1a Time:1570866717 Index:1000 Nonce:772f67c5bfd6b9bf Nextconsensus:AKcMgU2hd9DVm7YTqbKHS1SprWS4ufMhKr CrossStatesRoot: ChainID: Script:{InvocationScript:408201d61c3aaf87d059d7e8c4a884c022772a4d3751dedd0fe4f07d8f17bb7948bcc053ab0134441d256084c7ead4b409ff04b8614c75ab93b0ddcafad66b82e54036d46c32c36e128db1acd2f81b85c0505e1ea850aa8801ba03f6ad46d05becca99f949cd28324fc7402e88eadb49c206614618dd377fc8d53d74a2d562e5fe954060f1cfbf7f332857a7658dc8ce6c29a12b2d90925c7c2a10c2862329901c1ff16224f209fb5c7a2c22d8863feb0b531d24ab61cbe81769ed7168764ff398e382 VerificationScript:53210301f2f3c3122526f8e2330693c6256c4b325a0068bc9955713567e7a3c09498a52103460a483d65cb6139ae5b55d8a2e9f8bbf0bbfa10a407224e93b99481792d4db2210286b70fb44980dd13a54ece4b5f88ef491e501fef7fead5804a25e7e1bc5dcf732103dc31257e8f394989325d4f86f47587e405c52d1d0be9898f81991137d5a3e04b54ae} Confirmations:1023 NextBlockHash:0x6247a10d325ef423152dafcbd328b2db3ec8a68aae8d106fe82a58c4ac799825} Tx:[{Txid:0x22b82e30630f480c7287193162d88e448aabab93eddfc42fee4dcc6b3ce9da1a Size:10 Type:MinerTransaction Version:0 Attributes:[] Vin:[] Vout:[] SysFee:0 NetFee:0 Scripts:[] Nonce:3218520511 BlockHash: Confirmations:0 Blocktime:0 Script: Gas: Claims:[]}]}}
//--- PASS: TestRpcClient_GetBlockByIndex (0.34s)
//PASS

func TestRpcClient_GetBlockCount(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetBlockCount()
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetBlockCount
//2019/11/04 15:53:55 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:2023}
//--- PASS: TestRpcClient_GetBlockCount (0.33s)
//PASS

func TestRpcClient_GetBlockHeaderByHash(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetBlockHeaderByHash("2727ab449e02150fa66943cf6d8fcdf4af349480a558a0bdbb4eea550ffeb01f")
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetBlockHeaderByHash
//2019/11/04 16:16:55 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Hash:0x2727ab449e02150fa66943cf6d8fcdf4af349480a558a0bdbb4eea550ffeb01f Size:442 Version:0 Previousblockhash:0xbfe458b3bb7190070f76e4d09b9b677c62490970ad4673a8a6ae57350f99503b Merkleroot:0x30e1ad309fabb7dc6390074dfefe3fb375d2bdf9eb78ee4a9e7648dbd48cf9d4 Time:1572853450 Index:2022 Nonce:b1a375ca27a78e55 Nextconsensus:AKcMgU2hd9DVm7YTqbKHS1SprWS4ufMhKr CrossStatesRoot: ChainID: Script:{InvocationScript:404afc96d7f59d7c3459c6d5680eed9ff9f891ee28245219e38404318632bf08d8b5b2c13b1d8cfd92c444ec59609967cd5c0cc413261b4f608e125c083a5b214c40af65feb1e3450d4391b6bef3b7c3c9223e20183dc10529152253f0d773de0b11327868849b647f8083ab14e4b6cc2d3b11bfc645ef400e81a78d9b4ff35d311d4029dec98b568b14011f62c68d19d951645cececc5b2cee6c1e1e813d0714ce38217fbb673bc62b514a49ff0c96eda76065d1b5c4bb83d016d64f19951ffd2db48 VerificationScript:53210301f2f3c3122526f8e2330693c6256c4b325a0068bc9955713567e7a3c09498a52103460a483d65cb6139ae5b55d8a2e9f8bbf0bbfa10a407224e93b99481792d4db2210286b70fb44980dd13a54ece4b5f88ef491e501fef7fead5804a25e7e1bc5dcf732103dc31257e8f394989325d4f86f47587e405c52d1d0be9898f81991137d5a3e04b54ae} Confirmations:1 NextBlockHash:0x587ac1a5b048ed84535ebe781f5fc614f18853bcbc1a1d8da5bbb7ad51d41ab8}}
//--- PASS: TestRpcClient_GetBlockHeaderByHash (0.32s)
//PASS

func TestRpcClient_GetBlockHash(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetBlockHash(1000)
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetBlockHash
//2019/11/04 16:16:32 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:0x31c7521f4ac295deffb691689ecfffce687c77ac49704defa626024126b86de9}
//--- PASS: TestRpcClient_GetBlockHash (0.32s)
//PASS

// RpcSystemAssetTracker plugin required
func TestRpcClient_GetClaimable(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetClaimable("APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR")
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetClaimable
//2019/11/04 16:24:16 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Claimables:[{TxId:bc0fda55480440dbd5492de670d0fc44ecd336a20d55caeb43eedbe239ea3c65 N:0 Value:100000000 StartHeight:1870 EndHeight:2383 Generated:4104 SysFee:0 Unclaimed:4104}] Address:APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR Unclaimed:4104}}
//--- PASS: TestRpcClient_GetClaimable (0.33s)
//PASS

func TestRpcClient_GetConnectionCount(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetConnectionCount()
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetConnectionCount
//2019/11/04 16:25:18 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:3}
//--- PASS: TestRpcClient_GetConnectionCount (0.35s)
//PASS

func TestRpcClient_GetContractState(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetContractState(Nep5ScriptHash)
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetContractState
//2019/11/04 16:27:15 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Version:0 Hash:0x14df5d02f9a52d3e92ab8cdcce5fc76c743a9b26 Script:5ec56b6c766b00527ac46c766b51527ac46161680475c893e3009c6a52527ac46a52c36429006161145392ebf880d9a7119b17b2f1adaebc5f71c28e75616804efdfe6946a53527ac462920161680475c893e3609c6a54527ac46a54c36475016161680445995a5c6a55527ac46a00c30962616c616e63654f66876a56527ac46a56c36416006a51c300c361e0010154016a53527ac46245016a00c308646563696d616c73876a57527ac46a57c364110061e00100c0016a53527ac4621f016a00c3096d696e74546f6b656e876a58527ac46a58c364110061e00100de016a53527ac462f8006a00c3046e616d65876a59527ac46a59c364110061e00100c5026a53527ac462d6006a00c30673796d626f6c876a5a527ac46a5ac364110061e00100ad026a53527ac462b2006a00c312737570706f727465645374616e6461726473876a5b527ac46a5bc364110061e0010088026a53527ac46282006a00c30b746f74616c537570706c79876a5c527ac46a5cc364110061e0010084026a53527ac46259006a00c3087472616e73666572876a5d527ac46a5dc36437006a51c300c36a51c351c36a51c352c36a55c3615379517955727551727552795279547275527275e001049a026a53527ac4620d0061006a53527ac46203006a53c3616c756655c56b6c766b00527ac4616a00c36a51527ac46a51c3c001149c009c6a53527ac46a53c36439003254686520706172616d65746572206163636f756e742053484f554c442062652032302d62797465206164647265737365732e6175f06168048418d60d056173736574617ce001021e046a52527ac46a52c36a51c3617ce0010232046a54527ac46203006a54c3616c756600c56b58616c756653c56b6c766b00527ac4616a00c3616804a16ba92e6a51527ac46a51c36410006a51c361680477fd08c8620400516a52527ac46203006a52c3616c756655c56b616168048418d60d08636f6e7472616374617ce001029d036a00527ac46a00c30b746f74616c537570706c79617ce00102de036a51527ac46a51c3009e6a52527ac46a52c3640d0061006a53527ac462b000616a00c30b746f74616c537570706c79610400e1f505615272e00003d703616168048418d60d056173736574617ce0010230036a54527ac46a54c361145392ebf880d9a7119b17b2f1adaebc5f71c28e75610400e1f505615272e00003d70361610061145392ebf880d9a7119b17b2f1adaebc5f71c28e75610400e1f5056152720b7472616e7366657272656454c168124e656f2e52756e74696d652e4e6f7469667961516a53527ac46203006a53c3616c756600c56b044655434b616c756600c56b03575446616c756600c56b53c57600054e45502d35c47651054e45502d37c47652064e45502d3130c4616c756652c56b616168048418d60d08636f6e7472616374617ce0010258026a00527ac46a00c30b746f74616c537570706c79617ce0010299026a51527ac46203006a51c3616c756653c56b6c766b00527ac46c766b51527ac46c766b52527ac451616c75665fc56b6c766b00527ac46c766b51527ac46c766b52527ac46c766b53527ac4616a00c3c00114907c907c9e630f006a51c3c001149c009c620400516a57527ac46a57c3643e003754686520706172616d65746572732066726f6d20616e6420746f2053484f554c442062652032302d62797465206164647265737365732e6175f06a52c300a16a58527ac46a58c36433002c54686520706172616d6574657220616d6f756e74204d5553542062652067726561746572207468616e20302e6175f06a51c361e0010155fd009c6a59527ac46a59c3640c00006a5a527ac4622a016a00c3616804efdfe694630d006a00c36a53c39e620400006a5b527ac46a5bc3640c00006a5a527ac462fe006168048418d60d056173736574617ce00102f1006a54527ac46a54c36a00c3617ce0010205016a55527ac46a55c36a52c39f6a5c527ac46a5cc3640c00006a5a527ac462b8006a00c36a51c39c6a5d527ac46a5dc3640c00516a5a527ac4629d006a55c36a52c39c6a5e527ac46a5ec36414006a54c36a00c3617ce000029901616219006a54c36a00c36a55c36a52c394615272e000033f01616a54c36a51c3617ce0010284006a56527ac46a54c36a51c36a56c36a52c393615272e00003170161616a00c36a51c36a52c36152720b7472616e7366657272656454c168124e656f2e52756e74696d652e4e6f7469667961516a5a527ac46203006a5ac3616c756652c56b6c766b00527ac46c766b51527ac46152c5766a00c3007cc4766a51c3517cc4616c756653c56b6c766b00527ac46c766b51527ac46a00c351c301007e6a51c37e6a52527ac46a00c300c36a52c3617c68041f2e7b07616c756653c56b6c766b00527ac46c766b51527ac46a00c351c301007e6a51c37e6a52527ac46a00c300c36a52c3617c68041f2e7b07616c756654c56b6c766b00527ac46c766b51527ac46c766b52527ac46a00c351c301007e6a51c37e6a53527ac46a00c300c36a53c36a52c3615272680452a141f5616c756654c56b6c766b00527ac46c766b51527ac46c766b52527ac46a00c351c301007e6a51c37e6a53527ac46a00c300c36a53c36a52c3615272680452a141f5616c756653c56b6c766b00527ac46c766b51527ac46a00c351c301007e6a51c37e6a52527ac46a00c300c36a52c3617c6804ef7cef5d616c7566 Parameters:[String Array] Returntype:ByteArray Name:t CodeVersion:1 Author:j Email:z Description:1 Properties:{Storage:true DynamicInvoke:false}}}
//--- PASS: TestRpcClient_GetContractState (0.33s)
//PASS

// RpcNep5Tracker plugin required
func TestRpcClient_GetNep5Balances(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetNep5Balances("APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR")
	log.Printf("%+v", result)
}

func TestRpcClient_GetNep5Transfers(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetNep5Transfers("APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR")
	log.Printf("%+v", result)
}

// RpcWallet
func TestRpcClient_GetNewAddress(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetNewAddress()
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetNewAddress
//2019/11/04 16:55:01 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:AVfppvZcLeynVzsCe5W52yH925rXZARWVd}
//--- PASS: TestRpcClient_GetNewAddress (0.64s)
//PASS

func TestRpcClient_GetPeers(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetPeers()
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetPeers
//2019/11/04 16:57:57 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Unconnected:[] Bad:[] Connected:[{Address:127.0.0.1 Port:20001} {Address:127.0.0.1 Port:30001} {Address:127.0.0.1 Port:10001}]}}
//--- PASS: TestRpcClient_GetPeers (0.33s)
//PASS

func TestRpcClient_GetRawMemPool(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetRawMemPool()
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetRawMemPool
//2019/11/04 16:59:37 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:[]}
//--- PASS: TestRpcClient_GetRawMemPool (0.48s)
//PASS

func TestRpcClient_GetRawTransaction(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetRawTransaction("0x0f3009286f9e59ab7cd7a232c9580a145a02cccd155a90c805e90eb9bfed3e40")
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetRawTransaction
//2019/11/04 17:10:37 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Txid:0x0f3009286f9e59ab7cd7a232c9580a145a02cccd155a90c805e90eb9bfed3e40 Size:218 Type:InvocationTransaction Version:1 Attributes:[{Usage:Script Data:5392ebf880d9a7119b17b2f1adaebc5f71c28e75}] Vin:[] Vout:[] SysFee:0 NetFee:0 Scripts:[{Invocation:4017a5a8de6c2fbd2c3f9a66bfdd40cde769e27fc2fe5fcabcf25a866a7e42bacc4619c02133284d3d67f7c889398b6b585dc60908b75ffa9d64d7bc333bc4de44 Verification:2103dc623d556a79437ffefa0133516ec91545f363b772149901ac589b5e927ad55fac}] Nonce:0 BlockHash:0x9d88c792a68f97e5fa4076adc5341c2e7efd8298386610cfc83e25ac35837b08 Confirmations:822 Blocktime:1572853150 Script:048096980014ed53b14908f15918abb6ba8790a80387f8a5fd89145392ebf880d9a7119b17b2f1adaebc5f71c28e7553c1087472616e7366657267269b3a746cc75fcedc8cab923e2da5f9025ddf14f1 Gas:0 Claims:[]}}
//--- PASS: TestRpcClient_GetRawTransaction (0.33s)
//PASS

func TestRpcClient_GetStorage(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetStorage(Nep5ScriptHash, "636f6e747261637400746f74616c537570706c79") // key= "contract" + "00" + "totalSupply"
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetStorage
//2019/11/04 17:43:23 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:00e1f505}
//--- PASS: TestRpcClient_GetStorage (0.36s)
//PASS

func TestRpcClient_GetTransactionHeight(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetTransactionHeight("0x0f3009286f9e59ab7cd7a232c9580a145a02cccd155a90c805e90eb9bfed3e40")
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetTransactionHeight
//2019/11/04 17:24:40 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:1978}
//--- PASS: TestRpcClient_GetTransactionHeight (0.31s)
//PASS

func TestRpcClient_GetTxOut(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetTxOut("0x011863ed7e21e73eba017e1887ae3085771661ae624d631190e4457ef4202ad1", 0)
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetTxOut
//2019/11/04 17:34:49 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{N:0 Asset:0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b Value:100000000 Address:APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR}}
//--- PASS: TestRpcClient_GetTxOut (0.32s)
//PASS

// need RpcSystemAssetTracker
func TestRpcClient_GetUnclaimed(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetUnclaimed(WalletAddress)
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetUnclaimed
//2019/11/04 17:47:51 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Available:4104 Unavailable:4088 Unclaimed:8192}}
//--- PASS: TestRpcClient_GetUnclaimed (0.36s)
//PASS

func TestRpcClient_GetUnclaimedGas(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetUnclaimedGas()
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetUnclaimedGas
//2019/11/04 17:48:41 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Available:4104 Unavailable:4096}}
//--- PASS: TestRpcClient_GetUnclaimedGas (0.37s)
//PASS

func TestRpcClient_GetUnspents(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetUnspents(WalletAddress)
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetUnspents
//2019/11/04 17:51:04 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Balance:[{Unspent:[{Txid:883524d20d2f43789453a2e3b4d4cb7726bbceb71b6fffadc4916654da9ad90a N:0 Value:6996} {Txid:db961a3cd480934dc31c83f7665f859310d7b2c51d0a441df7c3ae2a87e14a28 N:0 Value:1120} {Txid:520bfe6fea7c4dd31f3cc3f3b5b9098dbc6d7b671575c7c2c28da53b4b79167f N:0 Value:81.97877} {Txid:9251ea1b8eb58ab908a33523e6712f63b1d78386b3a05011914d23feda973f87 N:0 Value:6762}] AssetHash:602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7 Asset:GAS AssetSymbol:GAS Amount:0} {Unspent:[{Txid:011863ed7e21e73eba017e1887ae3085771661ae624d631190e4457ef4202ad1 N:0 Value:1e+08}] AssetHash:c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b Asset:NEO AssetSymbol:NEO Amount:100000000}] Address:APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR}}
//--- PASS: TestRpcClient_GetUnspents (0.35s)
//PASS

func TestRpcClient_GetValidators(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetValidators()
	log.Printf("%+v", result)
}

func TestRpcClient_GetVersion(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetVersion()
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetVersion
//2019/11/04 17:58:04 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Port:50001 Nonce:1717092459 Useragent:/Neo:2.10.3/}}
//--- PASS: TestRpcClient_GetVersion (0.35s)
//PASS

func TestRpcClient_GetWalletHeight(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.GetWalletHeight()
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_GetWalletHeight
//2019/11/04 17:58:54 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:2894}
//--- PASS: TestRpcClient_GetWalletHeight (0.33s)
//PASS

// TODO
func TestRpcClient_ImportPrivKey(t *testing.T) {

}

func TestRpcClient_InvokeFunction(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.InvokeFunction(Nep5ScriptHash, "name")
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_InvokeFunction
//2019/11/04 18:07:27 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Script:00c1046e616d6567269b3a746cc75fcedc8cab923e2da5f9025ddf14 State:HALT GasConsumed:0.094 Stack:[{Type:ByteArray Value:4655434b}]}}
//--- PASS: TestRpcClient_InvokeFunction (0.38s)
//PASS

func TestRpcClient_InvokeScript(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.InvokeScript("00c1046e616d6567269b3a746cc75fcedc8cab923e2da5f9025ddf14")
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_InvokeScript
//2019/11/04 18:09:02 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Script:00c1046e616d6567269b3a746cc75fcedc8cab923e2da5f9025ddf14 State:HALT GasConsumed:0.094 Stack:[{Type:ByteArray Value:4655434b}]}}
//--- PASS: TestRpcClient_InvokeScript (0.34s)
//PASS

// RpcWallet
func TestRpcClient_ListAddress(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.ListAddress()
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_ListAddress
//2019/11/04 18:21:27 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:[{Address:APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR HasKey:true Label: WatchOnly:false}]}
//--- PASS: TestRpcClient_ListAddress (0.35s)
//PASS

func TestRpcClient_ListPlugins(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.ListPlugins()
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_ListPlugins
//2019/11/04 18:24:03 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:[{Name:ApplicationLogs Version:2.10.3.0 Interfaces:[IRpcPlugin IPersistencePlugin]} {Name:RpcNep5Tracker Version:2.10.3.0 Interfaces:[IPersistencePlugin IRpcPlugin]} {Name:RpcSystemAssetTrackerPlugin Version:2.10.3.0 Interfaces:[IPersistencePlugin IRpcPlugin]} {Name:RpcWallet Version:2.10.3.0 Interfaces:[IRpcPlugin]} {Name:SimplePolicyPlugin Version:2.10.3.0 Interfaces:[ILogPlugin IPolicyPlugin]}]}
//--- PASS: TestRpcClient_ListPlugins (0.32s)
//PASS

func TestRpcClient_SendFrom(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.SendFrom(AssetIdNeo, WalletAddress, WalletAddress, 100000000, 0, WalletAddress)
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_SendFrom
//2019/11/04 18:44:09 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Txid:0x5bfbda6774a96430c7a67e590c37c00b4c167c22268878d49a46849b15fa55b0 Size:202 Type:ContractTransaction Version:0 Attributes:[] Vin:[{Txid:0xef69ebd50ca876857e9f44c9901460d81dbfb5b8eaab2b7975325b82990bf3c2 Vout:0}] Vout:[{N:0 Asset:0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b Value:100000000 Address:APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR}] SysFee:0 NetFee:0 Scripts:[{Invocation:40bcb785e9c3a8c9df5dd77f5d9007607ebbf45ed52eaec63a6cf366432e31e6ef40a559ba4c8594ec8d2bbd8f5a5e194cee92017868b3ee1e0d5a5cf15be68b12 Verification:2103dc623d556a79437ffefa0133516ec91545f363b772149901ac589b5e927ad55fac}] Nonce:0 BlockHash: Confirmations:0 Blocktime:0 Script: Gas: Claims:[]}}
//--- PASS: TestRpcClient_SendFrom (0.34s)
//PASS

// TODO
func TestRpcClient_SendRawTransaction(t *testing.T) {

}

func TestRpcClient_SendToAddress(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.SendToAddress(AssetIdNeo, WalletAddress, 100000000, 0, WalletAddress)
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_SendToAddress
//2019/11/04 18:44:34 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Txid:0x21bf1497ff54aef61e389d2180f5b4174963e52eaa1c5e717da14ba977c9e7f9 Size:202 Type:ContractTransaction Version:0 Attributes:[] Vin:[{Txid:0x5bfbda6774a96430c7a67e590c37c00b4c167c22268878d49a46849b15fa55b0 Vout:0}] Vout:[{N:0 Asset:0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b Value:100000000 Address:APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR}] SysFee:0 NetFee:0 Scripts:[{Invocation:40420f043d3c4d46b0979c3243329bcf9e3870eab68ef827243db81e0b04f62aa0a081efbee5d1368bf9e4d97b3edc8e7ba700a94d1799a4933a47e8947935b200 Verification:2103dc623d556a79437ffefa0133516ec91545f363b772149901ac589b5e927ad55fac}] Nonce:0 BlockHash: Confirmations:0 Blocktime:0 Script: Gas: Claims:[]}}
//--- PASS: TestRpcClient_SendToAddress (0.34s)
//PASS

func TestRpcClient_SubmitBlock(t *testing.T) {

}

func TestRpcClient_ValidateAddress(t *testing.T) {
	client := rpc.NewClient(LocalEndPoint)
	if client == nil {
		t.Fail()
	}

	result := client.ValidateAddress(WalletAddress)
	log.Printf("%+v", result)
}

//=== RUN   TestRpcClient_ValidateAddress
//2019/11/04 18:46:00 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Address:APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR IsValid:true}}
//--- PASS: TestRpcClient_ValidateAddress (0.35s)
//PASS
