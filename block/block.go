package block

import "github.com/joeqian10/neo-gogogo/tx"

type Block struct {
	BlockHeader
	Tx []*tx.Transaction
}
