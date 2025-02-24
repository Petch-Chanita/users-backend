package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Image     string    `json:"image"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `gorm:"unique" json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Auto Migrate: การสร้างตารางในฐานข้อมูล
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
