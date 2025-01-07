package postgre

import (
	"auth-sample-app/internal/domain/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type DB struct {
	Db *gorm.DB
}

func ConnectDb() *gorm.DB {

	_ = godotenv.Load(".env")

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=UTC",
		dbHost, dbUser, dbPass, dbName)

	// Connect to the database
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

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
