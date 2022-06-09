package main

import (
	"net/http"

	"github.com/gemm123/registration-committee/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB("host=localhost user=postgres password=gemmq123456 dbname=registration_committee port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	database.Migrate()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{})
	})

	r.Run()
}
