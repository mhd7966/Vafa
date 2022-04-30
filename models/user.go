package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	NationalID string `json:"national_id" binding:"required, len=10"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Status     bool   `json:"status" defaul:"1"`
}
