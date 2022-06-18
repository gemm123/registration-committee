package repository

import (
	"errors"

	"github.com/gemm123/registration-committee/helper"
	"github.com/gemm123/registration-committee/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user models.User) error
	Authenticated(email string, password string) (models.User, error)
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

func (r *repository) Authenticated(email string, password string) (models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	ok := helper.CheckPassword(password, user.Password)
	if !ok {
		return models.User{}, errors.New("email or password incorrect")
	}

	return user, nil
}
