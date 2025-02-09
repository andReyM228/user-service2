package cars

import (
	"database/sql"
	"time"
	"user_service/internal/domain/cars"
)

const tableName = "cars"

type carDB struct {
	// TODO: int64
	ID        int `gorm:"primaryKey"`
	Name      string
	Model     string
	Price     int64
	Image     string
	Info      sql.NullString
	CreatedAt time.Time `gorm:"column:created_at"`
}

type carsDB []carDB

func fromDomain(car cars.Car) carDB {
	return carDB{
		ID:    car.ID,
		Name:  car.Name,
		Model: car.Model,
		Price: car.Price,
		Image: car.Image,
		Info: sql.NullString{
			String: car.Info,
		},
		CreatedAt: car.CreatedAt,
	}
}

func (c carDB) toDomain() cars.Car {
	return cars.Car{
		ID:        c.ID,
		Name:      c.Name,
		Model:     c.Model,
		Price:     c.Price,
		Image:     c.Image,
		Info:      c.Info.String,
		CreatedAt: c.CreatedAt,
	}
}

func toDomainList(carsDB []carDB) cars.Cars {
	result := make([]cars.Car, 0, len(carsDB))

	for _, car := range carsDB {
		result = append(result, car.toDomain())
	}

	return result
}

func (carDB) TableName() string {
	return tableName
}
