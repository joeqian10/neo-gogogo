package rpc_test

import (
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)
/*
If you want to run all the tests, you'd better have a private net prepared for testing,
since a private net is more flexible and you will have plenty of neo and gas to spend.
You need to change the LocalEndPoint to your node, and install all required plugins.
*/

var LocalEndPoint = "http://localhost:50003"
var TestNetEndPoint = "http://seed1.ngd.network:20332"

var LocalClient = rpc.NewClient(LocalEndPoint)
var TestNetClient  = rpc.NewClient(TestNetEndPoint)

var LocalWalletAddress = "APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR"
var TestNetWalletAddress = "AUrE5r4NHznrgvqoFAGhoUbu96PE5YeDZY"

var AssetIdNeo = "c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b"
var Nep5ScriptHash = "0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263"



func TestNewClient(t *testing.T) {
	client := rpc.NewClient(TestNetEndPoint)
	if client == nil {
		t.Fail()
	}
	url := client.Endpoint
	//log.Printf("%s", url.Fragment)
	//log.Printf("%s", url.Host)
	//log.Printf("%s", url.Opaque)
	//log.Printf("%s", url.Path)
	//log.Printf("%s", url.RawPath)
	//log.Printf("%s", url.RawQuery)
	//log.Printf("%s", url.Scheme)
	//log.Printf("%v", client)
	assert.Equal(t, "seed1.ngd.network:20332", url.Host)
	assert.Equal(t, "http", url.Scheme)
}

// RpcWallet plugin required
func TestRpcClient_ClaimGas(t *testing.T) {
	response := LocalClient.ClaimGas(LocalWalletAddress)
	//log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t, "ClaimTransaction", r.Type)
	assert.Equal(t, 0, r.Version)
}

func TestRpcClient_GetAccountState(t *testing.T) {
	response := LocalClient.GetAccountState(LocalWalletAddress)
	//log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t,0, r.Version)
	assert.Equal(t, "0x758ec2715fbcaeadf1b2179b11a7d980f8eb9253", r.ScriptHash)
}

//=== RUN   TestRpcClient_GetAccountState
//2019/11/04 15:41:01 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Version:0 ScriptHash:0x758ec2715fbcaeadf1b2179b11a7d980f8eb9253 Frozen:false Votes:[] Balances:[{Asset:0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b Value:100000000} {Asset:0x602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7 Value:14959.97877}]}}
//--- PASS: TestRpcClient_GetAccountState (0.33s)
//PASS

// ApplicationLogs plugin required
func TestRpcClient_GetApplicationLog(t *testing.T) {
	response := TestNetClient.GetApplicationLog("0xca159430e3d72227c06a3880244111aea0368ecc09fa8c2eade001a1bbcc7d4a")
	//log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t, "0xca159430e3d72227c06a3880244111aea0368ecc09fa8c2eade001a1bbcc7d4a", r.TxId)
	e := r.Executions[0]
	assert.Equal(t, "Application", e.Trigger)
	assert.Equal(t, "0x003bd113b3bc841657f3a84db8546daa6e4953c3", e.Contract)
	assert.Equal(t, "HALT", e.VMState)
	n := e.Notifications[0]
	assert.Equal(t, "0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263", n.Contract)
}

//=== RUN   TestRpcClient_GetApplicationLog
//2019/11/07 18:48:32 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{TxId:0xca159430e3d72227c06a3880244111aea0368ecc09fa8c2eade001a1bbcc7d4a Executions:[{Trigger:Application Contract:0x003bd113b3bc841657f3a84db8546daa6e4953c3 VMState:HALT GasConsumed:2.855 Stack:[{Type:Integer Value:1}] Notifications:[{Contract:0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263 State:{Type:Array Value:[map[type:ByteArray value:7472616e73666572] map[type:ByteArray value:5c564ab204122ddce30eb9a6accbfa23b27cc3ac] map[type:ByteArray value:8f6c5be89c0cb6579e44a8bf9bfd2ecbcc11dfdc] map[type:ByteArray value:00203d88792d]]}}]}]}}
//--- PASS: TestRpcClient_GetApplicationLog (0.54s)
//PASS

func TestRpcClient_GetAssetState(t *testing.T) {
	response := TestNetClient.GetAssetState(AssetIdNeo)
	//log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t, 0, r.Version)
	assert.Equal(t, "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b", r.Id)
	assert.Equal(t, "GoverningToken", r.Type)
	assert.Equal(t, "AntShare", r.Name[1].Name)
	assert.Equal(t, "100000000", r.Amount)
	assert.Equal(t, 0, r.Precision)
	assert.Equal(t, "Abf2qMs1pzQb8kYk9RuxtUb9jtRKJVuBJt", r.Admin)
	assert.Equal(t, 4000000, r.Expiration)
}

//=== RUN   TestRpcClient_GetAssetState
//2019/11/04 15:43:06 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Version:0 Id:0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b Type:GoverningToken Name:[{Lang:zh-CN Name:小蚁股} {Lang:en Name:AntShare}] Amount:100000000 Available:100000000 Precision:0 Owner:00 Admin:Abf2qMs1pzQb8kYk9RuxtUb9jtRKJVuBJt Issuer:Abf2qMs1pzQb8kYk9RuxtUb9jtRKJVuBJt Expiration:4000000 Frozen:false}}
//--- PASS: TestRpcClient_GetAssetState (0.35s)
//PASS

// RpcWallet plugin required
func TestRpcClient_GetBalance(t *testing.T) {
	response := LocalClient.GetBalance(AssetIdNeo)
	//log.Printf("%+v", response)
	r:= response.Result
	assert.Equal(t, 100000000, r.Balance)
}

//=== RUN   TestRpcClient_GetBalance
//2019/11/04 15:49:03 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Balance:100000000 Confirmed:100000000}}
//--- PASS: TestRpcClient_GetBalance (0.36s)
//PASS

// BestBlockHash is changing, so you may stop the chain to test this API, and change the best block hash to yours
func TestRpcClient_GetBestBlockHash(t *testing.T) {
	response := LocalClient.GetBestBlockHash()
	//log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t, "0x2727ab449e02150fa66943cf6d8fcdf4af349480a558a0bdbb4eea550ffeb01f", r)
}

//=== RUN   TestRpcClient_GetBestBlockHash
//2019/11/04 15:49:53 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:0x2727ab449e02150fa66943cf6d8fcdf4af349480a558a0bdbb4eea550ffeb01f}
//--- PASS: TestRpcClient_GetBestBlockHash (0.35s)
//PASS

func TestRpcClient_GetBlockByHash(t *testing.T) {
	response := TestNetClient.GetBlockByHash("035212da3f0e73cd41e3f6e22ccbedaac064e4150ad6dd2bed3eeff420be3179")
	//log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t, "0x035212da3f0e73cd41e3f6e22ccbedaac064e4150ad6dd2bed3eeff420be3179", r.Hash)
	assert.Equal(t, 1521, r.Size)
	assert.Equal(t, 0, r.Version)
	assert.Equal(t, "0xff5396db837d368acd334c19f98b7bc8885b5efcbd85fa02e6a5558c4966e840", r.Previousblockhash)
	assert.Equal(t, "0x3216ea4203e7a90d188cc97eabbfa0bbfc6589debbcacbc43eeff0952b380757", r.Merkleroot)
	assert.Equal(t, 1573123342, r.Time)
	assert.Equal(t, 3386365, r.Index)
	assert.Equal(t, "a6e6d82b50273b82", r.Nonce)
	assert.Equal(t, "AUNSizuErA3dv1a2ag2ozvikkQS7hhPY1X", r.Nextconsensus)
}

