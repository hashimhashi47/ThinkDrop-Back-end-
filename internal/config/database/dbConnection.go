package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

// -> Make gorm connection with GORM to database
func Connection() *gorm.DB {

	// err := godotenv.Load("../../.env")
	// if err != nil {
	// 	log.Fatal("❌ Error loading .env file")
	// }
	var err error

	_ = godotenv.Load()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	sqlDB, _ := DB.DB()
	fmt.Println(sqlDB.Stats())

	fmt.Print("Database integerted succesfully✅")

	return DB
}
