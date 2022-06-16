package service

import (
	"net/http"
	"strconv"

	"github.com/gemm123/registration-committee/models"
	"github.com/gemm123/registration-committee/repository"
	"github.com/gin-gonic/gin"
)

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

func (s *service) Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.tmpl", gin.H{})
}

func (s *service) PostRegister(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	gender := c.PostForm("gender")
	generation := c.PostForm("generation")
	intGeneration, _ := strconv.Atoi(generation)

	newUser := models.User{
		Name:       name,
		Email:      email,
		Password:   password,
		Gender:     gender,
		Generation: intGeneration,
	}

	err := s.repository.CreateUser(newUser)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "register.tmpl", gin.H{
			"error": err,
		})
		return
	}

	c.Redirect(http.StatusFound, "/login")
}
