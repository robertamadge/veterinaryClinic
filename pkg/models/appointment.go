package models

import "time"

type Appointment struct {
	Id        int       `json:"id" gorm:"primary_key"`
	PetsID    uint      `gorm:"foreign_key:Id"`
	DoctorsID uint      `gorm:"foreign_key:Id"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
