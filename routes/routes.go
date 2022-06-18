package routes

import (
	"github.com/gemm123/registration-committee/database"
	"github.com/gemm123/registration-committee/repository"
	"github.com/gemm123/registration-committee/service"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	repository := repository.NewRepository(database.DB)
	service := service.NewService(repository)

	r.GET("/", service.Home)
	r.GET("/login", service.Login)
	r.POST("/login", service.PostLogin)
	r.GET("/register", service.Register)
	r.POST("/register", service.PostRegister)
	r.GET("/logout", service.Logout)
}
