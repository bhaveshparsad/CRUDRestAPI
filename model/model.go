package model

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	EmailId   string `json:"email"`
}
