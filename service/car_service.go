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

// NewCarService creates a new CarService instance.
//
// Takes a repository.Car as a parameter.
// Returns a pointer to a CarService.
func NewCarService(repo repository.Car) *CarService {
	return &CarService{repo: repo}
}

// GetCars returns a list of cars.
//
// Takes a string as a parameter.
// Returns a slice of cars and an error.
func (c *CarService) GetCars(regNumber, mark string, page, pageSize string) ([]model.Car_model, error) {
	l.Info().Msg("GetCars service called")
	l.Debug().Msg(fmt.Sprintf("GetCars service called with regNum: %s, mark: %s, page: %s, pageSize: %s", regNumber, mark, page, pageSize))

	return c.repo.GetCars(regNumber, mark, page, pageSize)
}

// DeleteCar deletes a car by ID.
//
// Takes an integer ID as a parameter and returns an error.
func (c *CarService) DeleteCar(id int) error {
	l.Info().Msg("DeleteCar service called")
	l.Debug().Msg(fmt.Sprintf("DeleteCar service called with id: %d", id))

	return c.repo.DeleteCar(id)
}

// UpdateCar updates a car by ID.
//
// Takes an integer ID and a CarUpdate struct as parameters and returns an error.
func (c *CarService) UpdateCar(id int, car model.CarUpdate) error {
	l.Info().Msg("UpdateCar service called")
	l.Debug().Msg(fmt.Sprintf("UpdateCar service called with id: %d, car: %+v", id, car))

	return c.repo.UpdateCar(id, car)
}

// AddCar adds a new car.
//
// Takes a CarAdd struct as a parameter and returns an error.
func (c *CarService) AddCar(carAdd model.CarAdd) error {
	l.Info().Msg("AddCar service called")
	l.Debug().Msg(fmt.Sprintf("AddCar service called with car: %+v", carAdd))

	for _, regNum := range carAdd.RegNums {

		apiCar, err := callAPI(regNum)
		l.Debug().Msg(fmt.Sprintf("API response: %+v", apiCar))

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

// callAPI calls an external API and returns a CarApi struct.
//
// Takes a string as a parameter and returns a CarApi struct and an error.
func callAPI(regNum string) (model.CarApi, error) {
	l.Info().Msg("callAPI service called")
	l.Debug().Msg(fmt.Sprintf("callAPI service called with regNum: %s", regNum))

	url := fmt.Sprintf("https://external-api.com/info?regNum=%s", regNum)

	resp, err := http.Get(url)
	if err != nil {
		return model.CarApi{}, err
	}

	l.Debug().Msg(fmt.Sprintf("API response in method callAPI: %+v", resp))

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
