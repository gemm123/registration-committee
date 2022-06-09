package models

import (
	"time"

	"gorm.io/gorm"
)

type Committee struct {
	gorm.Model
	Name          string
	Year          time.Time
	Registrations []Registration
}
