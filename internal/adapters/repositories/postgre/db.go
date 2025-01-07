package postgre

import (
	"auth-sample-app/internal/domain/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type DB struct {
	Db *gorm.DB
}

func ConnectDb() *gorm.DB {

	if os.Getenv("ENV") != "prod" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("Error loading .env file")
		}
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Failed to connect to database! %s\n", dsn)
	} else {
		fmt.Println("Database connection successfully established")
	}

	var ms = []interface{}{&models.User{}}

	errMigra := db.AutoMigrate(ms...)

	if errMigra != nil {
		fmt.Printf("Failed to migrate database! %s\n", dsn)
	}

	return db
}
