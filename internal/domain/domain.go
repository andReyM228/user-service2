package domain

import (
	"time"
)

const (
	FieldID        = "id"
	FieldName      = "name"
	FieldSurname   = "surname"
	FieldPhone     = "phone"
	FieldEmail     = "email"
	FieldIsAdmin   = "is_admin"
	FieldPassword  = "password"
	FieldChatID    = "chat_id"
	FieldCreatedAt = "created_at"
)

type User struct {
	ID             int
	Name           string
	Surname        string
	Phone          string
	Email          string
	AccountAddress string `db:"account_address"`
	IsAdmin        bool   `db:"is_admin"`
	Password       string
	ChatID         int64 `db:"chat_id"`
	Cars           []Car
	CreatedAt      time.Time `db:"created_at"`
}

type Car struct {
	ID        int
	Name      string
	Model     string
	Price     int64
	Image     string
	CreatedAt time.Time `db:"created_at"`
}

type Cars struct {
	Cars []Car
}

type UserCar struct {
	ID        int64
	UserID    int64
	CarID     int64
	CreatedAt time.Time `db:"created_at"`
}

type UserCars struct {
	Cars []UserCar
}
