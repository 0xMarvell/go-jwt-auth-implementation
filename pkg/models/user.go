package models

import "gorm.io/gorm"

// User is the database model for user objects
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}