//=== RUN   TestRpcClient_GetBlockByHash
//2019/11/08 11:14:42 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{RpcBlockHeader:{Hash:0x035212da3f0e73cd41e3f6e22ccbedaac064e4150ad6dd2bed3eeff420be3179 Size:1521 Version:0 Previousblockhash:0xff5396db837d368acd334c19f98b7bc8885b5efcbd85fa02e6a5558c4966e840 Merkleroot:0x3216ea4203e7a90d188cc97eabbfa0bbfc6589debbcacbc43eeff0952b380757 Time:1573123342 Index:3386365 Nonce:a6e6d82b50273b82 Nextconsensus:AUNSizuErA3dv1a2ag2ozvikkQS7hhPY1X CrossStatesRoot: ChainID: Script:{InvocationScript:40c5ad2bbbcbb76fa9bdb6bd19da2b37f6cf0c12fe2e1471da1c5a1af983c706fc8ecdfb90d031b26e89e6bf2159004c9bed89e435d4f672013c4f90d2d6ae026840d61a6fe68741138f3e65b762a0fd858ca46f8a8bcd433de08ef15a2272dd790eb3fb4f04ad55dcd07b58869dad2a43a48abca3b30f49325bc1d3a7673257dbb04055488c2bd94b99f479f0c42aa2bf167ece07484dce3a217c9f4246893168d6b40e20461f9115d9d7c5995271df4c472894af4b33fdc0116f1da63de21a378c32409b5b4216cfd7bc8442893971f33348ba63a231988de7379bd4c59fdb1bad783d3934e53cbd91f44e06c591354f9dd8825c30031ac2370c762a8e818ca24c6c1540d80ea89c01aed1e43ccc93be261613181d0130c2db5afb6b198d8d2655878a743b806c2d1f915e982b1dda8bf855a148051d05d9285ab0e9d6ba0c5b07a8eecd VerificationScript:552103028007d683ceb4dc9084300d0cf16fe6d47a726e586bf3d63559cec13305565221030ef96257401b803da5dd201233e2be828795672b775dd674d69df83f7aec1e3621025bdf3f181f53e9696227843950deb72dcd374ded17c057159513c3d0abe20b64210266b588e350ab63b850e55dbfed0feeda44410a30966341b371014b803a15af072103c089d7122b840a4935234e82e26ae5efd0c2acb627239dc9f207311337b6f2c12103fd95a9cb3098e6447d0de9f76cc97fd5e36830f9c7044457c15a0e81316bf28f2103fea219d4ccfd7641cebbb2439740bb4bd7c4730c1abd6ca1dc44386533816df957ae} Confirmations:3495 NextBlockHash:0x7f514b6d785b52adfeee56919d9deb12059516145aaae36f997cd79890c11bac} Tx:[{Txid:0x75f1de0f6aaab785138fb8a5183018d25465278eb38c44917db87b61a7f1c588 Size:10 Type:MinerTransaction Version:0 Attributes:[] Vin:[] Vout:[] SysFee:0 NetFee:0 Scripts:[] Nonce:1344748418 BlockHash: Confirmations:0 Blocktime:0 Script: Gas: Claims:[]} {Txid:0x18147d0916e1f2fbbcc26a3bb5fd593b90ea86c8a27be4496eebbccb8fe99de9 Size:247 Type:InvocationTransaction Version:1 Attributes:[{Usage:Script Data:10d46912932d6ebcd1d3c4a27a1a8ea77e68ac95} {Usage:Remark Data:0000016e45752e389b187108}] Vin:[] Vout:[] SysFee:0 NetFee:0 Scripts:[{Invocation:40f24367766e1fc0e1a40a9565d7395d5661d29f058668d1374077c4d8a217de9c8848509785f79eda1be3afb8881a29993aeb973963281638277e68ce64b822f7 Verification:2102b1ca89d1ac9006a795e35b92dc801f42eff7d05626ecfcf243e20aff4cb79a4aac}] Nonce:0 BlockHash: Confirmations:0 Blocktime:0 Script:1410d46912932d6ebcd1d3c4a27a1a8ea77e68ac950020000101001a000010020000000056054b2042000019c5420830b10420500444bb53c11063726561746550726f6d6f437574696567f55b45d0235e1b009eb6cffc25a56d338d2c39d3 Gas:0 Claims:[]} {Txid:0x452d52a1e8963746e80ea92a45660d85aeb9d6cd04e0065439fc95270db1810d Size:196 Type:InvocationTransaction Version:1 Attributes:[{Usage:Script Data:10d46912932d6ebcd1d3c4a27a1a8ea77e68ac95} {Usage:Remark Data:0000016e457520059c8d6488}] Vin:[] Vout:[] SysFee:0 NetFee:0 Scripts:[{Invocation:40da7201b8a684d32a1c6d5a4665ef5f99ef50cfab2e6b42eb9972fa8edbeef8646eb46bb1d0caa74080d898f5affb3f6d67e445691e0aa5f0af71994f5976cebf Verification:2102b1ca89d1ac9006a795e35b92dc801f42eff7d05626ecfcf243e20aff4cb79a4aac}] Nonce:0 BlockHash: Confirmations:0 Blocktime:0 Script:00029d0652c1106368616e676547656e65726174696f6e67f55b45d0235e1b009eb6cffc25a56d338d2c39d3 Gas:0 Claims:[]} {Txid:0x45ea381e940f0e089a72f07f1a05d4bfa2be85410e96eea472c55e7b11872d5a Size:196 Type:InvocationTransaction Version:1 Attributes:[{Usage:Script Data:10d46912932d6ebcd1d3c4a27a1a8ea77e68ac95} {Usage:Remark Data:0000016e45751f72959c1668}] Vin:[] Vout:[] SysFee:0 NetFee:0 Scripts:[{Invocation:40ae6931e6fa6398e0ccb448419debd96b2ce7ef76d4f4a172ada527f3da28d2fc77f2ea2aac04186c6ad575cd17d117820ee6d3395f1ebfb900b29972b69d20ac Verification:2102b1ca89d1ac9006a795e35b92dc801f42eff7d05626ecfcf243e20aff4cb79a4aac}] Nonce:0 BlockHash: Confirmations:0 Blocktime:0 Script:0002460552c1106368616e676547656e65726174696f6e67f55b45d0235e1b009eb6cffc25a56d338d2c39d3 Gas:0 Claims:[]} {Txid:0x82ad53594683ca5fc44e658fc1ed265b86070e4466417df3f9eadf5e207b87b2 Size:196 Type:InvocationTransaction Version:1 Attributes:[{Usage:Script Data:10d46912932d6ebcd1d3c4a27a1a8ea77e68ac95} {Usage:Remark Data:0000016e457520981613af83}] Vin:[] Vout:[] SysFee:0 NetFee:0 Scripts:[{Invocation:404e82e563c8a87155bcfa805e02f45f293d2568632a6aebb8eef6f4a4e51328b43aba45e363fac26461ac821bdc633e4ae9f5836b628820f8b382924c213b6a53 Verification:2102b1ca89d1ac9006a795e35b92dc801f42eff7d05626ecfcf243e20aff4cb79a4aac}] Nonce:0 BlockHash: Confirmations:0 Blocktime:0 Script:0002800552c1106368616e676547656e65726174696f6e67f55b45d0235e1b009eb6cffc25a56d338d2c39d3 Gas:0 Claims:[]}]}}
//--- PASS: TestRpcClient_GetBlockByHash (0.54s)
//PASS

func TestRpcClient_GetBlockByIndex(t *testing.T) {
	response := TestNetClient.GetBlockByIndex(3386365)
	//log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t, "0x035212da3f0e73cd41e3f6e22ccbedaac064e4150ad6dd2bed3eeff420be3179", r.Hash)
	assert.Equal(t, 1521, r.Size)
	assert.Equal(t, 0, r.Version)
	assert.Equal(t, "0xff5396db837d368acd334c19f98b7bc8885b5efcbd85fa02e6a5558c4966e840", r.Previousblockhash)
	assert.Equal(t, "0x3216ea4203e7a90d188cc97eabbfa0bbfc6589debbcacbc43eeff0952b380757", r.Merkleroot)
	assert.Equal(t, 1573123342, r.Time)
	assert.Equal(t, 3386365, r.Index)
	assert.Equal(t, "a6e6d82b50273b82", r.Nonce)
	assert.Equal(t, "AUNSizuErA3dv1a2ag2ozvikkQS7hhPY1X", r.Nextconsensus)
}

