package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	NationalID string `json:"national_id"  validate:"required,len=10"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Status     int    `json:"status" default:"1"`
}
