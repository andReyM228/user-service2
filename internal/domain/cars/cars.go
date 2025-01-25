package cars

import "time"

type Car struct {
	// TODO: int64
	ID        int `gorm:"primaryKey"`
	Name      string
	Model     string
	Price     int64
	Image     string
	Info      string
	CreatedAt time.Time `gorm:"column:created_at"`
}

type Cars []Car
