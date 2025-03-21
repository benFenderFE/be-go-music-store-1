package config

import (
	"log"
	"os"

	"be-go-music-store/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db

	SeedMockData()
}

type Instrument struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Brand string
	Price float64
	Stock int
}

// 🔥 เพิ่ม Mockup Data
func SeedMockData() {
	var count int64
	DB.Model(&models.Product{}).Count(&count)
	if count == 0 { // ถ้า table ว่างให้เติมข้อมูล
		products := []models.Product{
			{Name: "Fender Precision Bass", Brand: "Fender", Price: 1200.00, Stock: 10},
			{Name: "Gibson Les Paul", Brand: "Gibson", Price: 2500.00, Stock: 5},
			{Name: "Roland TD-17KV", Brand: "Roland", Price: 1500.00, Stock: 3},
			{Name: "Yamaha P-125 Digital Piano", Brand: "Yamaha", Price: 750.00, Stock: 7},
			{Name: "Shure SM58 Microphone", Brand: "Shure", Price: 100.00, Stock: 20},
		}
		DB.Create(&products)
	}
}
