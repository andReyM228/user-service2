package car_txs

import (
	"time"
)

type CarTx struct {
	ID        int64
	TxHash    string
	Kind      Kind
	Status    Status
	Error     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CarTxs struct {
	Transactions []CarTx
}
