package models

type Product struct {
	ID    uint    `gorm:"primaryKey"`
	Name  string  `json:"name"`
	Brand string  `json:"brand"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}
