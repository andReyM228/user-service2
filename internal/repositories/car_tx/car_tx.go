package car_tx

import (
	"database/sql"
	"errors"
	"github.com/andReyM228/lib/errs"
	"github.com/andReyM228/lib/log"
	"github.com/jmoiron/sqlx"
	"user_service/internal/domain"
)

type Repository struct {
	db  *sqlx.DB
	log log.Logger
}

func NewRepository(database *sqlx.DB, log log.Logger) Repository {
	return Repository{
		db:  database,
		log: log,
	}
}

func (r Repository) Get(id int64) (domain.CarTx, error) {
	var carTx domain.CarTx

	if err := r.db.Get(&carTx, "SELECT * FROM car_transactions WHERE id = $1", id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.log.Info(err.Error())
			return domain.CarTx{}, errs.NotFoundError{What: "car"}
		}

		r.log.Error(err.Error())
		return domain.CarTx{}, errs.InternalError{Cause: err.Error()}
	}

	return carTx, nil
}

func (r Repository) GetAll(kind string) (domain.CarTxs, error) {
	var cars domain.CarTxs

	if err := r.db.Select(&cars, "SELECT * FROM car_transactions WHERE kind = $1", kind); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.log.Info(err.Error())
			return domain.CarTxs{}, errs.NotFoundError{What: "cars"}
		}

		r.log.Error(err.Error())
		return domain.CarTxs{}, errs.InternalError{Cause: err.Error()}
	}

	return cars, nil
}

func (r Repository) Create(transaction domain.CarTx) error {
	if _, err := r.db.Exec("INSERT INTO car_transactions (tx_hash, kind) VALUES ($1, $2)", transaction.TxHash, transaction.Kind); err != nil {
		r.log.Error(err.Error())
		return errs.InternalError{Cause: err.Error()}
	}

	return nil
}

func (r Repository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM car_transactions WHERE id = $1", id)
	if err != nil {
		r.log.Error(err.Error())
		return errs.InternalError{Cause: err.Error()}
	}

	return nil
}
