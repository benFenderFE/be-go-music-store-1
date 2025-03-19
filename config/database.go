package config

import (
	"log"

	"be-go-music-store/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "postgresql://postgres:SKdpzWKInBFBCOlSMksMJYHQWZsJIQfY@ballast.proxy.rlwy.net:10576/railway"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	DB.AutoMigrate(&models.Product{}) // สร้าง table อัตโนมัติ

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
