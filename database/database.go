package database

import (
	"log"

	"github.com/gemm123/registration-committee/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDB(dsn string) {
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("cant connect to database")
	}

	log.Println("connected to database")
}

func Migrate() {
	// DB.Migrator().DropTable(&models.User{}, &models.Registration{}, &models.Committee{})
	DB.AutoMigrate(&models.User{}, &models.Committee{}, &models.Registration{})
}

func CloseConnection() {
	dbSQL, _ := DB.DB()
	dbSQL.Close()
}
