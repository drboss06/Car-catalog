package repository

import (
	"carDirectory/logger"
	"carDirectory/model"
	"github.com/jmoiron/sqlx"
)

type Car interface {
	GetCars(regNumber, mark string, page, pageSize string) ([]model.Car_model, error)
	DeleteCar(id int) error
	UpdateCar(idInt int, car model.CarUpdate) error
	AddCar(apiCar model.CarApi) error
}

var l = logger.Get()

type CarRepository struct {
	Car
}

// NewCarRepository creates a new CarRepository instance.
//
// It takes a pointer to a sqlx.DB as a parameter.
// Returns a pointer to a CarRepository.
func NewCarRepository(db *sqlx.DB) *CarRepository {
	return &CarRepository{Car: NewCarPostgres(db)}
}
