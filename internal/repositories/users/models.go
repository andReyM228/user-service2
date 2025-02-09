package users

import (
	"time"
	"user_service/internal/domain/cars"
	"user_service/internal/domain/users"
)

const tableName = "users"

type userDB struct {
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

type usersDB []userDB

func fromDomain(user users.User) userDB {
	return userDB{
		ID:             user.ID,
		Name:           user.Name,
		Surname:        user.Surname,
		Phone:          user.Phone,
		Email:          user.Email,
		AccountAddress: user.AccountAddress,
		IsAdmin:        user.IsAdmin,
		Password:       user.Password,
		ChatID:         user.ChatID,
		Cars:           user.Cars,
		CreatedAt:      user.CreatedAt,
	}
}

func (u userDB) toDomain() users.User {
	return users.User{
		ID:             u.ID,
		Name:           u.Name,
		Surname:        u.Surname,
		Phone:          u.Phone,
		Email:          u.Email,
		AccountAddress: u.AccountAddress,
		IsAdmin:        u.IsAdmin,
		Password:       u.Password,
		ChatID:         u.ChatID,
		Cars:           u.Cars,
		CreatedAt:      u.CreatedAt,
	}
}

func toDomainList(usersDB usersDB) users.Users {
	result := make([]users.User, 0, len(usersDB))

	for _, user := range usersDB {
		result = append(result, user.toDomain())
	}

	return result
}

func (userDB) TableName() string {
	return tableName
}