//=== RUN   TestRpcClient_GetBlockByIndex
//2019/11/08 11:16:03 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{RpcBlockHeader:{Hash:0x035212da3f0e73cd41e3f6e22ccbedaac064e4150ad6dd2bed3eeff420be3179 Size:1521 Version:0 Previousblockhash:0xff5396db837d368acd334c19f98b7bc8885b5efcbd85fa02e6a5558c4966e840 Merkleroot:0x3216ea4203e7a90d188cc97eabbfa0bbfc6589debbcacbc43eeff0952b380757 Time:1573123342 Index:3386365 Nonce:a6e6d82b50273b82 Nextconsensus:AUNSizuErA3dv1a2ag2ozvikkQS7hhPY1X CrossStatesRoot: ChainID: Script:{InvocationScript:40c5ad2bbbcbb76fa9bdb6bd19da2b37f6cf0c12fe2e1471da1c5a1af983c706fc8ecdfb90d031b26e89e6bf2159004c9bed89e435d4f672013c4f90d2d6ae026840d61a6fe68741138f3e65b762a0fd858ca46f8a8bcd433de08ef15a2272dd790eb3fb4f04ad55dcd07b58869dad2a43a48abca3b30f49325bc1d3a7673257dbb04055488c2bd94b99f479f0c42aa2bf167ece07484dce3a217c9f4246893168d6b40e20461f9115d9d7c5995271df4c472894af4b33fdc0116f1da63de21a378c32409b5b4216cfd7bc8442893971f33348ba63a231988de7379bd4c59fdb1bad783d3934e53cbd91f44e06c591354f9dd8825c30031ac2370c762a8e818ca24c6c1540d80ea89c01aed1e43ccc93be261613181d0130c2db5afb6b198d8d2655878a743b806c2d1f915e982b1dda8bf855a148051d05d9285ab0e9d6ba0c5b07a8eecd VerificationScript:552103028007d683ceb4dc9084300d0cf16fe6d47a726e586bf3d63559cec13305565221030ef96257401b803da5dd201233e2be828795672b775dd674d69df83f7aec1e3621025bdf3f181f53e9696227843950deb72dcd374ded17c057159513c3d0abe20b64210266b588e350ab63b850e55dbfed0feeda44410a30966341b371014b803a15af072103c089d7122b840a4935234e82e26ae5efd0c2acb627239dc9f207311337b6f2c12103fd95a9cb3098e6447d0de9f76cc97fd5e36830f9c7044457c15a0e81316bf28f2103fea219d4ccfd7641cebbb2439740bb4bd7c4730c1abd6ca1dc44386533816df957ae} Confirmations:3500 NextBlockHash:0x7f514b6d785b52adfeee56919d9deb12059516145aaae36f997cd79890c11bac} Tx:[{Txid:0x75f1de0f6aaab785138fb8a5183018d25465278eb38c44917db87b61a7f1c588 Size:10 Type:MinerTransaction Version:0 Attributes:[] Vin:[] Vout:[] SysFee:0 NetFee:0 Scripts:[] Nonce:1344748418 BlockHash: Confirmations:0 Blocktime:0 Script: Gas: Claims:[]} {Txid:0x18147d0916e1f2fbbcc26a3bb5fd593b90ea86c8a27be4496eebbccb8fe99de9 Size:247 Type:InvocationTransaction Version:1 Attributes:[{Usage:Script Data:10d46912932d6ebcd1d3c4a27a1a8ea77e68ac95} {Usage:Remark Data:0000016e45752e389b187108}] Vin:[] Vout:[] SysFee:0 NetFee:0 Scripts:[{Invocation:40f24367766e1fc0e1a40a9565d7395d5661d29f058668d1374077c4d8a217de9c8848509785f79eda1be3afb8881a29993aeb973963281638277e68ce64b822f7 Verification:2102b1ca89d1ac9006a795e35b92dc801f42eff7d05626ecfcf243e20aff4cb79a4aac}] Nonce:0 BlockHash: Confirmations:0 Blocktime:0 Script:1410d46912932d6ebcd1d3c4a27a1a8ea77e68ac950020000101001a000010020000000056054b2042000019c5420830b10420500444bb53c11063726561746550726f6d6f437574696567f55b45d0235e1b009eb6cffc25a56d338d2c39d3 Gas:0 Claims:[]} {Txid:0x452d52a1e8963746e80ea92a45660d85aeb9d6cd04e0065439fc95270db1810d Size:196 Type:InvocationTransaction Version:1 Attributes:[{Usage:Script Data:10d46912932d6ebcd1d3c4a27a1a8ea77e68ac95} {Usage:Remark Data:0000016e457520059c8d6488}] Vin:[] Vout:[] SysFee:0 NetFee:0 Scripts:[{Invocation:40da7201b8a684d32a1c6d5a4665ef5f99ef50cfab2e6b42eb9972fa8edbeef8646eb46bb1d0caa74080d898f5affb3f6d67e445691e0aa5f0af71994f5976cebf Verification:2102b1ca89d1ac9006a795e35b92dc801f42eff7d05626ecfcf243e20aff4cb79a4aac}] Nonce:0 BlockHash: Confirmations:0 Blocktime:0 Script:00029d0652c1106368616e676547656e65726174696f6e67f55b45d0235e1b009eb6cffc25a56d338d2c39d3 Gas:0 Claims:[]} {Txid:0x45ea381e940f0e089a72f07f1a05d4bfa2be85410e96eea472c55e7b11872d5a Size:196 Type:InvocationTransaction Version:1 Attributes:[{Usage:Script Data:10d46912932d6ebcd1d3c4a27a1a8ea77e68ac95} {Usage:Remark Data:0000016e45751f72959c1668}] Vin:[] Vout:[] SysFee:0 NetFee:0 Scripts:[{Invocation:40ae6931e6fa6398e0ccb448419debd96b2ce7ef76d4f4a172ada527f3da28d2fc77f2ea2aac04186c6ad575cd17d117820ee6d3395f1ebfb900b29972b69d20ac Verification:2102b1ca89d1ac9006a795e35b92dc801f42eff7d05626ecfcf243e20aff4cb79a4aac}] Nonce:0 BlockHash: Confirmations:0 Blocktime:0 Script:0002460552c1106368616e676547656e65726174696f6e67f55b45d0235e1b009eb6cffc25a56d338d2c39d3 Gas:0 Claims:[]} {Txid:0x82ad53594683ca5fc44e658fc1ed265b86070e4466417df3f9eadf5e207b87b2 Size:196 Type:InvocationTransaction Version:1 Attributes:[{Usage:Script Data:10d46912932d6ebcd1d3c4a27a1a8ea77e68ac95} {Usage:Remark Data:0000016e457520981613af83}] Vin:[] Vout:[] SysFee:0 NetFee:0 Scripts:[{Invocation:404e82e563c8a87155bcfa805e02f45f293d2568632a6aebb8eef6f4a4e51328b43aba45e363fac26461ac821bdc633e4ae9f5836b628820f8b382924c213b6a53 Verification:2102b1ca89d1ac9006a795e35b92dc801f42eff7d05626ecfcf243e20aff4cb79a4aac}] Nonce:0 BlockHash: Confirmations:0 Blocktime:0 Script:0002800552c1106368616e676547656e65726174696f6e67f55b45d0235e1b009eb6cffc25a56d338d2c39d3 Gas:0 Claims:[]}]}}
//--- PASS: TestRpcClient_GetBlockByIndex (0.52s)
//PASS

// You may stop the chain to test this API on your private net.
func TestRpcClient_GetBlockCount(t *testing.T) {
	response := LocalClient.GetBlockCount()
	//log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t, 2023, r)
}

//=== RUN   TestRpcClient_GetBlockCount
//2019/11/04 15:53:55 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:2023}
//--- PASS: TestRpcClient_GetBlockCount (0.33s)
//PASS

func TestRpcClient_GetBlockHeaderByHash(t *testing.T) {
	response := TestNetClient.GetBlockHeaderByHash("035212da3f0e73cd41e3f6e22ccbedaac064e4150ad6dd2bed3eeff420be3179")
	log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t, "0x035212da3f0e73cd41e3f6e22ccbedaac064e4150ad6dd2bed3eeff420be3179", r.Hash)
	assert.Equal(t, 676, r.Size)
	assert.Equal(t, 0, r.Version)
	assert.Equal(t, "0xff5396db837d368acd334c19f98b7bc8885b5efcbd85fa02e6a5558c4966e840", r.Previousblockhash)
	assert.Equal(t, "0x3216ea4203e7a90d188cc97eabbfa0bbfc6589debbcacbc43eeff0952b380757", r.Merkleroot)
	assert.Equal(t, 1573123342, r.Time)
	assert.Equal(t, 3386365, r.Index)
	assert.Equal(t, "a6e6d82b50273b82", r.Nonce)
	assert.Equal(t, "AUNSizuErA3dv1a2ag2ozvikkQS7hhPY1X", r.Nextconsensus)
}

