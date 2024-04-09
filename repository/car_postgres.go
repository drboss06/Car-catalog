package repository

import (
	"carDirectory/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
)

type CarPostgres struct {
	db *sqlx.DB
}

// NewCarPostgres creates a new CarPostgres instance.
//
// Takes a pointer to a sqlx.DB as the parameter.
// Returns a pointer to a CarPostgres instance.
func NewCarPostgres(db *sqlx.DB) *CarPostgres {
	return &CarPostgres{db: db}
}

// NewCarPostgres creates a new CarPostgres instance.
//
// db: the database connection for CarPostgres.
// *CarPostgres: returns a pointer to the CarPostgres struct.
func (c *CarPostgres) GetCars(regNumber, mark string, page, pageSize string) ([]model.Car_model, error) {
	l.Info().Msg("GetCars service called")

	query := "SELECT * FROM car WHERE TRUE"

	args := []interface{}{}

	if regNumber != "" {
		query += " AND $1 = ANY(regNum)"
		args = append(args, regNumber)
	}
	if mark != "" {
		query += " AND mark = $2"
		args = append(args, mark)
	}

	query += " ORDER BY id"

	if page != "" && pageSize != "" {
		query += " LIMIT $4 OFFSET $5"
		pageInt, _ := strconv.Atoi(page)
		pageSizeInt, _ := strconv.Atoi(pageSize)
		args = append(args, pageSizeInt, (pageInt-1)*pageSizeInt)
	}

	l.Debug().Msg(fmt.Sprintf("GetCars service called with query: %s, args: %v", query, args))

	rows, err := c.db.Query(query, args...)

	l.Debug().Msg(fmt.Sprintf("GetCars service called with rows: %+v", rows))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []model.Car_model

	for rows.Next() {
		var car model.Car_model

		if err := rows.Scan(&car.Id, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner_id); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		l.Debug().Msg(fmt.Sprintf("GetCars service called with car: %+v", car))

		cars = append(cars, car)
	}

	l.Debug().Msg(fmt.Sprintf("GetCars service called with cars: %+v", cars))

	l.Info().Msg("Cars fetched successfully")

	return cars, nil
}

// DeleteCar deletes a car from the database given its ID.
//
// Parameter(s):
// - id int: the ID of the car to be deleted.
// Return type(s):
// - error: an error, if any, during the deletion process.
func (c *CarPostgres) DeleteCar(id int) error {
	l.Info().Msg("DeleteCar service called")

	resp, err := c.db.Exec("DELETE FROM car WHERE id = $1", id)

	l.Debug().Msg(fmt.Sprintf("DeleteCar service called with resp: %+v", resp))

	if err != nil {
		return err
	}

	l.Info().Msg("Car deleted successfully")

	return nil
}

// UpdateCar updates a car in the database based on the provided ID and new car information.
//
// Parameters:
// - id: the ID of the car to update
// - car: the new car information to update
// Return type: error
func (c *CarPostgres) UpdateCar(id int, car model.CarUpdate) error {
	l.Info().Msg("UpdateCar service called")

	query := "UPDATE car SET"
	setValues := make([]string, 0)
	params := make([]interface{}, 0)
	argId := 1

	if car.RegNum != "" {
		setValues = append(setValues, fmt.Sprintf("regNum=$%d", argId))
		params = append(params, car.RegNum)
		argId++
	}
	if car.Mark != "" {
		setValues = append(setValues, fmt.Sprintf("mark=$%d", argId))
		params = append(params, car.Mark)
		argId++
	}
	if car.Model != "" {
		setValues = append(setValues, fmt.Sprintf("model=$%d", argId))
		params = append(params, car.Model)
		argId++
	}
	if car.Year != 0 {
		setValues = append(setValues, fmt.Sprintf("year=$%d", argId))
		params = append(params, car.Year)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query = fmt.Sprintf("UPDATE car SET %s WHERE id = $%d", setQuery, argId)

	l.Debug().Msg(fmt.Sprintf("UpdateCar service called with query: %s, params: %v", query, params))

	params = append(params, id)

	_, err := c.db.Exec(query, params...)
	if err != nil {
		return err
	}

	l.Info().Msg("Car updated successfully")

	return nil
}

// AddCar adds a new car to the database.
//
// Parameter(s):
// - apiCar: the car information to be added.
// Return type(s):
// - error: an error, if any, during the addition process.
func (c *CarPostgres) AddCar(apiCar model.CarApi) error {
	l.Info().Msg("AddCar service called")

	query := fmt.Sprintf(`
        WITH people AS (
    	INSERT INTO people (name, surname, patronymic)
        VALUES ($1, $2, $3)
        RETURNING id
		)
		INSERT INTO car (regNum, mark, model, year, owner)
		SELECT $4, $5, $6, $7, id
		FROM people;
    `,
		apiCar.Owner.Name, apiCar.Owner.Surname, apiCar.Owner.Patronymic,
		apiCar.RegNum, apiCar.Mark, apiCar.Model, apiCar.Year)

	_, err := c.db.Exec(query, apiCar.Owner.Name, apiCar.Owner.Surname, apiCar.Owner.Patronymic,
		apiCar.RegNum, apiCar.Mark, apiCar.Model, apiCar.Year)

	l.Debug().Msg(fmt.Sprintf("AddCar service called with query: %s", query))

	if err != nil {
		return err
	}

	l.Info().Msg("Car added successfully")

	return nil
}
