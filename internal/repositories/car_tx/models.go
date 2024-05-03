package car_tx

import (
	"database/sql"
	"time"
	"user_service/internal/domain"
)

type CarTxDB struct {
	ID        int64
	TxHash    string
	Kind      string
	Status    string
	Error     sql.NullString
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func fromDomain(carTx domain.CarTx) CarTxDB {
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

func (c CarTxDB) toDomain() domain.CarTx {
	return domain.CarTx{
		ID:        c.ID,
		TxHash:    c.TxHash,
		Kind:      domain.Kind(c.Kind),
		Status:    domain.Status(c.Status),
		Error:     c.Error.String,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func toDomainList(carTxsDB []CarTxDB) domain.CarTxs {
	result := make([]domain.CarTx, 0, len(carTxsDB))

	for _, tx := range carTxsDB {
		result = append(result, tx.toDomain())
	}

	return domain.CarTxs{
		Transactions: result,
	}
}
