package main

import (
	"log"
	"os"

	"users-backend/models"
	"users-backend/routes"
	"users-backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// ตั้งค่า CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // อนุญาตให้ทุกโดเมนสามารถเข้าถึงได้
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // เมธอดที่อนุญาต
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // เฮดเดอร์ที่อนุญาต
		AllowCredentials: true,                                                // ถ้าต้องการส่งข้อมูล Cookie หรือ Credential
	}))

	// เชื่อมต่อฐานข้อมูล
	db, err := services.InitDB() // ฟังก์ชันในการเชื่อมต่อฐานข้อมูล
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// ตั้งค่า migration
	models.Migrate(db)

	routes.NewRouter(r)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	// เริ่มต้นเซิร์ฟเวอร์
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)

}
