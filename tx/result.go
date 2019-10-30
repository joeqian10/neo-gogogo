package transaction

import (
	"github.com/joeqian10/neo-gogogo/helper"
)

// Result represents the Result of a transaction.
type Result struct {
	AssetID helper.UInt256
	Amount  helper.Fixed8
}
