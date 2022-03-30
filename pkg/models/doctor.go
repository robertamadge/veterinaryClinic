package models

import "time"

type Doctor struct {
	Id        int       `json:"id" gorm:"primary_key"`
	Crm       int       `json:"crm" gorm:"unique_index; not null"`
	Cpf       int       `json:"cpf" gorm:"unique_index; not null"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
