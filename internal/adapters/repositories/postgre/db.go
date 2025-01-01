package postgre

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"remember-me/internal/domain/models"
)

type DB struct {
	Db *gorm.DB
}

func ConnectDb() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbHost, dbUser, dbPass, dbName, dbPort)

	// Connect to the database
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		fmt.Printf("Failed to connect to database! %s\n", dsn)
	}

	fmt.Println("Database connection successfully established")

	var ms = []interface{}{&models.User{}}

	errMigra := db.AutoMigrate(ms...)

	if errMigra != nil {
		fmt.Printf("Failed to migrate database! %s\n", dsn)
	}

	return db
}
