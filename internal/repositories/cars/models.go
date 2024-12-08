package cars

import (
	"database/sql"
	"time"
	"user_service/internal/domain"
)

type CarDB struct {
	ID        int
	Name      string
	Model     string
	Price     int64
	Image     string
	Info      sql.NullString
	CreatedAt time.Time `db:"created_at"`
}

func fromDomain(car domain.Car) CarDB {
	return CarDB{
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

func (c CarDB) toDomain() domain.Car {
	return domain.Car{
		ID:        c.ID,
		Name:      c.Name,
		Model:     c.Model,
		Price:     c.Price,
		Image:     c.Image,
		Info:      c.Info.String,
		CreatedAt: c.CreatedAt,
	}
}

func toDomainList(carsDB []CarDB) domain.Cars {
	result := make([]domain.Car, 0, len(carsDB))

	for _, car := range carsDB {
		result = append(result, car.toDomain())
	}

	return result
}
