package service

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gemm123/registration-committee/helper"
	"github.com/gemm123/registration-committee/models"
	"github.com/gemm123/registration-committee/repository"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var authenticated int

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) Home(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"user":          user,
		"authenticated": authenticated,
	})
}

func (s *service) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{})
}

func (s *service) PostLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	session := sessions.Default(c)

	user, err := s.repository.Authenticated(email, password)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.tmpl", gin.H{
			"error": err,
		})
		return
	}

	session.Set("user", user.Name)
	if err := session.Save(); err != nil {
		c.HTML(http.StatusInternalServerError, "login.tmpl", gin.H{
			"error": err,
		})
		return
	}

	authenticated = 1

	c.Redirect(http.StatusFound, "/")

}

func (s *service) Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.tmpl", gin.H{})
}

func (s *service) PostRegister(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	gender := c.PostForm("gender")

	password := c.PostForm("password")
	hashPass, err := helper.HashPassword(password)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "register.tmpl", gin.H{
			"error": err,
		})
		return
	}

	generation := c.PostForm("generation")
	intGeneration, _ := strconv.Atoi(generation)

	newUser := models.User{
		Name:       name,
		Email:      email,
		Password:   hashPass,
		Gender:     gender,
		Generation: intGeneration,
	}

	err = s.repository.CreateUser(newUser)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "register.tmpl", gin.H{
			"error": err,
		})
		return
	}

	c.Redirect(http.StatusFound, "/login")
}

func (s *service) Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.HTML(http.StatusInternalServerError, "home.tmpl", gin.H{
			"err":           errors.New("cant get session"),
			"authenticated": authenticated,
		})
		return
	}

	session.Delete("user")
	if err := session.Save(); err != nil {
		c.HTML(http.StatusInternalServerError, "home.tmpl", gin.H{
			"err":           errors.New("cant save session"),
			"authenticated": authenticated,
		})
		return
	}

	authenticated = 0

	c.Redirect(http.StatusFound, "/")
}

func (s *service) Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{
		"authenticated": authenticated,
	})
}
