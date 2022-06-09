package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email         string
	Password      string
	Name          string
	Gender        string
	Generation    int
	Registrations []Registration
}
