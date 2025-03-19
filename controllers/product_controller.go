package controllers

import (
	"be-go-music-store/config"
	"be-go-music-store/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ดึงสินค้าทั้งหมด
func GetProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

// เพิ่มสินค้า
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&product)
	c.JSON(http.StatusCreated, product)
}

// อัปเดตสินค้า
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

// ลบสินค้า
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Product{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