//=== RUN   TestRpcClient_GetBlockHeaderByHash
//2019/11/08 11:21:10 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Hash:0x035212da3f0e73cd41e3f6e22ccbedaac064e4150ad6dd2bed3eeff420be3179 Size:676 Version:0 Previousblockhash:0xff5396db837d368acd334c19f98b7bc8885b5efcbd85fa02e6a5558c4966e840 Merkleroot:0x3216ea4203e7a90d188cc97eabbfa0bbfc6589debbcacbc43eeff0952b380757 Time:1573123342 Index:3386365 Nonce:a6e6d82b50273b82 Nextconsensus:AUNSizuErA3dv1a2ag2ozvikkQS7hhPY1X CrossStatesRoot: ChainID: Script:{InvocationScript:40c5ad2bbbcbb76fa9bdb6bd19da2b37f6cf0c12fe2e1471da1c5a1af983c706fc8ecdfb90d031b26e89e6bf2159004c9bed89e435d4f672013c4f90d2d6ae026840d61a6fe68741138f3e65b762a0fd858ca46f8a8bcd433de08ef15a2272dd790eb3fb4f04ad55dcd07b58869dad2a43a48abca3b30f49325bc1d3a7673257dbb04055488c2bd94b99f479f0c42aa2bf167ece07484dce3a217c9f4246893168d6b40e20461f9115d9d7c5995271df4c472894af4b33fdc0116f1da63de21a378c32409b5b4216cfd7bc8442893971f33348ba63a231988de7379bd4c59fdb1bad783d3934e53cbd91f44e06c591354f9dd8825c30031ac2370c762a8e818ca24c6c1540d80ea89c01aed1e43ccc93be261613181d0130c2db5afb6b198d8d2655878a743b806c2d1f915e982b1dda8bf855a148051d05d9285ab0e9d6ba0c5b07a8eecd VerificationScript:552103028007d683ceb4dc9084300d0cf16fe6d47a726e586bf3d63559cec13305565221030ef96257401b803da5dd201233e2be828795672b775dd674d69df83f7aec1e3621025bdf3f181f53e9696227843950deb72dcd374ded17c057159513c3d0abe20b64210266b588e350ab63b850e55dbfed0feeda44410a30966341b371014b803a15af072103c089d7122b840a4935234e82e26ae5efd0c2acb627239dc9f207311337b6f2c12103fd95a9cb3098e6447d0de9f76cc97fd5e36830f9c7044457c15a0e81316bf28f2103fea219d4ccfd7641cebbb2439740bb4bd7c4730c1abd6ca1dc44386533816df957ae} Confirmations:3519 NextBlockHash:0x7f514b6d785b52adfeee56919d9deb12059516145aaae36f997cd79890c11bac}}
//--- PASS: TestRpcClient_GetBlockHeaderByHash (0.53s)
//PASS

func TestRpcClient_GetBlockHash(t *testing.T) {
	response := TestNetClient.GetBlockHash(3386365)
	log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t, "0x035212da3f0e73cd41e3f6e22ccbedaac064e4150ad6dd2bed3eeff420be3179", r)
}

//=== RUN   TestRpcClient_GetBlockHash
//2019/11/08 11:22:25 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:0x035212da3f0e73cd41e3f6e22ccbedaac064e4150ad6dd2bed3eeff420be3179}
//--- PASS: TestRpcClient_GetBlockHash (0.53s)
//PASS

// RpcSystemAssetTracker plugin required, better to use your private net
func TestRpcClient_GetClaimable(t *testing.T) {
	response := LocalClient.GetClaimable("APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR")
	//log.Printf("%+v", response)
	r := response.Result
	c := r.Claimables[0]
	assert.Equal(t, "bc0fda55480440dbd5492de670d0fc44ecd336a20d55caeb43eedbe239ea3c65", c.TxId)
	assert.Equal(t, 0, c.N)
	assert.Equal(t, "APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR", r.Address)
	assert.Equal(t, 4104, r.Unclaimed)
}

//=== RUN   TestRpcClient_GetClaimable
//2019/11/04 16:24:16 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Claimables:[{TxId:bc0fda55480440dbd5492de670d0fc44ecd336a20d55caeb43eedbe239ea3c65 N:0 Value:100000000 StartHeight:1870 EndHeight:2383 Generated:4104 SysFee:0 Unclaimed:4104}] Address:APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR Unclaimed:4104}}
//--- PASS: TestRpcClient_GetClaimable (0.33s)
//PASS

func TestRpcClient_GetConnectionCount(t *testing.T) {
	response := TestNetClient.GetConnectionCount()
	//log.Printf("%+v", response)
	assert.Equal(t, 48, response.Result)
}

//=== RUN   TestRpcClient_GetConnectionCount
//2019/11/04 16:25:18 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:48}
//--- PASS: TestRpcClient_GetConnectionCount (0.35s)
//PASS


func TestRpcClient_GetContractState(t *testing.T) {
	response := TestNetClient.GetContractState(Nep5ScriptHash)
	//log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t, 0, r.Version)
	assert.Equal(t, "0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263", r.Hash)
	assert.Equal(t, "QLC", r.Name)
}

