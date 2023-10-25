package main

import "time"

type Person struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic,omitempty"`
	Surname    string `json:"surname"`
	Gender     string `json:"gender"`
	Age        int    `json:"age"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Country    []Country
}

type Country struct {
	CountryID   string  `json:"country_id"`
	Probability float32 `json:"probability"`
}

func NewPerson(id int, name, patronymic, surname string) Person {
	return Person{
		ID:         id,
		Name:       name,
		Patronymic: patronymic,
		Surname:    surname,
	}
}
