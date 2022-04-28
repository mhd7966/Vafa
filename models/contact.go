package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `json:"name" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Status    bool   `json:"status" binding:"required"`
}
