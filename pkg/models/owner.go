package models

import (
	"time"
)

type Owner struct {
	Id        int       `json:"id" gorm:"primary_key"`
	Cpf       int       `json:"cpf" gorm:"unique_index; not null"`
	Name      string    `json:"name"`
	Username  string    `json:"username" gorm:"unique_index; not null"`
	Email     string    `json:"email" gorm:"unique_index; not null" db:"username"`
	Mobile    string    `json:"mobile"`
	Password  string    `json:"password" db:"password"`
	PetsID    uint      `gorm:"foreign_key:Id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
