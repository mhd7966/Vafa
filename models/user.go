package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	NationalID string `json:"national_id"  validate:"len=10"`
	FirstName  string `json:"first_name"  validate:"alpha"`
	LastName   string `json:"last_name"  validate:"alpha"`
	Status     int    `json:"status" default:"1"`
}