//=== RUN   TestRpcClient_GetContractState
//2019/11/08 11:45:25 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Version:0 Hash:0xb9d7ea3062e6aeeb3e8ad9548220c4ba1361d263 Script:011ac56b6c766b00527ac46c766b51527ac4616168164e656f2e52756e74696d652e47657454726967676572009c6c766b54527ac46c766b54c3643e0061145c564ab204122ddce30eb9a6accbfa23b27cc3ac6168184e656f2e52756e74696d652e436865636b5769746e6573736c766b55527ac46259036168164e656f2e52756e74696d652e47657454726967676572609c6c766b56527ac46c766b56c364b702616c766b00c304696e6974876c766b57527ac46c766b57c364110061654e036c766b55527ac46206036c766b00c30a6d696e74546f6b656e73876c766b58527ac46c766b58c36411006165df066c766b55527ac462d8026c766b00c30b746f74616c537570706c79876c766b59527ac46c766b59c3641100616537096c766b55527ac462a9026c766b00c3046e616d65876c766b5a527ac46c766b5ac3641100616594026c766b55527ac46281026c766b00c30673796d626f6c876c766b5b527ac46c766b5bc364110061657d026c766b55527ac46257026c766b00c30a746f74616c546f6b656e876c766b5c527ac46c766b5cc3641100616562026c766b55527ac46229026c766b00c30869636f546f6b656e876c766b5d527ac46c766b5dc364110061658f036c766b55527ac462fd016c766b00c30669636f4e656f876c766b5e527ac46c766b5ec36411006165ca036c766b55527ac462d3016c766b00c306656e6449636f876c766b5f527ac46c766b5fc364110061650c046c766b55527ac462a9016c766b00c3087472616e73666572876c766b60527ac46c766b60c3647900616c766b51c3c0539c009c6c766b0114527ac46c766b0114c3640e00006c766b55527ac46264016c766b51c300c36c766b0111527ac46c766b51c351c36c766b0112527ac46c766b51c352c36c766b0113527ac46c766b0111c36c766b0112c36c766b0113c361527265f2076c766b55527ac46215016c766b00c30962616c616e63654f66876c766b0115527ac46c766b0115c3644d00616c766b51c3c0519c009c6c766b0117527ac46c766b0117c3640e00006c766b55527ac462cd006c766b51c300c36c766b0116527ac46c766b0116c36165cb096c766b55527ac462aa006c766b00c308646563696d616c73876c766b0118527ac46c766b0118c36411006165ad006c766b55527ac4627c00616165990b6c766b52527ac461650d0d6c766b53527ac46c766b53c300907c907ca1630e006c766b52c3c000a0620400006c766b0119527ac46c766b0119c3642f00616c766b52c36c766b53c3617c06726566756e6453c168124e656f2e52756e74696d652e4e6f746966796161006c766b55527ac46203006c766b55c3616c756600c56b0b516c696e6b20546f6b656e616c756600c56b03514c43616c756600c56b58616c756600c56b07008048efefd801616c756653c56b616168164e656f2e53746f726167652e476574436f6e746578740b746f74616c537570706c79617c680f4e656f2e53746f726167652e4765746c766b00527ac46c766b00c3c000a06c766b51527ac46c766b51c3640e00006c766b52527ac462df006168164e656f2e53746f726167652e476574436f6e74657874145c564ab204122ddce30eb9a6accbfa23b27cc3ac0800803dafbe50d300615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e746578740b746f74616c537570706c790800803dafbe50d300615272680f4e656f2e53746f726167652e5075746100145c564ab204122ddce30eb9a6accbfa23b27cc3ac0800803dafbe50d300615272087472616e7366657254c168124e656f2e52756e74696d652e4e6f7469667961516c766b52527ac46203006c766b52c3616c756652c56b616168164e656f2e53746f726167652e476574436f6e746578740b746f74616c537570706c79617c680f4e656f2e53746f726167652e4765746c766b00527ac46c766b00c30800803dafbe50d300946c766b51527ac46203006c766b51c3616c756652c56b616168164e656f2e53746f726167652e476574436f6e746578740b746f74616c537570706c79617c680f4e656f2e53746f726167652e4765746c766b00527ac46c766b00c30800803dafbe50d30094050008711b0c966c766b51527ac46203006c766b51c3616c756655c56b61145c564ab204122ddce30eb9a6accbfa23b27cc3ac6168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b52527ac46c766b52c3640e00006c766b53527ac4624e016168164e656f2e53746f726167652e476574436f6e746578740b746f74616c537570706c79617c680f4e656f2e53746f726167652e4765746c766b00527ac4080000869eae29d5006c766b00c3946c766b51527ac46c766b51c300a16c766b54527ac46c766b54c3640f0061006c766b53527ac462d7006168164e656f2e53746f726167652e476574436f6e746578740b746f74616c537570706c79080000869eae29d500615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e74657874145c564ab204122ddce30eb9a6accbfa23b27cc3ac6c766b51c3615272680f4e656f2e53746f726167652e5075746100145c564ab204122ddce30eb9a6accbfa23b27cc3ac6c766b51c3615272087472616e7366657254c168124e656f2e52756e74696d652e4e6f7469667961516c766b53527ac46203006c766b53c3616c75665cc56b61616520076c766b00527ac46c766b00c3c0009c6c766b58527ac46c766b58c3640f0061006c766b59527ac4624f026168184e656f2e426c6f636b636861696e2e4765744865696768746168184e656f2e426c6f636b636861696e2e4765744865616465726168174e656f2e4865616465722e47657454696d657374616d706c766b51527ac46c766b51c304d0013d5a946c766b52527ac46c766b52c36c766b00c3617c656e096c766b53527ac46c766b52c36165b2046c766b54527ac46c766b54c3009c6c766b5a527ac46c766b5ac3643900616c766b00c36c766b53c3617c06726566756e6453c168124e656f2e52756e74696d652e4e6f7469667961006c766b59527ac46274016c766b00c36c766b53c36c766b54c361527265b8046c766b55527ac46c766b55c3009c6c766b5b527ac46c766b5bc3640f0061006c766b59527ac46236016168164e656f2e53746f726167652e476574436f6e746578746c766b00c3617c680f4e656f2e53746f726167652e4765746c766b56527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b00c36c766b55c36c766b56c393615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e746578740b746f74616c537570706c79617c680f4e656f2e53746f726167652e4765746c766b57527ac46168164e656f2e53746f726167652e476574436f6e746578740b746f74616c537570706c796c766b55c36c766b57c393615272680f4e656f2e53746f726167652e50757461006c766b00c36c766b55c3615272087472616e7366657254c168124e656f2e52756e74696d652e4e6f7469667961516c766b59527ac46203006c766b59c3616c756651c56b616168164e656f2e53746f726167652e476574436f6e746578740b746f74616c537570706c79617c680f4e656f2e53746f726167652e4765746c766b00527ac46203006c766b00c3616c75665bc56b6c766b00527ac46c766b51527ac46c766b52527ac4616c766b52c300a16c766b55527ac46c766b55c3640e00006c766b56527ac46204026c766b00c36168184e656f2e52756e74696d652e436865636b5769746e657373009c6c766b57527ac46c766b57c3640e00006c766b56527ac462c8016c766b00c36c766b51c39c6c766b58527ac46c766b58c3640e00516c766b56527ac462a3016168164e656f2e53746f726167652e476574436f6e746578746c766b00c3617c680f4e656f2e53746f726167652e4765746c766b53527ac46c766b53c36c766b52c39f6c766b59527ac46c766b59c3640e00006c766b56527ac46246016c766b53c36c766b52c39c6c766b5a527ac46c766b5ac3643b006168164e656f2e53746f726167652e476574436f6e746578746c766b00c3617c68124e656f2e53746f726167652e44656c657465616241006168164e656f2e53746f726167652e476574436f6e746578746c766b00c36c766b53c36c766b52c394615272680f4e656f2e53746f726167652e507574616168164e656f2e53746f726167652e476574436f6e746578746c766b51c3617c680f4e656f2e53746f726167652e4765746c766b54527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b51c36c766b54c36c766b52c393615272680f4e656f2e53746f726167652e507574616c766b00c36c766b51c36c766b52c3615272087472616e7366657254c168124e656f2e52756e74696d652e4e6f7469667961516c766b56527ac46203006c766b56c3616c756652c56b6c766b00527ac4616168164e656f2e53746f726167652e476574436f6e746578746c766b00c3617c680f4e656f2e53746f726167652e4765746c766b51527ac46203006c766b51c3616c756654c56b6c766b00527ac4616c766b00c3009f6c766b51527ac46c766b51c3640f0061006c766b52527ac4623b006c766b00c30380de28a0009c6c766b53527ac46c766b53c364140061050008711b0c6c766b52527ac4620f0061006c766b52527ac46203006c766b52c3616c756659c56b6c766b00527ac46c766b51527ac46c766b52527ac4616c766b51c30400e1f505966c766b52c3956c766b53527ac46168164e656f2e53746f726167652e476574436f6e746578740b746f74616c537570706c79617c680f4e656f2e53746f726167652e4765746c766b54527ac4080000869eae29d5006c766b54c3946c766b55527ac46c766b55c300a16c766b56527ac46c766b56c3643900616c766b00c36c766b51c3617c06726566756e6453c168124e656f2e52756e74696d652e4e6f7469667961006c766b57527ac46276006c766b55c36c766b53c39f6c766b58527ac46c766b58c3644d00616c766b00c36c766b53c36c766b55c3946c766b52c3960400e1f50595617c06726566756e6453c168124e656f2e52756e74696d652e4e6f74696679616c766b55c36c766b53527ac4616c766b53c36c766b57527ac46203006c766b57c3616c756657c56b6161682953797374656d2e457865637574696f6e456e67696e652e476574536372697074436f6e7461696e65726c766b00527ac46c766b00c361681d4e656f2e5472616e73616374696f6e2e4765745265666572656e6365736c766b51527ac4616c766b51c36c766b52527ac4006c766b53527ac4629d006c766b52c36c766b53c3c36c766b54527ac4616c766b54c36168154e656f2e4f75747075742e47657441737365744964209b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc59c6c766b55527ac46c766b55c3642d006c766b54c36168184e656f2e4f75747075742e476574536372697074486173686c766b56527ac4622c00616c766b53c351936c766b53527ac46c766b53c36c766b52c3c09f635aff006c766b56527ac46203006c766b56c3616c756651c56b6161682d53797374656d2e457865637574696f6e456e67696e652e476574457865637574696e67536372697074486173686c766b00527ac46203006c766b00c3616c756658c56b6161682953797374656d2e457865637574696f6e456e67696e652e476574536372697074436f6e7461696e65726c766b00527ac46c766b00c361681a4e656f2e5472616e73616374696f6e2e4765744f7574707574736c766b51527ac4006c766b52527ac4616c766b51c36c766b53527ac4006c766b54527ac462cd006c766b53c36c766b54c3c36c766b55527ac4616c766b55c36168184e656f2e4f75747075742e47657453637269707448617368616505ff907c907c9e6345006c766b55c36168154e656f2e4f75747075742e47657441737365744964209b7cffdaa674beae0f930ebe6085af9093e5fe56b34a5c220ccdcf6efc336fc59c620400006c766b56527ac46c766b56c3642d00616c766b52c36c766b55c36168134e656f2e4f75747075742e47657456616c7565936c766b52527ac461616c766b54c351936c766b54527ac46c766b54c36c766b53c3c09f632aff6c766b52c36c766b57527ac46203006c766b57c3616c75665ac56b6c766b00527ac46c766b51527ac46161657cfe6c766b52527ac46c766b00c300a16311006c766b00c3026054a0009c620400006c766b53527ac46c766b53c364870161067072656669786c766b51c37e6c766b54527ac46168164e656f2e53746f726167652e476574436f6e746578746c766b54c3617c680f4e656f2e53746f726167652e4765746c766b55527ac404002f68596c766b55c3946c766b56527ac46c766b56c300a0009c6c766b57527ac46c766b57c3643900616c766b51c36c766b52c3617c06726566756e6453c168124e656f2e52756e74696d652e4e6f7469667961006c766b58527ac462e9006c766b56c36c766b52c39f6c766b59527ac46c766b59c3648100616168164e656f2e53746f726167652e476574436f6e746578746c766b54c36c766b56c36c766b55c393615272680f4e656f2e53746f726167652e507574616c766b51c36c766b52c36c766b56c394617c06726566756e6453c168124e656f2e52756e74696d652e4e6f74696679616c766b56c36c766b58527ac46251006168164e656f2e53746f726167652e476574436f6e746578746c766b54c36c766b52c36c766b55c393615272680f4e656f2e53746f726167652e50757461616c766b52c36c766b58527ac46203006c766b58c3616c7566 Parameters:[String Array] Returntype:ByteArray Name:QLC CodeVersion:1.0 Author:qlink Email:admin@qlink.mobi Description:qlink token		 Properties:{Storage:true DynamicInvoke:false}}}
//--- PASS: TestRpcClient_GetContractState (2.25s)
//PASS

