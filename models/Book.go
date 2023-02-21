package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name   string `json:"name"`
	Genre  string `json:"genre"`
	Author string `json:"author"`
}
