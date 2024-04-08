package service

import (
	"carDirectory/model"
	"carDirectory/repository"
	"fmt"
	"github.com/goccy/go-json"
	"net/http"
)

type CarService struct {
	repo repository.Car
}

func NewCarService(repo repository.Car) *CarService {
	return &CarService{repo: repo}
}

func (c *CarService) GetCars(regNumber, mark string, page, pageSize string) ([]model.Car_model, error) {
	return c.repo.GetCars(regNumber, mark, page, pageSize)
}

func (c *CarService) DeleteCar(id int) error {
	return c.repo.DeleteCar(id)
}

func (c *CarService) UpdateCar(id int, car model.CarUpdate) error {

	fmt.Println(car)

	return c.repo.UpdateCar(id, car)
}

func (c *CarService) AddCar(carAdd model.CarAdd) error {
	for _, regNum := range carAdd.RegNums {

		apiCar, err := callAPI(regNum)
		if err != nil {
			return err
		}

		if err := c.repo.AddCar(apiCar); err != nil {
			return err
		}
		fmt.Printf("Adding car with registration number %s\n", regNum)
	}

	return nil
}

func callAPI(regNum string) (model.CarApi, error) {
	url := fmt.Sprintf("https://external-api.com/info?regNum=%s", regNum)

	resp, err := http.Get(url)
	if err != nil {
		return model.CarApi{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.CarApi{}, fmt.Errorf("external API returned non-200 status code: %d", resp.StatusCode)
	}

	var car model.CarApi
	if err := json.NewDecoder(resp.Body).Decode(&car); err != nil {
		return model.CarApi{}, err
	}

	return car, nil
}
