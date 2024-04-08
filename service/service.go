package service

import (
	"carDirectory/model"
	"carDirectory/repository"
)

type Car interface {
	GetCars(regNumber, mark string, page, pageSize string) ([]model.Car_model, error)
	DeleteCar(id int) error
	UpdateCar(idInt int, car model.CarUpdate) error
	AddCar(carAdd model.CarAdd) error
}

type Service struct {
	Car
}

func NewService(repo *repository.CarRepository) *Service {
	return &Service{
		Car: NewCarService(repo.Car),
	}
}
