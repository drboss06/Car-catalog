package repository

import (
	"carDirectory/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type CarPostgres struct {
	db *sqlx.DB
}

func NewCarPostgres(db *sqlx.DB) *CarPostgres {
	return &CarPostgres{db: db}
}

// NewCarPostgres creates a new CarPostgres instance.
//
// db: the database connection for CarPostgres.
// *CarPostgres: returns a pointer to the CarPostgres struct.
func (c *CarPostgres) GetCars(regNumber, mark string, page, pageSize string) ([]model.Car_model, error) {

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

	fmt.Println(query, args)

	rows, err := c.db.Query(query, args...)

	fmt.Println(rows, err, "****************************")

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

		fmt.Println(car)

		cars = append(cars, car)
	}

	fmt.Println(cars)

	return cars, nil
}

func (c *CarPostgres) DeleteCar(id int) error {
	_, err := c.db.Exec("DELETE FROM car WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}

func (c *CarPostgres) UpdateCar(id int, car model.CarUpdate) error {
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
	params = append(params, id)

	_, err := c.db.Exec(query, params...)
	if err != nil {
		return err
	}

	logrus.Println("Car updated successfully")

	return nil
}

func (c *CarPostgres) AddCar(apiCar model.CarApi) error {
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

	if err != nil {
		return err
	}

	return nil
}
