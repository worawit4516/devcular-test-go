package models

import (
	"gorm.io/gorm"
)

// Model for ORM
type Todo struct {
	gorm.Model
	ID          uint `gorm:"column:id;primary_key;"`
	Name        string
	Description string
}

// Model for trans request from JSON to Struct
type TodoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Model for response
type TodoResponse struct {
	TodoRequest
	ID uint `json:"id"`
}
