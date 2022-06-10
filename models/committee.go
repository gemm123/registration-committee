package models

import (
	"time"

	"gorm.io/gorm"
)

type Committee struct {
	gorm.Model
	Name          string `gorm:"type:varchar(255)"`
	Year          time.Time
	Registrations []Registration
}