// RpcNep5Tracker plugin required, better to use your private net
func TestRpcClient_GetNep5Balances(t *testing.T) {
	response := TestNetClient.GetNep5Balances(TestNetWalletAddress)
	//log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t, "b9d7ea3062e6aeeb3e8ad9548220c4ba1361d263", r.Balances[0].AssetHash)
	assert.Equal(t, "AUrE5r4NHznrgvqoFAGhoUbu96PE5YeDZY", r.Address)
}

//=== RUN   TestRpcClient_GetNep5Balances
//2019/11/08 14:08:37 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Balances:[{AssetHash:b9d7ea3062e6aeeb3e8ad9548220c4ba1361d263 Amount:0 LastUpdatedBlock:3275044}] Address:AUrE5r4NHznrgvqoFAGhoUbu96PE5YeDZY}}
//--- PASS: TestRpcClient_GetNep5Balances (3.52s)
//PASS

func TestRpcClient_GetNep5Transfers(t *testing.T) {
	response := TestNetClient.GetNep5Transfers(TestNetWalletAddress)
	//log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t, "AUrE5r4NHznrgvqoFAGhoUbu96PE5YeDZY", r.Address)
}

// need RpcWallet
func TestRpcClient_GetNewAddress(t *testing.T) {
	response := LocalClient.GetNewAddress()
	log.Printf("%+v", response)
}

//=== RUN   TestRpcClient_GetNewAddress
//2019/11/04 16:55:01 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:AVfppvZcLeynVzsCe5W52yH925rXZARWVd}
//--- PASS: TestRpcClient_GetNewAddress (0.64s)
//PASS

func TestRpcClient_GetPeers(t *testing.T) {
	response := TestNetClient.GetPeers()
	log.Printf("%+v", response)
}

//=== RUN   TestRpcClient_GetPeers
//2019/11/08 15:07:30 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Unconnected:[{Address:47.97.73.20 Port:20333} {Address:222.209.173.254 Port:20333} {Address:113.118.234.48 Port:20333} {Address:47.254.44.88 Port:20333} {Address:18.222.168.189 Port:10333} {Address:54.65.248.144 Port:20333} {Address:52.197.205.193 Port:20333} {Address:52.76.214.252 Port:20333} {Address:47.99.240.126 Port:20336} {Address:47.99.240.126 Port:20333} {Address:139.99.148.127 Port:20333} {Address:35.243.96.213 Port:20333} {Address:47.254.43.76 Port:20333} {Address:168.61.184.237 Port:20333} {Address:72.141.22.201 Port:20333} {Address:47.101.221.169 Port:20333} {Address:18.222.161.128 Port:10333} {Address:13.228.7.48 Port:20333} {Address:164.128.165.13 Port:20333} {Address:147.135.129.22 Port:20333} {Address:40.121.153.64 Port:20333} {Address:47.254.83.14 Port:20333} {Address:47.244.44.20 Port:20333} {Address:168.61.166.110 Port:20333} {Address:168.61.16.30 Port:20333} {Address:47.111.100.0 Port:20333} {Address:104.215.248.10 Port:20333} {Address:47.91.225.117 Port:20333} {Address:183.6.164.18 Port:20333} {Address:183.17.228.223 Port:20333} {Address:121.43.179.239 Port:20333} {Address:54.169.199.240 Port:20333} {Address:207.180.219.24 Port:20333} {Address:47.244.144.79 Port:20333} {Address:18.136.142.79 Port:30333} {Address:18.136.142.79 Port:20333} {Address:168.61.148.37 Port:20333}] Bad:[] Connected:[{Address:147.135.129.22 Port:20333} {Address:113.29.236.150 Port:20333} {Address:139.99.148.127 Port:20333} {Address:82.201.63.117 Port:20333} {Address:168.61.16.30 Port:20333} {Address:18.222.161.128 Port:10333} {Address:183.6.164.18 Port:20333} {Address:18.179.119.79 Port:20333} {Address:18.176.58.0 Port:20333} {Address:47.75.146.164 Port:20333} {Address:35.192.172.11 Port:0} {Address:40.121.153.64 Port:20333} {Address:104.196.172.132 Port:20333} {Address:47.75.217.176 Port:20333} {Address:142.44.138.161 Port:20333} {Address:47.101.221.169 Port:20333} {Address:18.191.171.240 Port:10333} {Address:172.255.99.84 Port:20333} {Address:183.6.164.18 Port:20333} {Address:18.136.142.79 Port:20333} {Address:207.180.219.24 Port:20333} {Address:47.99.240.126 Port:20333} {Address:47.99.240.126 Port:20336} {Address:47.111.100.0 Port:20333} {Address:47.244.144.79 Port:20333} {Address:47.254.44.88 Port:20333} {Address:52.130.66.169 Port:20333} {Address:104.215.248.10 Port:20333} {Address:54.238.172.91 Port:20333} {Address:5.35.241.70 Port:20333} {Address:119.139.196.60 Port:20333} {Address:121.43.179.239 Port:20333} {Address:149.129.175.157 Port:20333} {Address:47.52.159.228 Port:20333} {Address:54.65.248.144 Port:20333} {Address:47.244.44.20 Port:20333} {Address:172.105.228.12 Port:20333} {Address:52.76.214.252 Port:20333} {Address:47.254.43.76 Port:20333} {Address:13.76.173.63 Port:20333} {Address:172.105.199.143 Port:20333} {Address:47.97.73.20 Port:20333} {Address:164.128.165.13 Port:20333} {Address:104.248.132.109 Port:20333} {Address:139.217.114.241 Port:20333} {Address:47.90.28.83 Port:20333} {Address:18.136.142.79 Port:30333} {Address:13.58.169.218 Port:10333}]}}
//--- PASS: TestRpcClient_GetPeers (0.53s)
//PASS

func TestRpcClient_GetRawMemPool(t *testing.T) {
	response := TestNetClient.GetRawMemPool()
	log.Printf("%+v", response)
}

//=== RUN   TestRpcClient_GetRawMemPool
//2019/11/08 15:22:47 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:[]}
//--- PASS: TestRpcClient_GetRawMemPool (0.53s)
//PASS

func TestRpcClient_GetRawTransaction(t *testing.T) {
	response := TestNetClient.GetRawTransaction("0xca159430e3d72227c06a3880244111aea0368ecc09fa8c2eade001a1bbcc7d4a")
	//log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t, "0xca159430e3d72227c06a3880244111aea0368ecc09fa8c2eade001a1bbcc7d4a", r.Txid)
	assert.Equal(t, 242, r.Size)
	assert.Equal(t, "InvocationTransaction", r.Type)
	assert.Equal(t, 1, r.Version)
}

//=== RUN   TestRpcClient_GetRawTransaction
//2019/11/08 15:23:46 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Txid:0xca159430e3d72227c06a3880244111aea0368ecc09fa8c2eade001a1bbcc7d4a Size:242 Type:InvocationTransaction Version:1 Attributes:[{Usage:Script Data:5c564ab204122ddce30eb9a6accbfa23b27cc3ac} {Usage:Remark Data:313537313231383636323935373964363035643631}] Vin:[] Vout:[] SysFee:0 NetFee:0 Scripts:[{Invocation:4002a84056e9bf04ed47a6307c3030ac92704cb71a8c2fd46f45593c8ce57a403de47e19a4171114e7ec881d9f45d7851712e8eb922d11ce3a0de5ea64b8310025 Verification:2103f19ffa8acecb480ab727b0bf9ee934162f6e2a4308b59c80b732529ebce6f53dac}] Nonce:0 BlockHash:0x1a1d7b2f6d54e7c9084353372dd526301a456900827ce8478fcff1a7a00766f7 Confirmations:115691 Blocktime:1571218675 Script:0600203d88792d148f6c5be89c0cb6579e44a8bf9bfd2ecbcc11dfdc145c564ab204122ddce30eb9a6accbfa23b27cc3ac53c1087472616e736665726763d26113bac4208254d98a3eebaee66230ead7b9 Gas:0 Claims:[]}}
//--- PASS: TestRpcClient_GetRawTransaction (0.77s)
//PASS

