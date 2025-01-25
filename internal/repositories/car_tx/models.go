package car_tx

import (
	"database/sql"
	"time"
	"user_service/internal/domain/car_txs"
)

type CarTxDB struct {
	ID        int64          `db:"id"`
	TxHash    string         `db:"tx_hash"`
	Kind      string         `db:"kind"`
	Status    string         `db:"status"`
	Error     sql.NullString `db:"error"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}

func fromDomain(carTx car_txs.CarTx) CarTxDB {
	return CarTxDB{
		ID:     carTx.ID,
		TxHash: carTx.TxHash,
		Kind:   carTx.Kind.String(),
		Status: carTx.Status.String(),
		Error: sql.NullString{
			String: carTx.Error,
			Valid:  carTx.Error != "",
		},
		CreatedAt: carTx.CreatedAt,
		UpdatedAt: carTx.UpdatedAt,
	}
}

func (c CarTxDB) toDomain() car_txs.CarTx {
	return car_txs.CarTx{
		ID:        c.ID,
		TxHash:    c.TxHash,
		Kind:      car_txs.Kind(c.Kind),
		Status:    car_txs.Status(c.Status),
		Error:     c.Error.String,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func toDomainList(carTxsDB []CarTxDB) car_txs.CarTxs {
	result := make([]car_txs.CarTx, 0, len(carTxsDB))

	for _, tx := range carTxsDB {
		result = append(result, tx.toDomain())
	}

	return car_txs.CarTxs{
		Transactions: result,
	}
}
