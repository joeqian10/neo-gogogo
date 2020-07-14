package block

import (
	//"github.com/joeqian10/neo-gogogo/rpc/models"
	"github.com/joeqian10/neo-gogogo/tx"
)

type Block struct {
	BlockHeader
	Tx []*tx.Transaction
}

//func NewBlockFromRPC(rpcBlock *models.RpcBlock) (*Block, error) {
//	var block = new &Block{
//		BlockHeader: NewBlockHeaderFromRPC(&rpcBlock.RpcBlockHeader),
//		Tx:          nil,
//	}
//}
