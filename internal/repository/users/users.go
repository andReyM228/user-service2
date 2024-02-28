package users

import (
	"database/sql"
	"errors"
	"fmt"

	"user_service/internal/domain"
	"user_service/internal/repository"

	"github.com/andReyM228/lib/log"
	"github.com/jmoiron/sqlx"
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

// TODO: обработка ошибок

func (r Repository) Get(field string, value any) (domain.User, error) {
	var user domain.User
	var cars []domain.Car

	if err := r.db.Get(&user, fmt.Sprintf("SELECT * FROM users WHERE %s = %v", field, value)); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.log.Info(err.Error())
			return domain.User{}, repository.NotFound{NotFound: "user"}
		}

		r.log.Error(err.Error())
		return domain.User{}, repository.InternalServerError{}
	}

	if err := r.db.Select(&cars, `
		SELECT cars.*
		FROM cars
		JOIN user_cars ON user_cars.car_id = cars.id
		JOIN users ON users.id = user_cars.user_id
		WHERE users.id = $1;
		`, user.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.log.Info(err.Error())
			return domain.User{}, repository.NotFound{NotFound: "users cars"}
		}

		r.log.Error(err.Error())
		return domain.User{}, repository.InternalServerError{}
	}

	user.Cars = cars
	return user, nil
}

func (r Repository) Update(user domain.User) error {
	if _, err := r.db.Exec("UPDATE users SET name = $1, surname = $2, phone = $3, email = $4, password = $5, chat_id = $6, account_address = $7 WHERE id = $5",
		user.Name, user.Surname, user.Phone, user.Email, user.ID, user.Password, user.ChatID, user.AccountAddress); err != nil {
		r.log.Error(err.Error())
		return repository.InternalServerError{}
	}

	return nil
}

func (r Repository) Create(user domain.User) error {
	if _, err := r.db.Exec("INSERT INTO users (name, surname, phone, email, password, chat_id, account_address) VALUES ($1, $2, $3, $4, $5, $6, $7)", user.Name, user.Surname, user.Phone, user.Email, user.Password, user.ChatID, user.AccountAddress); err != nil {
		r.log.Error(err.Error())
		return repository.InternalServerError{}
	}

	return nil
}

func (r Repository) Delete(id int64) error {
	if _, err := r.db.Exec("DELETE FROM users WHERE id = $1", id); err != nil {
		r.log.Error(err.Error())
		return repository.InternalServerError{}
	}

	return nil
}
