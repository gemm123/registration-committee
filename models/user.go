package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email         string `form:"email" gorm:"type:varchar(255)"`
	Password      string `form:"password" gorm:"type:varchar(255)"`
	Name          string `form:"name" gorm:"type:varchar(255)"`
	Gender        string `form:"gender" gorm:"type:varchar(255)"`
	Generation    int    `form:"generation"`
	Registrations []Registration
}
