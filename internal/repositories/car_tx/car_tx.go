package car_tx

import (
	"context"
	"errors"
	"github.com/andReyM228/lib/errs"
	"github.com/andReyM228/lib/log"
	"gorm.io/gorm"
	"user_service/internal/domain/car_txs"
	"user_service/internal/repositories"
)

var _ repositories.CarTx = Repository{}

type Repository struct {
	db  *gorm.DB
	log log.Logger
}

func NewRepository(database *gorm.DB, log log.Logger) Repository {
	return Repository{
		db:  database,
		log: log,
	}
}

func (r Repository) Get(ctx context.Context, txHash string) (car_txs.CarTx, error) {
	var transaction car_txs.CarTx

	query := r.db.WithContext(ctx).Where("tx_hash = ?", txHash)

	if err := query.First(&transaction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.log.Info(err.Error())
			return car_txs.CarTx{}, errs.NotFoundError{What: "car transaction"}
		}

		r.log.Error(err.Error())
		return car_txs.CarTx{}, errs.InternalError{Cause: err.Error()}
	}

	return transaction, nil
}

func (r Repository) GetAll(ctx context.Context, kind string) (car_txs.CarTxs, error) {
	var transactions car_txs.CarTxs

	query := r.db.WithContext(ctx).Where("kind = ?", kind)

	if err := query.Find(&transactions).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.log.Info(err.Error())
			return car_txs.CarTxs{}, errs.NotFoundError{What: "car transactions"}
		}

		r.log.Error(err.Error())
		return car_txs.CarTxs{}, errs.InternalError{Cause: err.Error()}
	}

	return transactions, nil
}

func (r Repository) Create(ctx context.Context, transaction car_txs.CarTx) (car_txs.CarTx, error) {
	if err := r.db.WithContext(ctx).Create(&transaction).Error; err != nil {
		r.log.Error(err.Error())
		return car_txs.CarTx{}, errs.InternalError{Cause: err.Error()}
	}

	return transaction, nil
}

func (r Repository) Update(ctx context.Context, transaction car_txs.CarTx) error {
	query := r.db.WithContext(ctx).Model(&car_txs.CarTx{}).Where("id = ?", transaction.ID)

	if err := query.Updates(transaction).Error; err != nil {
		r.log.Error(err.Error())
		return errs.InternalError{Cause: err.Error()}
	}

	return nil
}

func (r Repository) Delete(ctx context.Context, id int64) error {
	query := r.db.WithContext(ctx).Where("id = ?", id)

	if err := query.Delete(&car_txs.CarTx{}).Error; err != nil {
		r.log.Error(err.Error())
		return errs.InternalError{Cause: err.Error()}
	}

	return nil
}
