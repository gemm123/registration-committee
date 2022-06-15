package service

import (
	"net/http"

	"github.com/gemm123/registration-committee/repository"
	"github.com/gin-gonic/gin"
)

type Service interface{}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", gin.H{})
}

func (s *service) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{})
}
