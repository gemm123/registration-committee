package repository

import (
	"github.com/gemm123/registration-committee/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user models.User) error
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		DB: db,
	}
}

func (r *repository) CreateUser(user models.User) error {
	err := r.DB.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}
