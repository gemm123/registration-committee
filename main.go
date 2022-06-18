package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gemm123/registration-committee/database"
	"github.com/gemm123/registration-committee/render"
	"github.com/gemm123/registration-committee/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error load .env files")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	connectDB := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName)
	defer database.CloseConnection()

	database.InitDB(connectDB)
	database.Migrate()

	r := gin.Default()
	r.Static("/assets", "./assets")
	r.HTMLRender = render.LoadTemplates("./templates")
	r.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("secret"))))

	routes.Routes(r)
	r.Run()
}
