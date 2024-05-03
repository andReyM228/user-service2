package car_tx

import (
	"context"
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

func (r Repository) Get(ctx context.Context, txHash string) (domain.CarTx, error) {
	var transaction CarTxDB

	if err := r.db.Get(&transaction, "SELECT * FROM car_transactions WHERE tx_hash = $1", txHash); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.log.Info(err.Error())
			return domain.CarTx{}, errs.NotFoundError{What: "car"}
		}

		r.log.Error(err.Error())
		return domain.CarTx{}, errs.InternalError{Cause: err.Error()}
	}

	return transaction.toDomain(), nil
}

func (r Repository) GetAll(ctx context.Context, kind string) (domain.CarTxs, error) {
	var transactions []CarTxDB

	if err := r.db.Select(&transactions, "SELECT * FROM car_transactions WHERE kind = $1", kind); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.log.Info(err.Error())
			return domain.CarTxs{}, errs.NotFoundError{What: "cars"}
		}

		r.log.Error(err.Error())
		return domain.CarTxs{}, errs.InternalError{Cause: err.Error()}
	}

	return toDomainList(transactions), nil
}

func (r Repository) Create(ctx context.Context, transaction domain.CarTx) error {
	query := `
		INSERT INTO car_transactions (
			tx_hash,
		    status,
		    kind,
		    error
		) VALUES (
		    :tx_hash,
		    :status,
		    :kind,
		    :error,
		)
	`

	_, err := r.db.NamedExecContext(ctx, query, fromDomain(transaction))
	if err != nil {
		return errs.InternalError{Cause: err.Error()}
	}

	return nil
}

// TODO: сделать так во всех репах 3-х сервисов

func (r Repository) Update(ctx context.Context, transaction domain.CarTx) error {
	query := `
		UPDATE car_transactions SET 
			tx_hash = :tx_hash,
		    status = :status,
		    kind = :kind,
		    error = :error
		WHERE 
			id = :id
	`

	_, err := r.db.NamedExecContext(ctx, query, fromDomain(transaction))
	if err != nil {
		return errs.InternalError{Cause: err.Error()}
	}

	return nil
}

func (r Repository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec("DELETE FROM car_transactions WHERE id = $1", id)
	if err != nil {
		r.log.Error(err.Error())
		return errs.InternalError{Cause: err.Error()}
	}

	return nil
}
