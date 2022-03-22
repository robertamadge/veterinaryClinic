package models

type Owner struct {
	Id     int
	Cpf    int    `json:"cpf"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile int    `json:"mobile"`
	Pets   []Pet
}
