package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID          uint `gorm:"column:id;primary_key;"`
	Name        string
	Description string
}

type TodoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TodoResponse struct {
	TodoRequest
	ID uint `json:"id"`
}