func TestRpcClient_GetStorage(t *testing.T) {
	response := TestNetClient.GetStorage("0x2fabf24313d69629fa56c51716f94cd5cbd36b88", "636f6e747261637400746f74616c537570706c79") // key = "contract" + "00" + "totalSupply"
	//log.Printf("%+v", response)
	assert.Equal(t, "00e1f505", response.Result)
}

//=== RUN   TestRpcClient_GetStorage
//2019/11/08 15:37:17 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:00e1f505}
//--- PASS: TestRpcClient_GetStorage (0.85s)
//PASS

func TestRpcClient_GetTransactionHeight(t *testing.T) {
	response := TestNetClient.GetTransactionHeight("0xca159430e3d72227c06a3880244111aea0368ecc09fa8c2eade001a1bbcc7d4a")
	//log.Printf("%+v", response)
	assert.Equal(t, 3275044, response.Result)
}

//=== RUN   TestRpcClient_GetTransactionHeight
//2019/11/08 15:44:53 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:3275044}
//--- PASS: TestRpcClient_GetTransactionHeight (0.68s)
//PASS

func TestRpcClient_GetTxOut(t *testing.T) {
	response := TestNetClient.GetTxOut("0x300d3083620e31708e291dd0732fa0117c8e893b532ef3adc9e0f47b17e29254", 0)
	//log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t, 0, r.N)
	assert.Equal(t, "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b", r.Asset)
	assert.Equal(t, 1000000, r.Value)
	assert.Equal(t, "AazAnhssUfNyBC3rdBKseGuck7voaF5p68", r.Address)
}

//=== RUN   TestRpcClient_GetTxOut
//2019/11/08 15:49:13 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{N:0 Asset:0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b Value:1000000 Address:AazAnhssUfNyBC3rdBKseGuck7voaF5p68}}
//--- PASS: TestRpcClient_GetTxOut (3.35s)
//PASS

// need RpcSystemAssetTracker
func TestRpcClient_GetUnclaimed(t *testing.T) {
	response := TestNetClient.GetUnclaimed(TestNetWalletAddress)
	log.Printf("%+v", response)
}

//=== RUN   TestRpcClient_GetUnclaimed
//2019/11/08 15:57:44 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Available:0 Unavailable:0 Unclaimed:0}}
//--- PASS: TestRpcClient_GetUnclaimed (3.97s)
//PASS

// Use your private net
func TestRpcClient_GetUnclaimedGas(t *testing.T) {
	response := LocalClient.GetUnclaimedGas()
	//log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t, 4104, r.Available)
	assert.Equal(t, 4096, r.Unavailable)
}

//=== RUN   TestRpcClient_GetUnclaimedGas
//2019/11/04 17:48:41 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Available:4104 Unavailable:4096}}
//--- PASS: TestRpcClient_GetUnclaimedGas (0.37s)
//PASS

// need RpcSystemAssetTracker, Use your private net
func TestRpcClient_GetUnspents(t *testing.T) {
	response := LocalClient.GetUnspents(LocalWalletAddress)
	log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t, "602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7", r.Balance[0].AssetHash)
	assert.Equal(t, "GAS", r.Balance[0].Asset)
	assert.Equal(t, "c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b", r.Balance[1].AssetHash)
	assert.Equal(t, "NEO", r.Balance[1].Asset)
}

//=== RUN   TestRpcClient_GetUnspents
//2019/11/04 17:51:04 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Balance:[{Unspent:[{Txid:883524d20d2f43789453a2e3b4d4cb7726bbceb71b6fffadc4916654da9ad90a N:0 Value:6996} {Txid:db961a3cd480934dc31c83f7665f859310d7b2c51d0a441df7c3ae2a87e14a28 N:0 Value:1120} {Txid:520bfe6fea7c4dd31f3cc3f3b5b9098dbc6d7b671575c7c2c28da53b4b79167f N:0 Value:81.97877} {Txid:9251ea1b8eb58ab908a33523e6712f63b1d78386b3a05011914d23feda973f87 N:0 Value:6762}] AssetHash:602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7 Asset:GAS AssetSymbol:GAS Amount:0} {Unspent:[{Txid:011863ed7e21e73eba017e1887ae3085771661ae624d631190e4457ef4202ad1 N:0 Value:1e+08}] AssetHash:c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b Asset:NEO AssetSymbol:NEO Amount:100000000}] Address:APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR}}
//--- PASS: TestRpcClient_GetUnspents (0.35s)
//PASS

func TestRpcClient_GetValidators(t *testing.T) {
	response := TestNetClient.GetValidators()
	log.Printf("%+v", response)
}

//=== RUN   TestRpcClient_GetValidators
//2019/11/08 16:26:07 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:[{PublicKey:02170376729a7c4c9d12c14d862fe6596a4eecebf4178a4d432b725d91b267f8e8 Votes:0 Active:false} {PublicKey:023cfe95a0301acfa831edc8ba10c55d9f3a985e43e1349dd79a2521629987ecf9 Votes:0 Active:false} {PublicKey:02494f3ff953e45ca4254375187004f17293f90a1aa4b1a89bc07065bc1da521f6 Votes:0 Active:false} {PublicKey:025bdf3f181f53e9696227843950deb72dcd374ded17c057159513c3d0abe20b64 Votes:90101602 Active:true} {PublicKey:0266b588e350ab63b850e55dbfed0feeda44410a30966341b371014b803a15af07 Votes:90101602 Active:true} {PublicKey:02e96b953cfc6b88c7490794dfa5b790c53be54154abb06a537338185c84843b3e Votes:0 Active:false} {PublicKey:02f9f6d63fd321879d8b387131b7b46c3a34fe090f4eac44276245ca86b6c69370 Votes:0 Active:false} {PublicKey:02ff8ac54687f36bbc31a91b730cc385da8af0b581f2d59d82b5cfef824fd271f6 Votes:0 Active:false} {PublicKey:03028007d683ceb4dc9084300d0cf16fe6d47a726e586bf3d63559cec133055652 Votes:90101602 Active:true} {PublicKey:030ef96257401b803da5dd201233e2be828795672b775dd674d69df83f7aec1e36 Votes:90100000 Active:true} {PublicKey:031c316e37d57006db9ac637c44169f0067fad4133f473fbd95d2d64605e7193af Votes:0 Active:false} {PublicKey:0344925b6126c8ae58a078b5b2ce98de8fff15a22dac6f57ffd3b108c72a0670d1 Votes:0 Active:false} {PublicKey:035e700a50b082b6c6986853bcd13340b3f636ae1ae39a5c201c0b0ce543569b6b Votes:0 Active:false} {PublicKey:036c27adb03c6da87236ecd1dc48d32e2626da084abdd72f759798e177b85d8dd6 Votes:0 Active:false} {PublicKey:036c5a11d322219d3386a37bdbf93fcf749f62c72e5898079bc73ff78ef6c1d2cf Votes:0 Active:false} {PublicKey:039271a57098559112648661f91c966a8d05775dff649d47a43817ba5de769c874 Votes:0 Active:false} {PublicKey:03b1521081dbf29df4828df3c12616bfb3a5600f4381dcb8ad92f99ba1e9a6c06a Votes:0 Active:false} {PublicKey:03b8bfe058dc404a2f9510606aee0de69e3d8b47e25d7b3af670577373640d51dc Votes:1000 Active:false} {PublicKey:03c089d7122b840a4935234e82e26ae5efd0c2acb627239dc9f207311337b6f2c1 Votes:90100000 Active:true} {PublicKey:03cdabe37ad8f2269ad39de0e8ee395358e38a2f654ba4918de8a5c2de80190ff9 Votes:0 Active:false} {PublicKey:03d3a870ee14a6f8f772b430f09dd65f5328e887e0d08d4e280730d20e2586198c Votes:0 Active:false} {PublicKey:03e88a1cf8c37b0b76858224905e126005a560487894309f2a3d69f7067df7d415 Votes:0 Active:false} {PublicKey:03fd95a9cb3098e6447d0de9f76cc97fd5e36830f9c7044457c15a0e81316bf28f Votes:90100000 Active:true} {PublicKey:03fea219d4ccfd7641cebbb2439740bb4bd7c4730c1abd6ca1dc44386533816df9 Votes:90100000 Active:true}]}
//--- PASS: TestRpcClient_GetValidators (8.40s)
//PASS

