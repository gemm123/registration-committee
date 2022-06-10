package models

type Registration struct {
	Division    string `gorm:"type:varchar(255)"`
	Status      string `gorm:"type:varchar(255)"`
	UserID      uint
	CommitteeID uint
}
