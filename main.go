package main

import (
	"log"
	"os"

	"be-go-music-store/config"
	"be-go-music-store/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// โหลดค่าจากไฟล์ .env
	godotenv.Load()

	// เชื่อมต่อ Database
	config.ConnectDatabase()

	// สร้าง Gin Router
	r := gin.Default()

	// เพิ่ม Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// เรียกใช้งาน Routes
	routes.SetupRoutes(r)

	// เริ่มเซิร์ฟเวอร์
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server is running on port:", port)
	r.Run(":" + port)
}