func TestRpcClient_GetVersion(t *testing.T) {
	response := TestNetClient.GetVersion()
	//log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t, 20333, r.Port)
	assert.Equal(t, 1109366691, r.Nonce)
	assert.Equal(t, "/Neo:2.10.3/", r.Useragent)
}

//=== RUN   TestRpcClient_GetVersion
//2019/11/08 16:27:26 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Port:20333 Nonce:1109366691 Useragent:/Neo:2.10.3/}}
//--- PASS: TestRpcClient_GetVersion (0.50s)
//PASS

// need RpcWallet, use your onw private net
func TestRpcClient_GetWalletHeight(t *testing.T) {
	response := LocalClient.GetWalletHeight()
	log.Printf("%+v", response)
	assert.Equal(t, 2894, response.Result)
}

//=== RUN   TestRpcClient_GetWalletHeight
//2019/11/04 17:58:54 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:2894}
//--- PASS: TestRpcClient_GetWalletHeight (0.33s)
//PASS

// TODO
func TestRpcClient_ImportPrivKey(t *testing.T) {

}

func TestRpcClient_InvokeFunction(t *testing.T) {
	response := TestNetClient.InvokeFunction(Nep5ScriptHash, "name")
	//log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t, "00c1046e616d656763d26113bac4208254d98a3eebaee66230ead7b9", r.Script)
	assert.Equal(t, "HALT", r.State)
	assert.Equal(t, "516c696e6b20546f6b656e", r.Stack[0].Value)
}

//=== RUN   TestRpcClient_InvokeFunction
//2019/11/08 16:34:53 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Script:00c1046e616d656763d26113bac4208254d98a3eebaee66230ead7b9 State:HALT GasConsumed:0.126 Stack:[{Type:ByteArray Value:516c696e6b20546f6b656e}]}}
//--- PASS: TestRpcClient_InvokeFunction (0.55s)
//PASS

func TestRpcClient_InvokeScript(t *testing.T) {
	response := TestNetClient.InvokeScript("00c1046e616d656763d26113bac4208254d98a3eebaee66230ead7b9")
	//log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t, "00c1046e616d656763d26113bac4208254d98a3eebaee66230ead7b9", r.Script)
	assert.Equal(t, "HALT", r.State)
	assert.Equal(t, "516c696e6b20546f6b656e", r.Stack[0].Value)
}

//=== RUN   TestRpcClient_InvokeScript
//2019/11/08 16:41:13 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Script:00c1046e616d656763d26113bac4208254d98a3eebaee66230ead7b9 State:HALT GasConsumed:0.126 Stack:[{Type:ByteArray Value:516c696e6b20546f6b656e}] Tx:}}
//--- PASS: TestRpcClient_InvokeScript (0.51s)
//PASS

// need RpcWallet, use your own private net
func TestRpcClient_ListAddress(t *testing.T) {
	response := LocalClient.ListAddress()
	log.Printf("%+v", response)
}

//=== RUN   TestRpcClient_ListAddress
//2019/11/04 18:21:27 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:[{Address:APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR HasKey:true Label: WatchOnly:false}]}
//--- PASS: TestRpcClient_ListAddress (0.35s)
//PASS


func TestRpcClient_ListPlugins(t *testing.T) {
	response := TestNetClient.ListPlugins()
	log.Printf("%+v", response)
	r := response.Result
	assert.Equal(t, "ApplicationLogs", r[0].Name)
}

//=== RUN   TestRpcClient_ListPlugins
//2019/11/08 16:43:25 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:[{Name:ApplicationLogs Version:2.10.3.0 Interfaces:[IRpcPlugin IPersistencePlugin]} {Name:ImportBlocks Version:2.10.3.0 Interfaces:[]} {Name:RpcNep5Tracker Version:2.10.3.0 Interfaces:[IPersistencePlugin IRpcPlugin]} {Name:RpcSecurity Version:2.10.3.0 Interfaces:[IRpcPlugin]} {Name:RpcSystemAssetTrackerPlugin Version:2.10.3.0 Interfaces:[IPersistencePlugin IRpcPlugin]} {Name:RpcWallet Version:2.10.3.0 Interfaces:[IRpcPlugin]} {Name:SimplePolicyPlugin Version:2.10.3.0 Interfaces:[ILogPlugin IPolicyPlugin]}]}
//--- PASS: TestRpcClient_ListPlugins (3.99s)
//PASS

// use your own private net
func TestRpcClient_SendFrom(t *testing.T) {
	response := LocalClient.SendFrom(AssetIdNeo, LocalWalletAddress, LocalWalletAddress, 100000000, 0, LocalWalletAddress)
	log.Printf("%+v", response)
}

//=== RUN   TestRpcClient_SendFrom
//2019/11/04 18:44:09 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Txid:0x5bfbda6774a96430c7a67e590c37c00b4c167c22268878d49a46849b15fa55b0 Size:202 Type:ContractTransaction Version:0 Attributes:[] Vin:[{Txid:0xef69ebd50ca876857e9f44c9901460d81dbfb5b8eaab2b7975325b82990bf3c2 Vout:0}] Vout:[{N:0 Asset:0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b Value:100000000 Address:APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR}] SysFee:0 NetFee:0 Scripts:[{Invocation:40bcb785e9c3a8c9df5dd77f5d9007607ebbf45ed52eaec63a6cf366432e31e6ef40a559ba4c8594ec8d2bbd8f5a5e194cee92017868b3ee1e0d5a5cf15be68b12 Verification:2103dc623d556a79437ffefa0133516ec91545f363b772149901ac589b5e927ad55fac}] Nonce:0 BlockHash: Confirmations:0 Blocktime:0 Script: Gas: Claims:[]}}
//--- PASS: TestRpcClient_SendFrom (0.34s)
//PASS

// TODO
// use your own private net
func TestRpcClient_SendRawTransaction(t *testing.T) {

}

// use your own private net
func TestRpcClient_SendToAddress(t *testing.T) {
	response := LocalClient.SendToAddress(AssetIdNeo, LocalWalletAddress, 100000000, 0, LocalWalletAddress)
	log.Printf("%+v", response)
}

//=== RUN   TestRpcClient_SendToAddress
//2019/11/04 18:44:34 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Txid:0x21bf1497ff54aef61e389d2180f5b4174963e52eaa1c5e717da14ba977c9e7f9 Size:202 Type:ContractTransaction Version:0 Attributes:[] Vin:[{Txid:0x5bfbda6774a96430c7a67e590c37c00b4c167c22268878d49a46849b15fa55b0 Vout:0}] Vout:[{N:0 Asset:0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b Value:100000000 Address:APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR}] SysFee:0 NetFee:0 Scripts:[{Invocation:40420f043d3c4d46b0979c3243329bcf9e3870eab68ef827243db81e0b04f62aa0a081efbee5d1368bf9e4d97b3edc8e7ba700a94d1799a4933a47e8947935b200 Verification:2103dc623d556a79437ffefa0133516ec91545f363b772149901ac589b5e927ad55fac}] Nonce:0 BlockHash: Confirmations:0 Blocktime:0 Script: Gas: Claims:[]}}
//--- PASS: TestRpcClient_SendToAddress (0.34s)
//PASS

func TestRpcClient_SubmitBlock(t *testing.T) {

}

func TestRpcClient_ValidateAddress(t *testing.T) {
	response := TestNetClient.ValidateAddress(TestNetWalletAddress)
	//log.Printf("%+v", response)
	assert.Equal(t, true, response.Result.IsValid)
}

//=== RUN   TestRpcClient_ValidateAddress
//2019/11/04 18:46:00 {RpcResponse:{JsonRpc:2.0 ID:1} ErrorResponse:{Error:{Code:0 Message:}} Result:{Address:APPmjituYcgfNxjuQDy9vP73R2PmhFsYJR IsValid:true}}
//--- PASS: TestRpcClient_ValidateAddress (0.35s)
//PASS
