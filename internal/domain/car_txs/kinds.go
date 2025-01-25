package car_txs

const (
	KindBuy  = "buy"
	KindSell = "sell"
)

type Kind string

var Kinds = map[Kind]struct{}{
	KindBuy:  {},
	KindSell: {},
}

func (k Kind) String() string {
	return string(k)
}
