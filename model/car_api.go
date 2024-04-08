package model

type CarApi struct {
	RegNum string    `json:"regNum"`
	Mark   string    `json:"mark"`
	Model  string    `json:"model"`
	Year   int       `json:"year"`
	Owner  PeopleApi `json:"owner"`
}

type PeopleApi struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}
