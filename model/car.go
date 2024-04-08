package model

type Car_model struct {
	Id       int    `json:"-" db:"id"`
	RegNum   string `json:"regNum" db:"regNum"`
	Mark     string `json:"mark" db:"mark"`
	Model    string `json:"model" db:"model"`
	Year     int    `json:"year" db:"year"`
	Owner_id int    `json:"-" db:"owner"`
}
