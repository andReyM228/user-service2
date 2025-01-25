package cars

import (
	"time"
	"user_service/internal/domain/cars"
)

type carResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Model     string    `json:"model"`
	Price     int64     `json:"price"`
	Image     string    `json:"image"`
	Info      string    `json:"info"`
	CreatedAt time.Time `json:"createdAt"`
}

func toResponse(car cars.Car) carResponse {
	return carResponse{
		ID:        car.ID,
		Name:      car.Name,
		Model:     car.Model,
		Price:     car.Price,
		Image:     car.Image,
		Info:      car.Info,
		CreatedAt: car.CreatedAt,
	}
}

func toResponseList(cars cars.Cars) []carResponse {
	var result []carResponse

	for _, car := range cars {
		result = append(result, toResponse(car))
	}

	return result
}
