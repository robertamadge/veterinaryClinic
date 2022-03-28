package models

import (
	"time"
)

type Owner struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Cpf       int    `json:"cpf" gorm:"unique"`
	Name      string `json:"name"`
	Username  string `json:"username" gorm:"unique"`
	Email     string `json:"email" gorm:"unique"`
	Mobile    int    `json:"mobile"`
	Pets      []Pet  `json:"pets"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
