package config

import (
	"fmt"
	"log"
	"os"
	"wedding/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	// Memuat file .env
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Membaca nilai-nilai dari environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// Membuat Data Source Name (DSN)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Jakarta", dbHost, dbUser, dbPassword, dbName, dbPort)

	// Membuka koneksi ke database dengan GORM
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connected to database")

	// Melakukan migrasi ke database
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Gagal migrasi: %v", err)
	}
	log.Println("✅ Database migration completed")
}
