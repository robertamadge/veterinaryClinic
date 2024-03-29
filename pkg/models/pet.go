package models

import (
	"time"
)

type Pet struct {
	Id        int       `json:"id" gorm:"primary_key"`
	OwnersID  uint      `gorm:"foreign_key:Id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
