package users

import (
	"time"
	"user_service/internal/domain/cars"
)

type User struct {
	// TODO: int64
	ID             int `gorm:"primaryKey"`
	Name           string
	Surname        string
	Phone          string
	Email          string
	AccountAddress string `gorm:"column:account_address"`
	IsAdmin        bool   `gorm:"column:is_admin"`
	Password       string
	ChatID         int64      `gorm:"column:chat_id"`
	Cars           []cars.Car `gorm:"many2many:user_cars;joinForeignKey:UserID;joinReferences:CarID"`
	CreatedAt      time.Time  `gorm:"column:created_at"`
}

type Users []User
