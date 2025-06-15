package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Gagal memuat file .env, lanjut pakai environment bawaan")
		}
	}

	host := os.Getenv("DB_HOST_DEBUG")
	port := os.Getenv("DB_PORT_DEBUG")
	user := os.Getenv("DB_USER_DEBUG")
	pass := os.Getenv("DB_PASSWORD_DEBUG")
	dbname := os.Getenv("DB_NAME_DEBUG")
	timezone := os.Getenv("DB_TIMEZONE_DEBUG")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		host, user, pass, dbname, port, timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	DB = db
	fmt.Println("Database berhasil terhubung")
}
