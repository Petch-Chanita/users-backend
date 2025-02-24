package services

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

// InitDB - ฟังก์ชันในการเชื่อมต่อฐานข้อมูล
func InitDB() (*gorm.DB, error) {
	// โหลดค่า .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// ดึงค่า DATABASE_URL จากไฟล์ .env
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// กำหนดค่าการเชื่อมต่อฐานข้อมูล
	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(5)

	DB = db
	return db, nil
}
