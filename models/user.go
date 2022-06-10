package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email         string `gorm:"type:varchar(255)"`
	Password      string `gorm:"type:varchar(255)"`
	Name          string `gorm:"type:varchar(255)"`
	Gender        string `gorm:"type:varchar(255)"`
	Generation    int
	Registrations []Registration
}
