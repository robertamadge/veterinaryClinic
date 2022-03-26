package handlers

import (
	"gorm.io/gorm"
	"html/template"
)

type handler struct {
	DB *gorm.DB
}
var tpl *template.Template

func New(db *gorm.DB) handler {
	return handler{db}
}