package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Todo struct {
	gorm.Model
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
