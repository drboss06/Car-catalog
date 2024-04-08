package repository

import (
	"carDirectory/model"
	"github.com/jmoiron/sqlx"
)

type Car interface {
	GetCars(regNumber, mark string, page, pageSize string) ([]model.Car_model, error)
	DeleteCar(id int) error
	UpdateCar(idInt int, car model.CarUpdate) error
	AddCar(apiCar model.CarApi) error
}

type CarRepository struct {
	Car
}

func NewCarRepository(db *sqlx.DB) *CarRepository {
	return &CarRepository{Car: NewCarPostgres(db)}
}
