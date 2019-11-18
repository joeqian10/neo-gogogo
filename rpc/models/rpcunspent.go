package models

type RpcUnspent struct {
	Balance []UnspentBalance `json:"balance"`
	Address string `json:"address"`
}

type UnspentBalance struct {
	Unspent     []Unspent `json:"unspent"`
	AssetHash   string    `json:"asset_hash"`
	Asset       string    `json:"asset"`
	AssetSymbol string    `json:"asset_symbol"`
	Amount      float64   `json:"amount"`
}

type UnspentSlice []Unspent

func (us UnspentSlice) Len() int {
	return len(us)
}

func (us UnspentSlice) Less(i, j int) bool {
	return us[i].Value < us[j].Value
}

func (us UnspentSlice) Swap(i, j int) {
	t := us[i]
	us[i] = us[j]
	us[j] = t
}

func (us UnspentSlice) Sum() float64 {
	var s float64 = 0
	for _, u := range us {
		s += u.Value
	}
	return s
}

type Unspent struct {
	Txid  string  `json:"txid"`
	N     int     `json:"n"`
	Value float64 `json:"value"`
}

